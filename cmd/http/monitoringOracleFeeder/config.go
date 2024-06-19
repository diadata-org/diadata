package main

import (
	"flag"
	"path/filepath"

	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var kubeConfigCache *kubernetes.Clientset

func GetKubernetesConnection() *kubernetes.Clientset {
	switch utils.Getenv("KUBERNETES_CONFIG", string(KubernetesConfigInCluster)) {
	case string(KubernetesConfigInCluster):
		return GetInClusterConfig()
	case string(KubernetesConfigFromFile):
		return GetConfigFromPath(utils.Getenv("KUBERNETES_CONFIG_PATH", ""))
	}
	return nil
}

func GetInClusterConfig() *kubernetes.Clientset {
	log.Infoln("Using in cluster config")
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil
	}
	kube, configErr := kubernetes.NewForConfig(config)
	if configErr != nil {
		log.Error(configErr)
	}
	if kube == nil {
		log.Error("could not get kubernetes clientSet")
	}
	return kube
}

func GetConfigFromPath(path string) *kubernetes.Clientset {
	if kubeConfigCache != nil {
		return kubeConfigCache
	}
	var kubeConfig *string
	if len(path) == 0 || path == "" {
		path = filepath.Join(homedir.HomeDir(), ".kube", "config")
	}
	kubeConfig = flag.String("kubeconfig", path, "(optional) absolute path to the kubeconfig file")
	flag.Parse()
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		return nil
	}
	// creates the clientset
	var configErr error
	kubeConfigCache, configErr = kubernetes.NewForConfig(config)
	if configErr != nil {
		log.Error(configErr)
	}
	if kubeConfigCache == nil {
		log.Error("could not get kubernetes clientSet")
	}
	return kubeConfigCache
}
