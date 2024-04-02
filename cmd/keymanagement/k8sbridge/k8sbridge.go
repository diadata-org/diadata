package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"path/filepath"

	"k8sbridge/k8util"

	pb "github.com/diadata-org/diadata/pkg/dia/helpers/k8sbridge/protoc"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"

	"google.golang.org/grpc"

	"github.com/diadata-org/diadata/pkg/utils"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var log = logrus.New()

/*

1) grpc to get keys using public key
2) grpc to generate key and return public key

*/

type server struct {
	pb.UnimplementedK8SHelperServer
	kb *k8util.K8Bridge
}

func main() {

	namespace := os.Getenv("NAMESPACE")
	if namespace == "" {
		namespace = "dia-oracle-feeder"
	}

	signerurl := utils.Getenv("SIGNER_URL", "signer.dia-oracle-feeder:50052")
	diaRestURL := utils.Getenv("DIA_REST_URL", "https://api.diadata.org")
	diaGraphqlURL := utils.Getenv("DIA_GRAPHQL_URL", "https://api.diadata.org/graphql/query")
	postgresqlHost := utils.Getenv("POSTGRES_HOST_POD", "dia-postgresql.dia-db")

	oraclebaseimage := utils.Getenv("ORACLE_BASE_IMAGE", "us.icr.io/dia-registry/oracles/oracle-baseimage:latest")
	affinity := utils.Getenv("ORACLE_FEEDER_AFFINITY", "default")

	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		// try using kube config
		var kubeconfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		} else {
			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
		flag.Parse()

		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err.Error())
		}

	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	s := grpc.NewServer()
	kb := k8util.New(clientset, namespace, signerurl, diaRestURL, diaGraphqlURL, postgresqlHost, affinity, oraclebaseimage)

	pb.RegisterK8SHelperServer(s, &server{kb: kb})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

func (s *server) CreateKeypair(_ context.Context, request *pb.K8SHelperRequest) (*pb.KeyPair, error) {
	publickey, err := s.kb.GenerateKey(request.Keyname)
	if err != nil {
		return &pb.KeyPair{}, err

	}

	return &pb.KeyPair{Publickey: publickey}, nil
}

func (s *server) GetKey(ctx context.Context, request *pb.K8SHelperRequest) (*pb.KeyPair, error) {
	publickey, err := s.kb.GetKeys(request.Keyname)
	if err != nil {
		return &pb.KeyPair{}, err

	}

	return &pb.KeyPair{Publickey: fmt.Sprintf("%s", publickey.Data[".public"]), Privatekey: fmt.Sprintf("%s", publickey.Data[".private"])}, nil
}

func (s *server) CreatePod(ctx context.Context, request *pb.FeederConfig) (*pb.CreatePodResult, error) {
	err := s.kb.CreateOracleFeeder(ctx, request.FeederID, request.Creator, request.FeederAddress, request.Oracle, request.ChainID, request.Symbols, request.FeedSelection, request.Blockchainnode, request.Frequency, request.SleepSeconds, request.DeviationPermille, request.MandatoryFrequency)
	if err != nil {
		return &pb.CreatePodResult{}, err

	}

	return &pb.CreatePodResult{}, err
}

func (s *server) RestartPod(ctx context.Context, request *pb.FeederConfig) (*pb.RestartPodResult, error) {
	err := s.kb.CreateOracleFeeder(ctx, request.FeederID, request.Creator, request.FeederAddress, request.Oracle, request.ChainID, request.Symbols, request.FeedSelection, request.Blockchainnode, request.Frequency, request.SleepSeconds, request.DeviationPermille, request.MandatoryFrequency)
	if err != nil {
		return &pb.RestartPodResult{}, err

	}

	return &pb.RestartPodResult{}, err
}

func (s *server) DeletePod(ctx context.Context, request *pb.FeederConfig) (*pb.DeletePodResult, error) {
	err := s.kb.DeleteOracleFeeder(ctx, request.FeederID)
	if err != nil {
		return &pb.DeletePodResult{}, err

	}

	return &pb.DeletePodResult{}, err
}
