package main

type KubernetesConfig string

const (
	KubernetesConfigInCluster KubernetesConfig = "IN_CLUSTER"
	KubernetesConfigFromFile  KubernetesConfig = "FROM_FILE"
)
