package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "[v1] Hello, Kubernetes")
}

func main() {
	// 获取集群config路径
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()
	fmt.Sprintf("kubeconfig: %d\n", &kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// 获取集群连接客户端
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	stopCh := make(chan struct{})
	defer close(stopCh)
	// 每分钟进行resync,resync 会周期性执行list操作
	// 初始化一个共享informer
	sharedInformers := informers.NewSharedInformerFactory(clientset, time.Minute)
	podInformer := sharedInformers.Core().V1().Pods().Informer()

	podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mobj := obj.(v1.Object)
			fmt.Println("New Pod Add to Store: ", mobj.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oObj := oldObj.(v1.Object)
			nObj := newObj.(v1.Object)
			fmt.Println("Old Pod Update to New", oObj.GetName(), nObj.GetName())

		},
		DeleteFunc: func(obj interface{}) {
			mobj := obj.(v1.Object)
			fmt.Println("Pod Delete form Store: ", mobj.GetName())
		},
	})
	podInformer.Run(stopCh)

	nodeInformer := sharedInformers.Node().V1().RuntimeClasses().Informer()
	nodeInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mobj := obj.(v1.Object)
			fmt.Println("New Node Add to Store: ", mobj.GetName())
		},
		// UpdateFunc: func(oldObj, newObj interface{}) {
		// 	oObj := oldObj.(v1.Object)
		// 	nObj := newObj.(v1.Object)

		DeleteFunc: func(obj interface{}) {
			mobj := obj.(v1.Object)
			fmt.Println("Node Delete form Store: ", mobj.GetName())
		},
	})
	nodeInformer.Run(stopCh)

}
