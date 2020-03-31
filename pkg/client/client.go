package client

import (
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	// gcp fails hard without this
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func ClientConfig() clientcmd.ClientConfig {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{})
}

// get the kube client config to call kube API
func InitClient() *kubernetes.Clientset {
	clientConfig := ClientConfig()
	config, err := clientConfig.ClientConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: getting configurations is hard\n")
		os.Exit(1)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: creating clients is hard\n")
		os.Exit(1)
	}

	return clientset
}
