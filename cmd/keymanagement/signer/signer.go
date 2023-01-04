package main

import (
	"context"
	"encoding/json"
	"net"

	kr "github.com/99designs/keyring"
	"github.com/99designs/keyring/cmd/k8sbridge"
	pb "github.com/diadata-org/diadata/pkg/dia/helpers/signer/protoc"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var logger = logrus.New()

func main() {
	kr.Debug = true
	k8bridgeurl := utils.Getenv("K8SBRIDGE_URL", "127.0.0.1:50051")
	signerport := utils.Getenv("SIGNER_PORT", ":50052")

	ring, _ := kr.Open(kr.Config{
		ServiceName:     "signer",
		Server:          k8bridgeurl,
		AllowedBackends: []kr.BackendType{kr.K8Secret},
	})

	s := grpc.NewServer()

	pb.RegisterSignerServer(s, &server{Keyring: ring})

	lis, err := net.Listen("tcp", signerport)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	logger.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}

}

type server struct {
	pb.UnimplementedSignerServer
	Keyring kr.Keyring
}

func (s *server) Sign(_ context.Context, request *pb.DataToSign) (*pb.SignedData, error) {

	var keypair *k8sbridge.KeyPair

	item, err := s.Keyring.Get(request.By)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(item.Data, &keypair)
	logger.Infoln("public key", keypair.GetPublickey())
	pk, err := crypto.ToECDSA(common.FromHex(keypair.GetPrivatekey()))
	if err != nil {
		return nil, err
	}
	sig, err := crypto.Sign(request.Data, pk)
	return &pb.SignedData{Signed: sig}, err
}
