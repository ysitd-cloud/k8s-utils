package utils

import (
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func AutoConnect() (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error
	config, err = loadInOfClusterConfig()
	if err != nil {
		config, err = loadOutOfClusterConfig()
		if err != nil {
			return nil, err
		}
	}

	return kubernetes.NewForConfig(config)
}

func loadInOfClusterConfig() (config *rest.Config, err error) {
	return rest.InClusterConfig()
}

func loadOutOfClusterConfig() (config *rest.Config, err error) {
	file := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	return clientcmd.BuildConfigFromFlags("", file)
}
