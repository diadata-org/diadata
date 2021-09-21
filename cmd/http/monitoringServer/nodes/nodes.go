package nodes

import (
	"context"
	"github.com/diadata-org/diadata/http/monitoringServer/config"
	"github.com/diadata-org/diadata/http/monitoringServer/enums"
	log "github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetAllStates() (nodeInfos []config.State) {
	log.Infoln("Accessing Nodes Information")
	kube := config.GetKubernetesConnection()
	nodes, err := kube.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		log.Error(err)
	}
	for i := 0; i < len(nodes.Items); i++ {
		node := nodes.Items[i]
		nodeInfo := config.GetMajorHealthState(node.Name)
		var condition v1.NodeCondition
		for _, condition = range node.Status.Conditions {
			if condition.Type == "Ready" && condition.Status == v1.ConditionTrue {
				nodeInfo.State = enums.HealthStateOperational
			}
		}
		nodeInfos = append(nodeInfos, nodeInfo)
	}
	return
}
