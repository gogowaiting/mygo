package learn

import "testing"

func TestFind(t *testing.T) {
	d, ok := Find("basic-01-hello")
	if !ok {
		t.Fatal("expected basic-01-hello in catalog")
	}
	if d.Path == "" {
		t.Fatal("expected demo path")
	}
}

func TestTopics(t *testing.T) {
	topics := Topics()
	if len(topics) == 0 {
		t.Fatal("expected non-empty topics")
	}
}
