package main

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"strconv"
)

var maxReplicas int
var logLevel string
var featureFlag bool

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Read configuration from the ConfigMap
	if err := readConfig(clientset); err != nil {
		panic(err)
	}

	fmt.Printf("Max Replicas: %d\n", maxReplicas)
	fmt.Printf("Log Level: %s\n", logLevel)
	fmt.Printf("Feature Flag: %v\n", featureFlag)

	// Your controller logic goes here...

	// wait		-- Not valid anymore
	/*// Wait for termination signals
	stopCh := make(chan struct{})
	wait.Until(func() {
		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
		<-signalCh
		close(stopCh)
	}, os.Interrupt, syscall.SIGTERM)

	<-stopCh*/
}

func readConfig(clientset *kubernetes.Clientset) error {
	cm, err := clientset.CoreV1().ConfigMaps("namespace").Get("my-controller-config", metav1.GetOptions{})
	if err != nil {
		return err
	}

	maxReplicasStr := cm.Data["max_replicas"]
	logLevel = cm.Data["log_level"]
	featureFlagStr := cm.Data["feature_flag"]

	// Convert configuration values to appropriate types
	maxReplicas, _ = strconv.Atoi(maxReplicasStr)
	featureFlag, _ = strconv.ParseBool(featureFlagStr)

	return nil
}
