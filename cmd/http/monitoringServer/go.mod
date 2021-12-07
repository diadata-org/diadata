module github.com/diadata-org/diadata/http/monitoringServer

go 1.14

replace (
	k8s.io/api v1.5.2 => k8s.io/api v0.21.3
	k8s.io/client-go v1.5.2 => k8s.io/client-go v0.21.3
)

require (
	github.com/diadata-org/diadata v1.3.0
	github.com/gin-gonic/gin v1.7.2
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/oauth2 v0.0.0-20210628180205-a41e5a781914 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.21.3
	k8s.io/apimachinery v0.21.3
	k8s.io/client-go v0.21.3
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920 // indirect
)
