package main

import (
	"context"
	"encoding/json"
	"log"
	"path/filepath"

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

	pods, err := clientset.CoreV1().Pods("hello").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, pod := range pods.Items {
		j, err := json.MarshalIndent(pod.ObjectMeta.Name, "", "  ")

		if err != nil {
			log.Fatalf(err.Error())
		}

		log.Printf(string(j))
	}
}
