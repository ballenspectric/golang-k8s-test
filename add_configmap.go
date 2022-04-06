package main

import (
	"context"
	"log"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	kubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")
	log.Println(kubeconfigPath)

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)

	if err != nil {
		log.Fatalf(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.Fatalf(err.Error())
	}

	newConfigMap := v1.ConfigMap{
		Data: map[string]string{"how": "with golang", "why": "to learn"},
	}

	newConfigMap.Name = "golang-configmap"
	newConfigMap.Namespace = "hello"

	_, err = clientset.CoreV1().ConfigMaps("hello").Create(
		context.TODO(),
		&newConfigMap,
		metav1.CreateOptions{},
	)

	if err != nil {
		log.Fatalf(err.Error())
	}
}
