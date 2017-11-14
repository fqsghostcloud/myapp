package server

import (
	"flag"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/golang/glog"

	pb "myapp/models/grpc/service/echo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//set port
const (
	port = ":50051"
)

type echoer struct{}

//echo message which from client input
func (s *echoer) EchoHello(stream pb.Echo_EchoHelloServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if message := in.Message; len(message) > 0 {
			glog.Infoln(message)
			if err := stream.Send(&pb.Reply{Message: message}); err != nil {
				return err
			}
		}

	}

}

//echo current time
func (s *echoer) EchoTime(out *pb.Request, stream pb.Echo_EchoTimeServer) error {
	for {
		timer := time.NewTimer(5 * time.Second)
		<-timer.C
		currentTime := time.Now()
		glog.Infoln(currentTime)
		if err := stream.Send(&pb.Reply{Message: fmt.Sprint(currentTime)}); err != nil {
			return err
		}
	}
}

func Run() {
	flag.Parse()
	glog.Info("\n------------begin to run echo service!---------------\n")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		glog.Fatalf("faild to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &echoer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		glog.Fatalf("faild to server: %v", err)
	}

}
