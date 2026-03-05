package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()
	fmt.Printf("kubeconfig: %s\n", *kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("build kubeconfig failed: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("create kubernetes client failed: %v", err)
	}
	stopCh := make(chan struct{})
	go func() {
		<-ctx.Done()
		close(stopCh)
	}()

	sharedInformers := informers.NewSharedInformerFactory(clientset, time.Minute)
	podInformer := sharedInformers.Core().V1().Pods().Informer()

	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("New Pod Add to Store:", objectName(obj))
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			fmt.Println("Old Pod Update to New", objectName(oldObj), objectName(newObj))
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("Pod Delete form Store:", objectName(obj))
		},
	})

	runtimeClassInformer := sharedInformers.Node().V1().RuntimeClasses().Informer()
	runtimeClassInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("New RuntimeClass Add to Store:", objectName(obj))
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("RuntimeClass Delete from Store:", objectName(obj))
		},
	})

	sharedInformers.Start(stopCh)
	if !cache.WaitForCacheSync(stopCh, podInformer.HasSynced, runtimeClassInformer.HasSynced) {
		log.Fatal("cache sync timeout")
	}
	fmt.Println("informers running")
	<-stopCh
	fmt.Println("informers stopped")
}

func objectName(obj interface{}) string {
	switch o := obj.(type) {
	case v1.Object:
		return o.GetName()
	case cache.DeletedFinalStateUnknown:
		if m, ok := o.Obj.(v1.Object); ok {
			return m.GetName()
		}
		return "unknown(deleted-final-state)"
	default:
		return "unknown"
	}
}
