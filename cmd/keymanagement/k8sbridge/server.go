package main

import (
	"context"
	"fmt"

	"github.com/diadata-org/diadata/cmd/keymanagement/k8sbridge/k8util"
	pb "github.com/diadata-org/diadata/pkg/dia/helpers/k8sbridge/protoc"
)

type server struct {
	pb.UnimplementedK8SHelperServer
	kb *k8util.K8Bridge
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
