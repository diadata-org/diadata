# export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=./pkg/dia/helpers/k8sbridge --go_opt=paths=source_relative --go-grpc_out=./pkg/dia/helpers/k8sbridge --go-grpc_opt=paths=source_relative ./protoc/k8sbridge.proto
