dockerfile="Dockerfile-synthScraper"
imageName="synthscraper"

#  image: "us.icr.io/dia-registry/services/pairdiscoveryservice"


#   image: "us.icr.io/dia-registry/services/assetcollectionservice"

#   image: "us.icr.io/dia-registry/scrapers/genericcollector"
# type="services"

# dockerfile="Dockerfile-pairDiscoveryService"
# imageName="pairdiscoveryservice"
type="scrapers"
version="v1.4.1-rc-376"

build_and_push() {
        docker build -f "build/$1" -t "diadata.$2" .
        docker tag "diadata.$2" "us.icr.io/dia-registry/$3/$2:latest"
        docker push "us.icr.io/dia-registry/$3/$2:latest"

        docker tag "diadata.$2" "us.icr.io/dia-registry/$3/$2:$version"
        docker push "us.icr.io/dia-registry/$3/$2:$version"
}

build_and_push "$dockerfile" "$imageName" "$type"


