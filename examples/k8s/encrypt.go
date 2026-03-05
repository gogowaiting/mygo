//go:build demo

package main

/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"

	"k8s.io/helm/pkg/proto/hapi/chart"
	"k8s.io/helm/pkg/proto/hapi/release"
)

var sum = sha256.Sum224([]byte("5f4dcc3b5aa765d61d8327deb882cf99"))
var key = sum[:24]

func CheckKey(k string) bool {
	h := md5.New()
	_, _ = io.WriteString(h, k)
	return sha256.Sum224([]byte(hex.EncodeToString(h.Sum(nil)))) == sum
}

// Encrypt
// NOTE: kept function name for backward compatibility; implementation now uses AES-GCM.
func TripleDesEncrypt(origData []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aead.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := aead.Seal(nil, nonce, origData, nil)
	return append(nonce, ciphertext...), nil
}

// Decrypt
func TripleDesDecrypt(cryptedData []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	if len(cryptedData) < aead.NonceSize() {
		return nil, errors.New("ciphertext too short")
	}
	nonce := cryptedData[:aead.NonceSize()]
	ciphertext := cryptedData[aead.NonceSize():]
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, errors.New("authorization failed")
	}
	return plaintext, nil
}

// Chart Values Encrypt
func ChartValuesEncrypt(values string) (string, error) {
	valuesData := []byte(values)
	valuesData, err := TripleDesEncrypt(valuesData)
	if err != nil {
		return values, err
	}
	return base64.StdEncoding.EncodeToString(valuesData), nil
}

// Chart Values Decrypt
func ChartValuesDecrypt(values string) (string, error) {
	valuesData, err := base64.StdEncoding.DecodeString(values)
	if err != nil {
		return values, err
	}
	valuesData, err = TripleDesDecrypt(valuesData)
	if err != nil {
		return values, err
	}
	return string(valuesData), nil
}

// Chart Templates Encrypt
func ChartTemplatesEncrypt(templates []*chart.Template) ([]*chart.Template, error) {
	for _, v := range templates {
		var err error
		v.Data, err = TripleDesEncrypt(v.GetData())
		if err != nil {
			return templates, err
		}
	}
	return templates, nil
}

// Chart Templates Decrypt
func ChartTemplatesDecrypt(templates []*chart.Template) ([]*chart.Template, error) {
	for _, v := range templates {
		var err error
		v.Data, err = TripleDesDecrypt(v.GetData())
		if err != nil {
			return templates, err
		}
	}
	return templates, nil
}

// Hooks Encrypt
func HooksEncrypt(hooks []*release.Hook) ([]*release.Hook, error) {
	for _, h := range hooks {
		valuesData := []byte(h.Manifest)
		valuesData, err := TripleDesEncrypt(valuesData)
		if err != nil {
			return hooks, err
		}
		h.Manifest = base64.StdEncoding.EncodeToString(valuesData)
	}
	return hooks, nil
}

// Hooks Decrypt
func HooksDecrypt(hooks []*release.Hook) ([]*release.Hook, error) {
	for _, h := range hooks {
		valuesData, err := base64.StdEncoding.DecodeString(h.Manifest)
		if err != nil {
			return hooks, err
		}
		valuesData, err = TripleDesDecrypt(valuesData)
		if err != nil {
			return hooks, err
		}
		h.Manifest = string(valuesData)
	}
	return hooks, nil
}

// Chart Encrypt
func ChartEncrypt(chart *chart.Chart) (*chart.Chart, error) {
	var err error
	chart.Values.Raw, err = ChartValuesEncrypt(chart.Values.Raw)
	if err != nil {
		return nil, err
	}
	chart.Templates, err = ChartTemplatesEncrypt(chart.Templates)
	if err != nil {
		return nil, err
	}
	return chart, nil
}

// Chart Decrypt
func ChartDecrypt(chart *chart.Chart) (*chart.Chart, error) {
	var err error
	chart.Values.Raw, err = ChartValuesDecrypt(chart.Values.Raw)
	if err != nil {
		return nil, err
	}
	chart.Templates, err = ChartTemplatesDecrypt(chart.Templates)
	if err != nil {
		return nil, err
	}
	return chart, nil
}

// Release Encrypt
func ReleaseEncrypt(release *release.Release) (*release.Release, error) {
	var err error
	release.Chart, err = ChartEncrypt(release.Chart)
	if err != nil {
		return nil, err
	}
	release.Manifest, err = ChartValuesEncrypt(release.Manifest)
	if err != nil {
		return nil, err
	}
	release.Config.Raw, err = ChartValuesEncrypt(release.Config.Raw)
	if err != nil {
		return nil, err
	}
	release.Hooks, err = HooksEncrypt(release.Hooks)
	if err != nil {
		return nil, err
	}
	return release, nil
}

// Release Decrypt
func ReleaseDecrypt(release *release.Release) (*release.Release, error) {
	var err error
	release.Chart, err = ChartDecrypt(release.Chart)
	if err != nil {
		return nil, err
	}
	release.Manifest, err = ChartValuesDecrypt(release.Manifest)
	if err != nil {
		return nil, err
	}
	release.Config.Raw, err = ChartValuesDecrypt(release.Config.Raw)
	if err != nil {
		return nil, err
	}
	release.Hooks, err = HooksDecrypt(release.Hooks)
	if err != nil {
		return nil, err
	}
	return release, nil
}
