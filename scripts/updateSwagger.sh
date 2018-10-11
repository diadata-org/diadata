# https://github.com/swaggo/swag
cd $GOPATH/src/github.com/diadata-org/diadata
swag init -g ./internal/pkg/http/restServer/restServer.go 
rm -rf api/docs
mv docs api/docs
cd -
