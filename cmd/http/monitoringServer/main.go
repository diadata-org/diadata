package main

import (
	"context"
	"fmt"
	"github.com/diadata-org/diadata/http/monitoringServer/stateFilter"
	"github.com/diadata-org/diadata/pkg/utils/probes"
	"net/http"
	"strings"
	"time"

	"github.com/diadata-org/diadata/http/monitoringServer/config"
	"github.com/diadata-org/diadata/http/monitoringServer/databases"
	"github.com/diadata-org/diadata/http/monitoringServer/enums"
	"github.com/diadata-org/diadata/http/monitoringServer/nodes"
	"github.com/diadata-org/diadata/http/monitoringServer/platform"
	"github.com/diadata-org/diadata/pkg/dia/helpers/db"
	"github.com/diadata-org/diadata/pkg/http/restApi"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	StartupDone      = false
	CacheGlobalState []config.State
)

const CacheTTLSeconds = 5 * 60

func main() {
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	routerGroup := engine.Group("/monitoring")
	routerGroup.GET("/", GetAllStates)
	databases.AddRoutes(routerGroup)
	nodes.AddRoutes(routerGroup)
	platform.AddRoutes(routerGroup)
	readAllStates()

	log.Infoln("starting probes")
	StartupDone = true
	probes.Start(live, ready)

	err := engine.Run(utils.Getenv("LISTEN_PORT", ":8080"))
	if err != nil {
		log.Error(err)
	}
	ticker := time.NewTicker(time.Second * CacheTTLSeconds)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				readAllStates()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func ready() bool {
	return StartupDone
}

func live() bool {
	if !StartupDone {
		return false
	}
	return config.GetKubernetesConnection() != nil
}

type MonitoringGroup struct {
	id                   uuid.UUID
	groupName            string
	operationalThreshold float64
	minorThreshold       float64
	groupParentId        uuid.UUID
}

type MonitoringItem struct {
	itemName       string
	k8sNamespace   string
	k8sServiceName string
}

func getMonitoringGroupStates() []config.State {
	return getMonitoringGroupConfigStates(db.PostgresDatabase(), uuid.Nil)
}

func getMonitoringGroupConfigStates(conn *pgxpool.Pool, groupParentId uuid.UUID) (states []config.State) {
	parentWhere := ""
	if groupParentId != uuid.Nil {
		parentWhere = fmt.Sprintf(" and group_parent_id = %s", groupParentId)
	}
	//goland:noinspection SqlResolve
	query := fmt.Sprintf("select id, group_name, operational_threshold, minor_threshold from monitoring_groups where active = true %s", parentWhere)

	log.Info("reading service monitoring endpoints")
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Error("error reading endpoint from postgres", err)
		return nil
	}

	kube := config.GetKubernetesConnection()

	for rows.Next() {
		var monitoringGroup MonitoringGroup
		err := rows.Scan(
			&monitoringGroup.id,
			&monitoringGroup.groupName,
			&monitoringGroup.operationalThreshold,
			&monitoringGroup.minorThreshold,
		)
		monitorState := config.GetOperationalHealthState(monitoringGroup.groupName)
		itemQuery := fmt.Sprintf("select item_name, k8s_namespace, k8s_servicename from monitoring_items WHERE monitoring_group_id = '%s' AND active = true", monitoringGroup.id.String())
		itemRows, err := conn.Query(context.Background(), itemQuery)
		log.Info("Reading Group: ", monitoringGroup.groupName)
		if err != nil {
			log.Error("error reading endpoint from postgres ", err)
			return nil
		}
		for itemRows.Next() {
			var monitoringItem MonitoringItem
			err = itemRows.Scan(
				&monitoringItem.itemName,
				&monitoringItem.k8sNamespace,
				&monitoringItem.k8sServiceName,
			)
			itemState := config.GetOperationalHealthState(monitoringItem.itemName)
			listOptions := metaV1.ListOptions{
				LabelSelector: "app=" + monitoringItem.k8sServiceName,
			}
			services, serviceErr := kube.CoreV1().Services(monitoringItem.k8sNamespace).List(context.TODO(), listOptions)
			if serviceErr != nil {
				log.Error(serviceErr)
				return nil
			}
			for _, service := range services.Items {
				if strings.Contains(service.Name, monitoringItem.k8sServiceName) {
					if len(service.Status.Conditions) == 0 {
						itemState.State = enums.HealthStateOperational
					} else {
						itemState.State = enums.HealthStateMinor
					}
					break
				}
			}
			deployments, deploymentErr := kube.AppsV1().Deployments(monitoringItem.k8sNamespace).List(context.TODO(), listOptions)
			if deploymentErr != nil {
				log.Error(deploymentErr)
				return nil
			}
			for _, deployment := range deployments.Items {
				if strings.Contains(deployment.Name, monitoringItem.k8sServiceName) {
					for _, condition := range deployment.Status.Conditions {
						if condition.Status == "True" && (condition.Type == "Progressing" || condition.Type == "Available") {
							itemState.State = enums.HealthStateOperational
						} else {
							log.Warning("Condition could not be checked for deployment: " + monitoringItem.k8sServiceName + " output of conditions was: ")
							log.Warning(condition)
							itemState.State = enums.HealthStateMajor
							break
						}
					}
					break
				}
			}
			monitorState.Subsection = append(monitorState.Subsection, itemState)
		}
		//goland:noinspection GoDeferInLoop
		defer itemRows.Close()
		monitorState.State = stateFilter.CheckHealthState(monitorState.Subsection, monitoringGroup.operationalThreshold, monitoringGroup.minorThreshold)
		states = append(states, monitorState)
	}
	defer rows.Close()
	log.Info("finished reading static endpoints")
	return
}

func mergeStateSlicesAsSubsection(name string, states []config.State) (configState config.State) {
	configState = config.GetOperationalHealthState(name)

	for _, oneSlice := range states {
		configState.Subsection = append(configState.Subsection, oneSlice)
		configState.State = enums.CompareStates(oneSlice.State, configState.State)
	}

	return
}

func GetAllStates(context *gin.Context) {
	if len(CacheGlobalState) == 0 {
		restApi.SendError(context, http.StatusNotFound, nil)
	}
	context.JSON(http.StatusOK, CacheGlobalState)
}

func readAllStates() {
	states := getMonitoringGroupStates()
	if states == nil {
		states = []config.State{}
	}
	dbStates := databases.GetAllStates()
	states = append(states, mergeStateSlicesAsSubsection("Databases", dbStates))
	nodeStates := nodes.GetAllStates()
	states = append(states, mergeStateSlicesAsSubsection("Nodes", nodeStates))
	platformStates := platform.GetAllStates()
	states = append(states, mergeStateSlicesAsSubsection("Platform", platformStates))

	CacheGlobalState = states
}
