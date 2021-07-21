package nodes

import (
	"context"
	"github.com/diadata-org/diadata/http/monitoringServer/config"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetNodeInformation() (nodeInfos []NodeInfo) {
	log.Infoln("Accessing Nodes Information")
	kube, configErr := config.GetInClusterConfig()
	if configErr != nil {
		log.Error(configErr)
	}
	if kube == nil {
		log.Error("could not get kubernetes clientSet")
	}
	nodes, err := kube.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		log.Error(err)
	}
	for i := 0; i < len(nodes.Items); i++ {
		node := nodes.Items[i]
		ready := "Terminated"
		var condition v1.NodeCondition
		for _, condition = range node.Status.Conditions {
			if condition.Type == "Ready" && condition.Status == v1.ConditionTrue {
				ready = "Running"
			}
		}
		nodeInfos = append(nodeInfos, NodeInfo{node.Name, ready})
	}
	return
}
