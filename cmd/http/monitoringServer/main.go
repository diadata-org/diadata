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
	"time"
)

var StartupDone = false
var CacheGlobalState []config.State

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

	log.Infoln("starting probes")
	probes.Start(live, ready)

	err := engine.Run(utils.Getenv("LISTEN_PORT", ":8080"))
	StartupDone = true
	if err != nil {
		log.Error(err)
	}
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
	id            uuid.UUID
	groupName     string
	groupParentId uuid.UUID
	sectionType   enums.SectionType
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
	query := fmt.Sprintf("select id, group_name, section_type from monitoring_groups where active = true %s", parentWhere)

	log.Info("reading service monitoring endpoints")
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Error("error reading endpoint from postgres", err)
		return nil
	}

	kube := config.GetKubernetesConnection()

	for rows.Next() {
		var monitoringGroup MonitoringGroup
		var section string
		err := rows.Scan(
			&monitoringGroup.id,
			&monitoringGroup.groupName,
			&section,
		)
		if err != nil {
			log.Error("error scanning row from postgres ", err)
			return nil
		}
		monitoringGroup.sectionType, err = enums.GetSectionTypeFromString(section)
		if err != nil {
			log.Error("error getting section from postgres ", err)
			return nil
		}
		monitorState := config.GetOperationalHealthStateWithSection(monitoringGroup.groupName, monitoringGroup.sectionType)
		itemQuery := fmt.Sprintf("select item_name, k8s_namespace, k8s_servicename from monitoring_items WHERE monitoring_group_id = '%s' AND active = true", monitoringGroup.id.String())
		itemRows, err := conn.Query(context.Background(), itemQuery)
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
			itemState := config.GetOperationalHealthStateWithSection(monitoringItem.itemName, monitoringGroup.sectionType)
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
							itemState.State = enums.HealthStateMajor
							break
						}
					}
					break
				}
			}
			monitorState.Subsection = append(monitorState.Subsection, itemState)
			/*subStates := getMonitoringGroupConfigStates(conn, monitoringGroup.id)

			subState := config.GetOperationalHealthState(monitoringGroup.groupName + " Children")
			for _, subStateItem := range subStates {
				subState.Subsection = append(subState.Subsection, subStateItem)
			}
			*/
		}
		//goland:noinspection GoDeferInLoop
		defer itemRows.Close()
		for _, item := range monitorState.Subsection {
			monitorState.State = enums.CompareStates(item.State, monitorState.State)
		}
		states = append(states, monitorState)
	}
	defer rows.Close()
	log.Info("finished reading static endpoints")
	return
}

func mergeStateSlicesAsSubsection(name string, states []config.State) (configState config.State) {

	for _, oneSlice := range states {
		configState = config.GetOperationalHealthStateWithSection(name, oneSlice.SectionType)
		configState.Subsection = append(configState.Subsection, oneSlice)
	}

	switch configState.SectionType {
	case enums.DatabaseSection:
		configState.State = appendDatabaseSlicesStateLogic(states)
		break
	case enums.ScraperSection:
		configState.State = appendScraperSlicesStateLogic(states)
		break
	case enums.ServiceSection:
		configState.State = appendServiceSlicesStateLogic(states)
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
	states = append(states, mergeStateSlicesAsSubsection("databases", dbStates))
	nodeStates := nodes.GetAllStates()
	states = append(states, mergeStateSlicesAsSubsection("nodes", nodeStates))
	platformStates := platform.GetAllStates()
	states = append(states, mergeStateSlicesAsSubsection("platform", platformStates))

	CacheGlobalState = states
}

func countHealthyStates(states []config.State) (count int) {
	count = 0
	for _, state := range states {
		if state.State == enums.HealthStateOperational {
			count++
		}
	}
	return
}

func appendScraperSlicesStateLogic(states []config.State) enums.HealthSate {
	operational := float64(countHealthyStates(states)) / float64(len(states))
	if operational >= 0.8 {
		return enums.HealthStateOperational
	} else if operational >= 0.5 {
		return enums.HealthStateMinor
	} else if operational <= 0.5 {
		return enums.HealthStateMajor
	}
	return enums.HealthStateMaintenance
}

func appendServiceSlicesStateLogic(states []config.State) enums.HealthSate {
	operational := countHealthyStates(states)
	all := len(states)
	if operational == len(states) {
		return enums.HealthStateOperational
	} else if (all - operational) == 1 {
		return enums.HealthStateMinor
	} else if (all - operational) > 1 {
		return enums.HealthStateMajor
	}
	return enums.HealthStateMaintenance
}

func appendDatabaseSlicesStateLogic(states []config.State) enums.HealthSate {
	operational := countHealthyStates(states)
	all := len(states)
	if operational == len(states) {
		return enums.HealthStateOperational
	} else if (all - operational) == 1 {
		return enums.HealthStateMinor
	} else if (all - operational) > 1 {
		return enums.HealthStateMajor
	}
	return enums.HealthStateMaintenance
}
