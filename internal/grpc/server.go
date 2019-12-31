package grpc

import (
	"context"
	"log"
	"net"

	proto "github.com/golang/protobuf/proto"
	"github.com/linnv/logx"
	"google.golang.org/grpc"

	conf "qnmock/config"
	pb "qnmock/model/addrgcable"
)

var config *conf.Configuration
var addressGcables *conf.AddressGcables

func Init() {
	config = conf.Config()
	addressGcables = conf.GetAddressGcables()
}

type server struct {
	pb.UnimplementedGcableToubleServer
}

func (s *server) CheckToubleAreaGcable(ctx context.Context, in *pb.RequestGcableTrouble) (*pb.ResponseGcableTrouble, error) {
	logx.Debugfln("Received: %s,%v", in.String(), in.GetAddress())
	resp := &pb.ResponseGcableTrouble{}
	if addressGcables.MatchShort(in.GetAddress()) {
		resp.Status = *proto.Int32(NORMAL)
	} else if addressGcables.Match(in.GetAddress()) {
		resp.Status = *proto.Int32(NORMAL)
		resp.ShouldGiveMoreDetail = *proto.Bool(true)
	}
	return resp, nil
}

func StartEngine(exit chan struct{}) *grpc.Server {
	lis, err := net.Listen("tcp", config.GrpcServerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGcableToubleServer(s, &server{})

	go func() {
		logx.Debugf("grpc server runs at port: %+v\n", config.GrpcServerPort)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return s
}
