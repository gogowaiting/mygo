package interview

import "testing"

func TestFind(t *testing.T) {
	q, ok := Find("go-gmp")
	if !ok {
		t.Fatal("expected go-gmp exists")
	}
	if q.Topic == "" || q.Answer == "" {
		t.Fatal("question should have topic and answer")
	}
}

func TestTopicsNotEmpty(t *testing.T) {
	if len(Topics()) == 0 {
		t.Fatal("topics should not be empty")
	}
}
