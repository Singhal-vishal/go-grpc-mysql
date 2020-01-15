package main
import(
	"context"
	"../proto"
	"google.golang.org/grpc/reflection"
	"net"
	"google.golang.org/grpc"
)
//this server will implement server interface generate by service.proto file
type server struct {

}

func main(){
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e!= nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error){
	a, b := request.GetA(), request.GetB()
	result := a+b
	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error){
	a, b := request.GetA(), request.GetB()
	result := a*b
	return &proto.Response{Result: result}, nil
}
