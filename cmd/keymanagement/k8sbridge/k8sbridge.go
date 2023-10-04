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
	kb := k8util.New(clientset)

	pb.RegisterK8SHelperServer(s, &server{kb: kb})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Infoln("server listening at %v", lis.Addr())
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
