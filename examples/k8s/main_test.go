package main

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

func TestObjectNamePod(t *testing.T) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "demo-pod"},
	}
	if got := objectName(pod); got != "demo-pod" {
		t.Fatalf("expected demo-pod, got %q", got)
	}
}

func TestObjectNameDeletedFinalStateUnknown(t *testing.T) {
	tombstone := cache.DeletedFinalStateUnknown{
		Obj: &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "deleted-pod"}},
	}
	if got := objectName(tombstone); got != "deleted-pod" {
		t.Fatalf("expected deleted-pod, got %q", got)
	}
}
