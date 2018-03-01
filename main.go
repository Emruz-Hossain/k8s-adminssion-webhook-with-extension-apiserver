package main

import (
	//"k8s.io/code-generator"
	"flag"

	"github.com/appscode/go/log"
	"github.com/emruz-hossain/k8s-admission-webhook-with-extension-apiserver/controller"

	clientset "github.com/emruz-hossain/k8s-admission-webhook-with-extension-apiserver/client/clientset/versioned"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	masterURL  string
	kubeConfig string
)

func init() {
	flag.StringVar(&kubeConfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
func main() {
	flag.Parse()

	stopCh := make(chan struct{})
	defer close(stopCh)

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeConfig)
	if err != nil {
		log.Fatalf("Can't build config. Reason:  ", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Can't build kubeClient. Reason:", err.Error())
	}

	kubecarClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		log.Fatalln("Can't build kubecar client. Reason: ", err.Error())
	}
	options := controller.NewOptions()
	kubecarController := controller.NewKubecarController(kubeClient, kubecarClient, *options)

	go func() {
		log.Info("Starting controller....")
		kubecarController.Run(stopCh)
	}()

	select {}
}
