module github.com/diadata-org/diadata/http/monitoringServer

go 1.14

replace (
	k8s.io/api v1.5.2 => k8s.io/api v0.21.3
	k8s.io/client-go v1.5.2 => k8s.io/client-go v0.21.3
)

require (
	github.com/blang/semver v3.5.1+incompatible // indirect
	github.com/coreos/go-oidc v2.2.1+incompatible // indirect
	github.com/diadata-org/diadata v1.1.1-rc-17
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19 // indirect
	github.com/gin-gonic/gin v1.7.2
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	github.com/howeyc/gopass v0.0.0-20190910152052-7cb4b85ec19c // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/juju/ratelimit v1.0.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/swaggo/gin-swagger v1.3.0 // indirect
	golang.org/x/oauth2 v0.0.0-20210628180205-a41e5a781914 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.21.3
	k8s.io/apimachinery v0.21.3
	k8s.io/client-go v0.21.3
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920 // indirect
)
