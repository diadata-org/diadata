cd $GOPATH/src/github.com/diadata-org/api-golang
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color
go get -t -v ./...
go test -v ./... 
if [ $? -ne 0 ]
then
echo -e "${RED}Fail${NC}\n"
else
echo -e  "${GREEN}OK${NC}\n"
fi

