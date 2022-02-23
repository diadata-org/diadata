module github.com/diadata-org/diadata/http/monitoringServer

go 1.14

replace (
	k8s.io/api v1.5.2 => k8s.io/api v0.21.3
	k8s.io/client-go v1.5.2 => k8s.io/client-go v0.21.3
)

require (
	github.com/diadata-org/diadata v1.3.4
	github.com/gin-gonic/gin v1.7.2
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/sirupsen/logrus v1.8.1
	k8s.io/api v0.21.3
	k8s.io/apimachinery v0.21.3
	k8s.io/client-go v0.21.3
)
