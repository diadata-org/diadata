package main

import (
	"context"
	"fmt"
	"github.com/diadata-org/diadata/pkg/http/restApi"
	"github.com/diadata-org/diadata/pkg/utils/probes"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	StartupDone      = false
	CacheGlobalState []k8sService
)

type k8sService struct {
	namespace   string
	deployments []appsv1.Deployment
}

const CacheTTLSeconds = 5 * 60

func main() {
	engine := gin.Default()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	routerGroup := engine.Group(utils.Getenv("URL_FOLDER_PREFIX", "/monitoringNs"))
	routerGroup.GET("/:name", GetAllStates)
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
	return GetKubernetesConnection() != nil
}

func readAllStates() {
	log.Info("reading namespaces services")

	kube := GetKubernetesConnection()
	// Fetch all namespaces
	namespaces, err := kube.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error listing namespaces: %s\n", err)
		os.Exit(1)
	}

	var k8sServices []k8sService
	for _, ns := range namespaces.Items {
		fmt.Printf("listing services in namespace %s\n", ns.Name)
		var namespaceServices = k8sService{namespace: ns.Name}
		var deploymentItems []appsv1.Deployment
		deployments, err := kube.AppsV1().Deployments(ns.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("Error listing services in namespace %s: %s\n", ns.Name, err)
			continue
		}
		for _, service := range deployments.Items {
			deploymentItems = append(deploymentItems, service)
		}
		namespaceServices.deployments = deploymentItems
		fmt.Printf("listing services in namespace %s: %s\n", ns.Name, strconv.Itoa(len(deploymentItems)))
		k8sServices = append(k8sServices, namespaceServices)
	}

	CacheGlobalState = k8sServices
	log.Info("finished reading static endpoints")
	return
}

func GetAllStates(context *gin.Context) {
	if len(CacheGlobalState) > 0 {
		namespace := context.Param("name")
		for _, state := range CacheGlobalState {
			if state.namespace == namespace {
				context.JSON(http.StatusOK, state.deployments)
			}
		}
	}
	restApi.SendError(context, http.StatusNotFound, nil)
}
