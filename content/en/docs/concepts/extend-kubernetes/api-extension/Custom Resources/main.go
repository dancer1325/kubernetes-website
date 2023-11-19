package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	// Set up kubeconfig
	kubeconfig := flag.String("kubeconfig", getDefaultKubeconfig(), "path to the kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("Error building kubeconfig: %v\n", err)
		os.Exit(1)
	}

	// Create Kubernetes clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating clientset: %v\n", err)
		os.Exit(1)
	}

	// Define the Pod
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "nginx-pod",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx-container",
					Image: "nginx:latest",
				},
			},
		},
	}

	// Create the Pod
	createdPod, err := clientset.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error creating Pod: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Pod %s created successfully\n", createdPod.Name)
}

func getDefaultKubeconfig() string {
	home := homedir.HomeDir()
	return home + "/.kube/config"
}
