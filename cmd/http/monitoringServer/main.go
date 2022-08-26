package main

import (
	"context"
	"fmt"
	"github.com/diadata-org/diadata/http/monitoringServer/config"
	"github.com/diadata-org/diadata/http/monitoringServer/databases"
	"github.com/diadata-org/diadata/http/monitoringServer/enums"
	"github.com/diadata-org/diadata/http/monitoringServer/nodes"
	"github.com/diadata-org/diadata/http/monitoringServer/platform"
	"github.com/diadata-org/diadata/pkg/dia/helpers/db"
	"github.com/diadata-org/diadata/pkg/http/restApi"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/diadata-org/diadata/pkg/utils/probes"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"strings"
)

var startupDone = false

func main() {

	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	routerGroup := engine.Group("/monitoring")
	routerGroup.GET("/", GetAllStates)
	databases.AddRoutes(routerGroup)
	nodes.AddRoutes(routerGroup)
	platform.AddRoutes(routerGroup)
	startupDone = true
	// This environment variable is either set in docker-compose or empty

	go func() {
		err := engine.Run(utils.Getenv("LISTEN_PORT", ":8080"))
		if err != nil {
			log.Error(err)
		}
	}()

	log.Infoln("starting probes")
	probes.Start(live, ready)
}

func ready() bool {
	return startupDone
}

func live() bool {
	if !startupDone {
		return false
	}
	return config.GetKubernetesConnection() != nil
}

type MonitoringGroup struct {
	id            uuid.UUID
	groupName     string
	groupParentId uuid.UUID
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
		parentWhere = fmt.Sprintf(" where group_parent_id = %s", groupParentId)
	}
	query := fmt.Sprintf("select id, group_name from monitoring_groups %s", parentWhere)

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
			&monitoringGroup.groupParentId,
		)
		monitorState := config.GetOperationalHealthState(monitoringGroup.groupName)
		itemQuery := fmt.Sprintf("select item_name, k8s_namespace, k8s_servicename from monitoring_items WHERE monitoring_group_id = '%s' ", monitoringGroup.groupParentId.String())
		itemRows, err := conn.Query(context.Background(), itemQuery)
		if err != nil {
			log.Error("error reading endpoint from postgres", err)
			return nil
		}
		for itemRows.Next() {
			var monitoringItem MonitoringItem
			err = itemRows.Scan(
				&monitoringItem.itemName,
				&monitoringItem.k8sNamespace,
				&monitoringItem.k8sServiceName,
			)
			itemState := config.GetMajorHealthState(monitoringItem.k8sServiceName)
			listOptions := metaV1.ListOptions{}
			services, serviceErr := kube.CoreV1().Services(monitoringItem.k8sNamespace).List(context.TODO(), listOptions)
			if serviceErr != nil {
				log.Error(serviceErr)
				return nil
			}
			for _, service := range services.Items {
				if strings.Contains(service.Name, monitoringItem.k8sServiceName) {
					if len(service.Status.Conditions) == 0 {
						itemState.State = enums.HealthStateOperational
					}
					break
				}
			}
			monitorState.Subsection = append(monitorState.Subsection, itemState)
			subStates := getMonitoringGroupConfigStates(conn, monitoringGroup.id)

			subState := config.GetOperationalHealthState(monitoringGroup.groupName + " Children")
			for _, subStateItem := range subStates {
				subState.Subsection = append(subState.Subsection, subStateItem)
			}
		}
		fmt.Printf("%v", monitorState)
		defer itemRows.Close()
		states = append(states, monitorState)
	}
	defer rows.Close()
	log.Info("finished reading static endpoints")
	return
}

func mergeStateSlicesAsSubsection(name string, states []config.State) config.State {
	state := config.GetOperationalHealthState(name)

	for _, oneSlice := range states {
		state.Subsection = append(state.Subsection, oneSlice)
		switch oneSlice.State {
		case enums.HealthStateMajor:
			state.State = enums.HealthStateMajor
		case enums.HealthStateMinor:
			if state.State == enums.HealthStateOperational || state.State == enums.HealthStateMaintenance {
				state.State = enums.HealthStateMinor
			}
		case enums.HealthStateMaintenance:
			if state.State == enums.HealthStateOperational {
				state.State = enums.HealthStateOperational
			}
		}
	}

	return state
}

func GetAllStates(context *gin.Context) {
	states := getMonitoringGroupStates()
	if states == nil {
		states = []config.State{}
	}
	dbStates := databases.GetAllStates()
	states = append(states, mergeStateSlicesAsSubsection("databases", dbStates))
	nodeStates := nodes.GetAllStates()
	states = append(states, mergeStateSlicesAsSubsection("nodes", nodeStates))
	platformStates := platform.GetAllStates()
	states = append(states, mergeStateSlicesAsSubsection("platform", platformStates))

	if len(states) == 0 {
		restApi.SendError(context, http.StatusNotFound, nil)
	}
	context.JSON(http.StatusOK, states)
}
