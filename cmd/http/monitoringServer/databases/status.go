package databases

import (
	"context"
	"strings"

	"github.com/diadata-org/diadata/http/monitoringServer/config"
	"github.com/diadata-org/diadata/http/monitoringServer/enums"
	log "github.com/sirupsen/logrus"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var states []config.State

func GetAllStates() []config.State {
	states = []config.State{}

	states = append(states, InfluxState())
	states = append(states, RedisState())
	states = append(states, PostgresState())
	states = append(states, KafkaState())

	return states
}

func InfluxState() (state config.State) {
	return getServiceState("InfluxDB", "influxdb")
}

func RedisState() (state config.State) {
	return getServiceState("Redis", "dia-redis-master")
}

func PostgresState() (state config.State) {
	return getServiceState("Postgres", "dia-postgresql-ha-pgpool")
}

func KafkaState() (state config.State) {
	return getServiceState("Kafka", "kafka")
}

func getServiceState(stateName string, serviceName string) (state config.State) {
	state = config.GetMajorHealthState(stateName)
	kube := config.GetKubernetesConnection()
	listOptions := metaV1.ListOptions{}
	svcs, err := kube.CoreV1().Services("dia-db").List(context.TODO(), listOptions)
	if err != nil {
		log.Error(err)
		return
	}
	for _, service := range svcs.Items {
		if strings.Contains(service.Name, serviceName) {
			log.Infoln("service name: %v\n", service.Name)
			if len(service.Status.Conditions) == 0 {
				state.State = enums.HealthStateOperational
			}
			return
		}
	}
	return
}
