#!/usr/bin/env bash
#
# Description: This script is used to manage the DIA development environment. Helpful on start/stop the platform, and create/remove resources
# Vesion: 0.1.0

SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ]; do
    SCRIPT_FOLDER="$(cd -P "$(dirname "$SOURCE")" >/dev/null 2>&1 && pwd)"
    SOURCE="$(readlink "$SOURCE")"
    [[ $SOURCE != /* ]] && SOURCE="$SCRIPT_FOLDER/$SOURCE"
done
SCRIPT_FOLDER="$(cd -P "$(dirname "$SOURCE")" >/dev/null 2>&1 && pwd)"
unset SOURCE

function usage() {
    echo "DIA Test Space"
    echo
    echo "Usage:"
    echo "  env [OPTIONS...] COMMAND [ARGS...]"
    echo
    echo "Options:"
    echo "  -h --help                 Print help"
    echo "  -v --verbose              Verbose mode on"
    echo "  -c --cpus <num>           Define the num of virtual CPUs (eg: 8)"
    echo "     --full                 Full mode on"
    echo "     --single               Single mode on"
    echo
    echo "Available commands:"
    echo "  start                     Start cluster"
    echo "  stop                      Stop the cluster"
    echo "  delete                    Delete all cluster resources"
    echo "  build                     Build DIA platform and prepare container images"
    echo "  install                   Install DIA platform"
    echo "  uninstall                 Un-install DIA platform"
    echo "  snapshot                  Pull and import snapshot of data"
    echo "  create <resource>         Create a new resource"
    echo "    scraper-cex [id]          Create a CEX scraper"
    echo "    scraper-dex [id]          Create a DEX scraper"
    echo "    scraper-foreign [id]      Create a foreign scraper"
    echo "    scraper-liquidity [id]    Create a liquidity scraper"
    echo "    cronjob                   Create a cronjob for snapshots"
    echo "    demos-scraper           Create a demo with all scrapers"
    echo "  remove                      Remove a resource"
    echo "    scraper-cex [id]          Remove a CEX scraper"
    echo "    scraper-dex [id]          Remove a DEX scraper"
    echo "    scraper-foreign [id]      Remove a foreign scraper"
    echo "    scraper-liquidity [id]    Remove a liquidity scraper"
    echo "    cronjob                   Remove a cronjob for snapshots"
    echo "    demos-scraper             Remove a demo with all scrapers"
    echo "  shell                     Connect to enviornment shell"
    echo "  logs                      Print logs"
    echo "  clean                     Clean unused files"
    echo "  info                      Show detailed information"
    echo "  ping                      Make ping tests"
    echo "  data                      List data"
    echo "    exchanges               List exchanges"
    echo "    exchange                Show exchange metadata"
    echo "    blockchain              Show blockchain metadata"
    echo "    scraper-cex             Show CEX scraper metadata"
    echo "    scraper-dex             Show DEX scraper metadata"
    echo "    scraper-liquidity       Show liquidity scraper metadata"
    echo "    scraper-foreign         Show foreign scraper metadata"
    echo "  data-reset                Reset data"
    echo
    echo "Report bugs to: <https://github.com/diadata-org/diadata/issues>"
}

function _image_exist(){
    minikube_docker_query="$(minikube -p "${minikube_profile}" ssh -- "docker images -q $1:latest" 2> /dev/null)"
    if [[ "${minikube_docker_query}" == "" ]]; then
        echo "Image $1 is not present"
        return 1
    else
        echo "Image $1 is present"
        return 0
    fi
}

function _build_ifnotexist(){
    if [ "$arg_single_mode" = true ]; then
        docker_query="$(docker images -q "$2:latest" 2> /dev/null)"
        if [[ "${docker_query}" == "" ]]; then
            echo "Image $2 is not present, building ..."
            docker buildx build -f "build/$1" -t "$2:latest" .
        else
            echo "Image $2 is present"
        fi
    else
        minikube_docker_query="$(minikube -p "${minikube_profile}" ssh -- "docker images -q $2:latest" 2> /dev/null)"
        if [[ "${minikube_docker_query}" == "" ]]; then
            echo "Image $2 is not present, building ..."
            if [[ "$minikube_driver" == "docker" ]]; then
                eval "$(minikube -p "${minikube_profile}" docker-env)"
                docker buildx build -f "build/$1" -t "$2:latest" .
                eval "$(minikube -p "${minikube_profile}" docker-env --unset)"
            else
                minikube -p "${minikube_profile}" image build -f "build/$1" -t "$2:latest" .
            fi
        else
            echo "Image $2 is present"
        fi
    fi
}

function _minikube_profile_isrunning() {
    minikube status --profile "$1" --format '{{.Host}}' >/dev/null 2>&1 || return 1
    minikube status --profile "$1" --format '{{.Kubelet}}' >/dev/null 2>&1 || return 1
    minikube status --profile "$1" --format '{{.APIServer}}' >/dev/null 2>&1 || return 1
    return 0
}

function _info() {
    echo "DIA Version: $version_detected"
    echo "Print Working Directory: $PWD"
    echo "Script Directory: $SCRIPT_FOLDER"
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        echo "Operating System: GNU/Linux"
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        echo "Operating System: Mac OSX"
    fi
    if [ "$arg_single_mode" = true ]; then
        echo "Docker version: v$(docker version --format '{{.Client.Version}}')"
    else
        echo "Minikube Version: $(minikube version --short)"
        if [[ "$minikube_driver" == "docker" ]]; then
            echo "Minikube Driver: ${minikube_driver} v$(docker version --format '{{.Client.Version}}')"
        else
            echo "Minikube Driver: ${minikube_driver}"
        fi
        echo "Minikube K8s: ${minikube_k8s_version}"
        echo "Minikube CPUs: ${minikube_hw_cpus}"
        echo "Minikube Memory: ${minikube_hw_ram}"
        echo "Minikube Disk: ${minikube_hw_disk}"
        echo "Minikube Profile: ${minikube_profile}"
    fi
}

function command_exists() {
  command -v "$1" >/dev/null 2>&1
}

function _status_host(){
    echo "Uptime: $(uptime -p)"
}

function main() {
    # Arguments parsing
    local command=("$@")

    local args=
    args=$(getopt -o hnsvc: --long cpus:,disk-space:,nodetached,full,single,verbose,help -n './testenv.sh' -- "$@")
    # shellcheck disable=SC2181
    if [ $? != 0 ]; then
        usage 1>&2
        return 1
    fi
    set -euo pipefail
    eval set -- "$args"
    local arg_full_mode=false
    local arg_single_mode=false
    local arg_verbose_mode=false
    local arg_cpus=""
    while true; do
        case "$1" in
            --verbose | -v)
                arg_verbose_mode=true
                shift ;;
            --single)
                arg_single_mode=true
                shift ;;
            --full | -f)
                arg_full_mode=true
                shift ;;
            --cpus | -c)
                arg_cpus="$2"
                if [ "$arg_cpus" = "" ]; then
                    echo "Invalid parameter" >&2
                    return 1
                fi
                shift 2 ;;
            --help | -h     ) usage; return 0 ;;
            --              ) shift; break    ;;
            *               )        break    ;;
        esac
    done


    # Check dependencies
    if ! hash git 2> /dev/null; then echo "Git not found" >&2; return 1; fi
    if ! hash docker 2> /dev/null; then echo "Docker not found" >&2; return 1; fi
    if ! hash minikube 2> /dev/null; then echo "Minikube not found" >&2; return 1; fi

    # Safe checks
    if [[ "$OSTYPE" != "linux-gnu"* ]] && [[ "$OSTYPE" != "darwin"* ]]; then
        echo "Unknown operating system"
        exit 1
    fi
    if [ -z "${command[0]:-}" ]; then
        echo "No command specified" >&2
        exit 1
    fi

    # Variables
    local minikube_profile=dia
    local minikube_k8s_version=v1.25.7
    local minikube_hw_cpus=4
    local minikube_hw_ram=8g
    local minikube_hw_disk=50g
    local minikube_driver=docker
    local snapshot_docker_registry=https://registry.hub.docker.com/v2/
    local snapshot_docker_username=dia_contributor
    local snapshot_docker_password=dia_contributor_pw
    local snapshot_docker_email=dia_contributor@example.com
    local data_docker_registry=docker.io
    local data_docker_username=dia
    local data_docker_password=dia_pw
    local data_docker_email=dia@diadata.com
    # shellcheck disable=SC1091
    # TODO: this will break the script?
    if [ -e testenv.local ]; then
    source ./testspace/.testenv.local
    fi
    if [[ "$arg_cpus" != "" ]]; then
        minikube_hw_cpus="${arg_cpus}"
    fi
    local version_detected
    version_detected=$(git describe --tags --abbrev=0)
    declare -a demos_scraper_cex=("bitfinex" "bittrex" "coinbase" "mexc")
    declare -a demos_scraper_dex=("platypus" "orca")
    declare -a demos_scraper_liquidity=("platypus" "orca")
    declare -a demos_scraper_foreign=("yahoofinance")

    # Command
    case "${command[0]}" in
    start)
      echo "running start"
        if [ "${#command[@]}" -eq 1 ]; then
            if [ "$arg_verbose_mode" = true ]; then _info; fi
            if [ "$arg_single_mode" = true ]; then
                echo "WIP"
            else
                if ! _minikube_profile_isrunning "${minikube_profile}"; then
                    echo "Starting cluster ..."
                    minikube --profile "${minikube_profile}" start \
                        --kubernetes-version "${minikube_k8s_version}" \
                        --driver "${minikube_driver}" \
                        --cpus "${minikube_hw_cpus}" \
                        --memory "${minikube_hw_ram}" \
                        --disk-size "${minikube_hw_disk}" \
                        --mount-string="$(pwd):/mnt/diadata" \
                        --mount
                    echo "Minikube cluster started with success"
                    exit 0
                else
                    echo "Cluster is already running"
                    exit 1
                fi
            fi
        else
            echo "Unknown command" >&2
            exit 1
        fi
        ;;
    stop)
        if [ "${#command[@]}" -eq 1 ]; then
            echo "Stopping cluster ..."
            minikube -p "${minikube_profile}" stop
        else
            echo "Unknown command" >&2
            exit 1
        fi
        ;;
    delete)
        if [ "${#command[@]}" -eq 1 ]; then
            echo "Deleting cluster ..."
            minikube delete -p "${minikube_profile}"
        else
            echo "Unknown command" >&2
            exit 1
        fi
        ;;
    build)
        if [ "${#command[@]}" -eq 1 ]; then
            echo "Building images ..."
            _build_ifnotexist build/Dockerfile-DiadataBuild-114-Dev dia.build-114.dev
            _build_ifnotexist build/Dockerfile-DiadataBuild-117-Dev dia.build-117.dev
            _build_ifnotexist build/Dockerfile-DiadataBuild-119-Dev dia.build-119.dev
            _build_ifnotexist build/Dockerfile-DiadataBuild-120-Dev dia.build-120.dev
            _build_ifnotexist Dockerfile-filtersBlockService-Dev dia.filtersblockservice.dev
            _build_ifnotexist Dockerfile-tradesBlockService-Dev dia.tradesblockservice.dev
            _build_ifnotexist Dockerfile-pairDiscoveryService-Dev dia.pairdiscoveryservice.dev
            _build_ifnotexist Dockerfile-genericCollector-Dev dia.genericcollector.dev
            _build_ifnotexist Dockerfile-genericForeignScraper-Dev dia.genericforeignscraper.dev
            _build_ifnotexist Dockerfile-liquidityScraper-Dev dia.liquidityscraper.dev
            _build_ifnotexist Dockerfile-assetCollectionService-Dev dia.assetcollectionservice.dev
            if [ "$arg_full_mode" = true ]; then
                _build_ifnotexist Dockerfile-blockchainservice-Dev dia.blockchainservice.dev
            fi
        else
            echo "Unknown command" >&2
            exit 1
        fi
        ;;
    install)
        if [ "${#command[@]}" -eq 1 ]; then
            _image_exist dia.build-114.dev || exit 1
            _image_exist dia.build-117.dev || exit 1
            _image_exist dia.build-119.dev || exit 1
            _image_exist dia.build-120.dev || exit 1
            _image_exist dia.filtersblockservice.dev || exit 1
            _image_exist dia.tradesblockservice.dev || exit 1
            _image_exist dia.pairdiscoveryservice.dev || exit 1
            _image_exist dia.genericcollector.dev || exit 1
            _image_exist dia.genericforeignscraper.dev || exit 1
            _image_exist dia.liquidityscraper.dev || exit 1
            _image_exist dia.assetcollectionservice.dev || exit 1
            if [ "$arg_full_mode" = true ]; then
                _image_exist dia.blockchainservice.dev || exit 1
            fi
            echo "Installing services ..."
            minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/service-filtersblockservice.yaml
            minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/service-tradesblockservice.yaml
            minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/data-kafka.yaml
            minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/data-redis.yaml
            minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/data-influx.yaml
            if [ "$arg_full_mode" = true ]; then
                minikube -p "${minikube_profile}" kubectl -- create configmap postgres-schemma --from-file=deployments/config/pginit.sql
                minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/data-postgres.yaml
                minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/job-prepare.yaml
            else
                minikube -p "${minikube_profile}" kubectl -- \
                    create secret docker-registry regcred-read \
                    --docker-server="${data_docker_registry}" \
                    --docker-username="${data_docker_username}" \
                    --docker-password="${data_docker_password}" \
                    --docker-email="${data_docker_email}"
                minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/data-postgres-prepopulated.yaml
            fi
            echo "Installation ended successfully"
        else
            echo "Unknown command" >&2
            exit 1
        fi
        ;;
    uninstall)
        if [ "${#command[@]}" -eq 1 ]; then
            echo "Uninstalling services ..."
            minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/service-filtersblockservice.yaml || true
            minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/service-tradesblockservice.yaml || true
            minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/data-kafka.yaml || true
            minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/data-redis.yaml || true
            minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/data-influx.yaml || true
            if [ "$arg_full_mode" = true ]; then
                minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/data-postgres.yaml || true
                minikube -p "${minikube_profile}" kubectl -- delete configmap postgres-schemma || true
            else
                minikube -p "${minikube_profile}" kubectl -- delete secret regcred-read || true
                minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/data-postgres-prepopulated.yaml || true
            fi
            echo "Services uninstalled with success"
        else
            echo "Unknown command" >&2
            exit 1
        fi
        ;;
    snapshot)
        if [ "${#command[@]}" -gt 1 ]; then
        # Check if unzip, wget, and psql commands exist
        if ! command_exists unzip; then
          echo "Error: 'unzip' command not found. Please install unzip and try again."
          exit 1
        fi

        if ! command_exists wget; then
          echo "Error: 'wget' command not found. Please install wget and try again."
          exit 1
        fi

        if ! command_exists psql; then
          echo "Error: 'psql' command not found. Please install PostgreSQL client and try again."
          exit 1
        fi

        # Ask for variables with _EXTRACT and export them
        read -p "Enter Postgres Server: " PGHOST_EXTRACT
        read -p "Enter Postgres Port: " PGPORT_EXTRACT
        read -p "Enter Postgres User: " PGUSER_EXTRACT
        read -p "Enter Postgres Password: " PGPASSWORD_EXTRACT
        read -p "Enter Postgres Database: " PGDB_EXTRACT

        export PGHOST=${PGHOST_EXTRACT}
        export PGUSER=${PGUSER_EXTRACT}
        export PGDB=${PGDB_EXTRACT}
        export PGPASSWORD=${PGPASSWORD_EXTRACT}
        export PGPORT=${PGPORT_EXTRACT}

        # Download and unzip the file from rest.diadata.org
        wget -qO snapshot.zip http://rest.diadata.org/snapshot/latest
        unzip -o snapshot.zip

        # Run the psql commands
        psql --port ${PGPORT} --username ${PGUSER} --dbname ${PGDB} --file output-blockchain.sql
        psql --port ${PGPORT} --username ${PGUSER} --dbname ${PGDB} --file output-exchange.sql
        psql --port ${PGPORT} --username ${PGUSER} --dbname ${PGDB} --file output-asset.sql
        psql --port ${PGPORT} --username ${PGUSER} --dbname ${PGDB} --file output-pool.sql
        psql --port ${PGPORT} --username ${PGUSER} --dbname ${PGDB} --file output-poolasset.sql
        psql --port ${PGPORT} --username ${PGUSER} --dbname ${PGDB} --file output-exchangepair.sql

        # Delete the output-*.sql files
        rm output-blockchain.sql

        echo "Data import and processing completed successfully."
      fi
      ;;
    create)
        if [ "${#command[@]}" -gt 1 ]; then
            case "${command[1]}" in
            scraper-cex)
                if [ "${#command[@]}" -gt 2 ]; then
                    if [[ " ${demos_scraper_cex[*]} " == *" ${command[2]} "* ]]; then
                        minikube -p "${minikube_profile}" kubectl -- create -f "deployments/k8s-yaml/scraper-cex-${command[2]}.yaml"
                    else
                        echo "Unknown demo" >&2
                        exit 1
                    fi
                else
                    minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/scraper-cex.yaml
                fi
                ;;
            scraper-dex)
                if [ "${#command[@]}" -gt 2 ]; then
                    if [[ " ${demos_scraper_dex[*]} " == *" ${command[2]} "* ]]; then
                        minikube -p "${minikube_profile}" kubectl -- create -f "deployments/k8s-yaml/scraper-dex-${command[2]}.yaml"
                    else
                        echo "Unknown demo" >&2
                        exit 1
                    fi
                else
                    minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/scraper-dex.yaml
                fi
                ;;
            scraper-foreign)
                if [ "${#command[@]}" -gt 2 ]; then
                    if [[ " ${demos_scraper_foreign[*]} " == *" ${command[2]} "* ]]; then
                        minikube -p "${minikube_profile}" kubectl -- create -f "deployments/k8s-yaml/scraper-foreign-${command[2]}.yaml"
                    else
                        echo "Unknown demo" >&2
                        exit 1
                    fi
                else
                    minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/scraper-foreign.yaml
                fi
                ;;
            scraper-liquidity)
                if [ "${#command[@]}" -gt 2 ]; then
                    if [[ " ${demos_scraper_liquidity[*]} " == *" ${command[2]} "* ]]; then
                        minikube -p "${minikube_profile}" kubectl -- create -f "deployments/k8s-yaml/scraper-liquidity-${command[2]}.yaml"
                    else
                        echo "Unknown demo" >&2
                        exit 1
                    fi
                else
                    minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/scraper-liquidity.yaml
                fi
                ;;
            cronjob)
                if [ "${#command[@]}" -eq 2 ]; then
                    minikube -p "${minikube_profile}" kubectl -- create configmap postgres-entrypoint --from-file=deployments/config/postgres-docker-entrypoint.sh
                    minikube -p "${minikube_profile}" kubectl -- create configmap postgres-crondump --from-file=./scripts/dump.sh
                    minikube -p "${minikube_profile}" kubectl -- \
                        create secret docker-registry regcred \
                        --docker-server="${snapshot_docker_registry}" \
                        --docker-username="${snapshot_docker_username}" \
                        --docker-password="${snapshot_docker_password}" \
                        --docker-email="${snapshot_docker_email}"
                    minikube -p "${minikube_profile}" kubectl -- create -f "deployments/k8s-yaml/cronjob-snapshot.yaml"
                else
                    echo "Unknown command" >&2
                    exit 1
                fi
                ;;
            demos-scraper)
                echo "Creating demo ..."
                if [ "$arg_verbose_mode" = true ]; then echo "Creating CEX scrapers ..."; fi
                for i in "${demos_scraper_cex[@]}"
                do
                    minikube -p "${minikube_profile}" kubectl -- create -f "deployments/k8s-yaml/scraper-cex-$i.yaml"
                done
                if [ "$arg_verbose_mode" = true ]; then echo "Creating DEX scrapers ..."; fi
                for i in "${demos_scraper_dex[@]}"
                do
                    minikube -p "${minikube_profile}" kubectl -- create -f "deployments/k8s-yaml/scraper-dex-$i.yaml"
                done
                if [ "$arg_verbose_mode" = true ]; then echo "Creating foreign scrapers ..."; fi
                for i in "${demos_scraper_foreign[@]}"
                do
                    minikube -p "${minikube_profile}" kubectl -- create -f "deployments/k8s-yaml/scraper-foreign-$i.yaml"
                done
                if [ "$arg_verbose_mode" = true ]; then echo "Creating liquidity scrapers ..."; fi
                for i in "${demos_scraper_liquidity[@]}"
                do
                    minikube -p "${minikube_profile}" kubectl -- create -f "deployments/k8s-yaml/scraper-liquidity-$i.yaml"
                done
                echo "Demo created with success"
                ;;
            *)
                echo "Unknown command" >&2
                exit 1
                ;;
            esac
        else
            echo "Unknown command" >&2
            exit 1
        fi
        ;;
    remove)
        if [ "${#command[@]}" -gt 1 ]; then
            case "${command[1]}" in
            scraper-cex)
                if [ "${#command[@]}" -gt 2 ]; then
                    if [[ " ${demos_scraper_cex[*]} " == *" ${command[2]} "* ]]; then
                        minikube -p "${minikube_profile}" kubectl -- delete -f "deployments/k8s-yaml/scraper-cex-${command[2]}.yaml" || true
                    else
                        echo "Unknown demo" >&2
                        exit 1
                    fi
                else
                    minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/scraper-cex.yaml || true
                fi
                ;;
            scraper-dex)
                if [ "${#command[@]}" -gt 2 ]; then
                    if [[ " ${demos_scraper_dex[*]} " == *" ${command[2]} "* ]]; then
                        minikube -p "${minikube_profile}" kubectl -- delete -f "deployments/k8s-yaml/scraper-dex-${command[2]}.yaml" || true
                    else
                        echo "Unknown demo" >&2
                        exit 1
                    fi
                else
                    minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/scraper-dex.yaml || true
                fi
                ;;
            scraper-liquidity)
                if [ "${#command[@]}" -gt 2 ]; then
                    if [[ " ${demos_scraper_liquidity[*]} " == *" ${command[2]} "* ]]; then
                        minikube -p "${minikube_profile}" kubectl -- delete -f "deployments/k8s-yaml/scraper-liquidity-${command[2]}.yaml" || true
                    else
                        echo "Unknown demo" >&2
                        exit 1
                    fi
                else
                    minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/scraper-liquidity.yaml || true
                fi
                ;;
            scraper-foreign)
                if [ "${#command[@]}" -gt 2 ]; then
                    if [[ " ${demos_scraper_foreign[*]} " == *" ${command[2]} "* ]]; then
                        minikube -p "${minikube_profile}" kubectl -- delete -f "deployments/k8s-yaml/scraper-foreign-${command[2]}.yaml" || true
                    else
                        echo "Unknown demo" >&2
                        exit 1
                    fi
                else
                    minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/scraper-foreign.yaml || true
                fi
                ;;
            cronjob)
                if [ "${#command[@]}" -eq 2 ]; then
                    minikube -p "${minikube_profile}" kubectl -- delete -f "deployments/k8s-yaml/cronjob-snapshot.yaml" || true
                    minikube -p "${minikube_profile}" kubectl -- delete secret regcred || true
                    minikube -p "${minikube_profile}" kubectl -- delete configmap postgres-entrypoint || true
                    minikube -p "${minikube_profile}" kubectl -- delete configmap postgres-crondump || true
                else
                    echo "Unknown command" >&2
                    exit 1
                fi
                ;;
            demos-scraper)
                echo "Removing demo ..."
                if [ "$arg_verbose_mode" = true ]; then echo "Removing CEX scrapers ..."; fi
                for i in "${demos_scraper_cex[@]}"
                do
                    minikube -p "${minikube_profile}" kubectl -- delete -f "deployments/k8s-yaml/scraper-cex-$i.yaml"
                done
                if [ "$arg_verbose_mode" = true ]; then echo "Removing DEX scrapers ..."; fi
                for i in "${demos_scraper_dex[@]}"
                do
                    minikube -p "${minikube_profile}" kubectl -- delete -f "deployments/k8s-yaml/scraper-dex-$i.yaml"
                done
                if [ "$arg_verbose_mode" = true ]; then echo "Removing foreign scrapers ..."; fi
                for i in "${demos_scraper_foreign[@]}"
                do
                    minikube -p "${minikube_profile}" kubectl -- delete -f "deployments/k8s-yaml/scraper-foreign-$i.yaml"
                done
                if [ "$arg_verbose_mode" = true ]; then echo "Removing liquidity scrapers ..."; fi
                for i in "${demos_scraper_liquidity[@]}"
                do
                    minikube -p "${minikube_profile}" kubectl -- delete -f "deployments/k8s-yaml/scraper-liquidity-$i.yaml"
                done
                echo "Demo removed with success"
                ;;
            *)
                echo "Unknown command" >&2
                exit 1
                ;;
            esac
        else
            echo "Unknown command" >&2
            exit 1
        fi
        ;;
    shell)
        if [ "${#command[@]}" -eq 1 ]; then
            minikube -p "${minikube_profile}" ssh
        else
            echo "Unknown command" >&2
            exit 1
        fi
        ;;
    logs)
        if [ "${#command[@]}" -eq 1 ]; then
            minikube -p "${minikube_profile}" logs -f
        else
            echo "Unknown command" >&2
            exit 1
        fi
        ;;
    clean)
        if [ "${#command[@]}" -eq 1 ]; then
            if [[ "$minikube_driver" == "docker" ]]; then
                eval "$(minikube -p "${minikube_profile}" docker-env)"
                docker buildx prune -a -f
                eval "$(minikube -p "${minikube_profile}" docker-env --unset)"
            fi
            minikube -p "${minikube_profile}" ssh -- docker image prune -a -f
            minikube -p "${minikube_profile}" ssh -- docker volume prune -f
            minikube -p "${minikube_profile}" ssh -- docker network prune -f
            minikube -p "${minikube_profile}" ssh -- docker system prune -a -f
        else
            echo "Unknown command" >&2
            exit 1
        fi
        ;;
    info)
        if [ "${#command[@]}" -eq 1 ]; then
            minikube profile list
            _info
            # TODO: uptime show the host uptime, not the minikube: https://stackoverflow.com/a/28353785/2042014
            _minikube_uptime=$(minikube -p "${minikube_profile}" ssh -- uptime -p)
            echo "Minikube Uptime: $_minikube_uptime"
            echo; echo "Images:"
            minikube -p "${minikube_profile}" image ls --format table
            echo; echo "Status:"
            minikube -p "${minikube_profile}" status
            echo; echo "Disk available:"
            minikube -p "${minikube_profile}" ssh -- df -h
            if [[ "$minikube_driver" == "docker" ]]; then
                echo; echo "Disk available of Docker:"
                minikube -p "${minikube_profile}" ssh -- docker system df
            fi
        else
            echo "Unknown command" >&2
            exit 1
        fi
        ;;
    ping)
        if [ "${#command[@]}" -eq 1 ]; then
            echo "Ping tests:"
            minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/job-ping-redis.yaml >/dev/null 2>&1 || true
            minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/job-ping-redis.yaml >/dev/null 2>&1
            minikube -p "${minikube_profile}" kubectl -- wait --timeout=30s --for=condition=complete job/ping-redis >/dev/null 2>&1 && echo "  Redis OK" || echo "  Redis FAILED" >&2
            minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/job-ping-influx.yaml >/dev/null 2>&1 || true
            minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/job-ping-influx.yaml >/dev/null 2>&1
            minikube -p "${minikube_profile}" kubectl -- wait --timeout=30s --for=condition=complete job/ping-influx >/dev/null 2>&1 && echo "  InfluxDB OK" || echo "  InfluxDB FAILED" >&2
            minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/job-ping-postgres.yaml >/dev/null 2>&1 || true
            minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/job-ping-postgres.yaml >/dev/null 2>&1
            minikube -p "${minikube_profile}" kubectl -- wait --timeout=30s --for=condition=complete job/ping-postgres >/dev/null 2>&1 && echo "  PostgreSQL OK" || echo "  PostgreSQL FAILED" >&2
            minikube -p "${minikube_profile}" kubectl -- delete -f deployments/k8s-yaml/job-ping-kafka.yaml >/dev/null 2>&1 || true
            minikube -p "${minikube_profile}" kubectl -- create -f deployments/k8s-yaml/job-ping-kafka.yaml >/dev/null 2>&1
            minikube -p "${minikube_profile}" kubectl -- wait --timeout=30s --for=condition=complete job/ping-kafka >/dev/null 2>&1 && echo "  Kafka OK" || echo "  Kafka FAILED" >&2
        else
            echo "Unknown command" >&2
            exit 1
        fi
        ;;
    data-reset)
        # TODO: Also reset relational data and cache
        if [ "${#command[@]}" -eq 1 ]; then
            echo "Resetting timeseries data ..."
            minikube -p "${minikube_profile}" kubectl -- exec deployment/data-influx -- influx -host 'localhost' -port '8086' -database 'dia' -execute 'DROP MEASUREMENT foreignquotation'
            minikube -p "${minikube_profile}" kubectl -- exec deployment/data-influx -- influx -host 'localhost' -port '8086' -database 'dia' -execute 'DROP MEASUREMENT tradesTmp'
        else
            echo "Unknown command" >&2
            exit 1
        fi
        ;;
    data)
        # TODO: "SELECT ... | wc" is very slow, maybe investigate to use COUNT statement
        if [ "${#command[@]}" -gt 1 ]; then
            case "${command[1]}" in
            exchanges)
                num_exchanges=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM exchange")
                echo "Exchanges:"
                minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT * FROM exchange"
                echo "#Exchanges: $num_exchanges"
                ;;
            exchange)
                echo "Exchange:"
                read -r inputExchangeName
                minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -t -c "SELECT exchange_id FROM exchange WHERE name = '$inputExchangeName'" | if [[ $(wc -l) -ge 2 ]]; then echo "Found"; else echo "Not Found"; exit 1; fi
                isbridge=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT bridge FROM exchange WHERE name = '$inputExchangeName'")
                isactive=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT scraper_active FROM exchange WHERE name = '$inputExchangeName'")
                iscentralized=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -AXqt -c "SELECT centralized FROM exchange WHERE name = '$inputExchangeName'" | sed 's/\r$//')
                echo "Bridge: $isbridge"
                echo "Active: $isactive"
                echo "Centralized: $iscentralized"                
                if [[ "${iscentralized}" == "t" ]]; then
                    num_exchangepairs=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM exchangepair WHERE exchange = '$inputExchangeName'")
                    num_exchangepairassets=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM asset WHERE asset_id IN (SELECT id_quotetoken FROM exchangepair WHERE exchange = '$inputExchangeName') OR asset_id IN (SELECT id_basetoken FROM exchangepair WHERE exchange = '$inputExchangeName')")
                    echo "PostgreSQL #exchangepair: $num_exchangepairs (#asset: $num_exchangepairassets)"
                    num_exchangesymbols=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM exchangesymbol WHERE exchange = '$inputExchangeName'")
                    num_exchangesymbolsassets=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM asset WHERE asset_id IN (SELECT asset_id FROM exchangesymbol WHERE exchange = '$inputExchangeName')")
                    echo "PostgreSQL #exchangesymbol: $num_exchangesymbols (#asset: $num_exchangesymbolsassets)"
                else
                    blockchain=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT blockchain FROM exchange WHERE name = '$inputExchangeName'")
                    echo "Blockchain: $blockchain"
                    num_pools=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM pool WHERE exchange = '$inputExchangeName'")
                    last_liquidityupdate=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT time_stamp FROM poolasset WHERE pool_id IN (SELECT pool_id FROM pool WHERE exchange = '$inputExchangeName') ORDER BY time_stamp DESC LIMIT 1")
                    num_poolassets=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM poolasset WHERE pool_id IN (SELECT pool_id FROM pool WHERE exchange = '$inputExchangeName')")
                    num_poolassetsassets=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM asset WHERE asset_id IN (SELECT asset_id FROM poolasset WHERE pool_id IN (SELECT pool_id FROM pool WHERE exchange = '$inputExchangeName'))")
                    echo "PostgreSQL #pool: $num_pools (#poolasset: $num_poolassets updated at $last_liquidityupdate (#asset: $num_poolassetsassets))"
                fi
                num_trades=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-influx -- influx -host 'localhost' -port '8086' -database 'dia' -execute "SELECT * FROM tradesTmp WHERE exchange = '$inputExchangeName'" | wc -l)
                echo "InfluxDB #trades: $num_trades"
                ;;
            blockchain)
                echo "Blockchain:"
                read -r inputBlockchainName
                verificationmechanism=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT verificationmechanism FROM blockchain WHERE name = '$inputBlockchainName'")
                genesisdate=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT genesisdate FROM blockchain WHERE name = '$inputBlockchainName'")
                chainid=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT chain_id FROM blockchain WHERE name = '$inputBlockchainName'")
                echo "Consensus: $verificationmechanism"
                echo "Genesis Date: $genesisdate"
                echo "Chain ID: $chainid"
                num_blockchainassets=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM asset WHERE blockchain = '$inputBlockchainName'")
                echo "PostgreSQL #asset: $num_blockchainassets"
                ;;
            scraper-cex)
                echo wip
                ;;
            scraper-dex)
                echo wip
                ;;
            scraper-liquidity)
                echo wip
                ;;
            scraper-foreign)
                echo "Exchange:"
                read -r inputExchangeName
                num_foreignquotation=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-influx -- influx -host 'localhost' -port '8086' -database 'dia' -execute "SELECT * FROM foreignquotation WHERE source = '$inputExchangeName'" | wc -l)
                echo "InfluxDB #foreignquotation: $num_foreignquotation"
                minikube -p "${minikube_profile}" kubectl -- exec deployment/data-influx -- influx -host 'localhost' -port '8086' -database 'dia' -execute "SELECT * FROM foreignquotation WHERE source = '$inputExchangeName' ORDER BY time DESC LIMIT 10"
                ;;
            *)
                echo "Unknown command" >&2
                exit 1
                ;;
            esac
        else
            # InfluxDB
            num_trades=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-influx -- influx -host 'localhost' -port '8086' -database 'dia' -execute 'SELECT * FROM trades' | wc -l)
            echo "InfluxDB #trades: $num_trades"
            num_tradestmp=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-influx -- influx -host 'localhost' -port '8086' -database 'dia' -execute 'SELECT * FROM tradesTmp' | wc -l)
            echo "InfluxDB #tradesTmp: $num_tradestmp"
            num_filters=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-influx -- influx -host 'localhost' -port '8086' -database 'dia' -execute 'SELECT * FROM filters' | wc -l)
            echo "InfluxDB #filters: $num_filters"
            num_quotations=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-influx -- influx -host 'localhost' -port '8086' -database 'dia' -execute 'SELECT * FROM assetQuotations' | wc -l)
            echo "InfluxDB #assetQuotations: $num_quotations"
            num_foreignquotation=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-influx -- influx -host 'localhost' -port '8086' -database 'dia' -execute 'SELECT * FROM foreignquotation' | wc -l)
            echo "InfluxDB #foreignquotation: $num_foreignquotation"
            # PostgreSQL
            num_exchanges=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM exchange")
            echo "PostgreSQL #exchange: $num_exchanges"
            num_pools=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM pool")
            echo "PostgreSQL #pool: $num_pools"
            num_poolassets=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM poolasset")
            echo "PostgreSQL #poolasset: $num_poolassets"
            num_assets=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM asset")
            echo "PostgreSQL #asset: $num_assets"
            num_exchangepairs=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM exchangepair")
            echo "PostgreSQL #exchangepair: $num_exchangepairs"
            num_exchangesymbols=$(minikube -p "${minikube_profile}" kubectl -- exec deployment/data-postgres -- psql -U postgres -tA -c "SELECT COUNT(*) FROM exchangesymbol")
            echo "PostgreSQL #exchangesymbol: $num_exchangesymbols"
        fi
        ;;
    *)
        echo "Unknown command ${command[0]}" >&2
        exit 1
        ;;
    esac
}

main "$@"