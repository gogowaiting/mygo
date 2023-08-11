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
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/sha256"
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
	io.WriteString(h, k)
	if sha256.Sum224([]byte(hex.EncodeToString(h.Sum(nil)))) == sum {
		return true
	}
	return false
}

// Encrypt
func TripleDesEncrypt(origData []byte) (_ []byte, err error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	defer func() {
		if recover() != nil {
			err = errors.New("Authorization failed. ")
		}
	}()
	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	cryptedData := make([]byte, len(origData))
	blockMode.CryptBlocks(cryptedData, origData)
	return cryptedData, nil
}

// Decrypt
func TripleDesDecrypt(cryptedData []byte) (_ []byte, err error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	defer func() {
		if recover() != nil {
			err = errors.New("Prohibit the installation of unauthorized software packages. ")
		}
	}()
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	origData := make([]byte, len(cryptedData))
	blockMode.CryptBlocks(origData, cryptedData)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// Chart Values Encrypt
func ChartValuesEncrypt(values string) (string, error) {
	valuesData := []byte(values)
	valuesData, err := TripleDesEncrypt(valuesData)
	if err != nil {
		return values, err
	}
	return string(valuesData), nil
}

// Chart Values Decrypt
func ChartValuesDecrypt(values string) (string, error) {
	valuesData := []byte(values)
	valuesData, err := TripleDesDecrypt(valuesData)
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
		h.Manifest = string(valuesData)
	}
	return hooks, nil
}

// Hooks Decrypt
func HooksDecrypt(hooks []*release.Hook) ([]*release.Hook, error) {
	for _, h := range hooks {
		valuesData := []byte(h.Manifest)
		valuesData, err := TripleDesDecrypt(valuesData)
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
