package api

import (
	"fmt"
	"net"
	"strconv"

	grpcmw "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/userq11/grpc-test/api/auth"
	interceptors "github.com/userq11/grpc-test/api/interceptors"
	"github.com/userq11/grpc-test/api/users"
	innergrpc "github.com/userq11/grpc-test/protobufs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"xorm.io/xorm"
)

func Run(port int, db *xorm.Engine) {
	list, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcmw.ChainUnaryServer(interceptors.GlobalRepoInjector(db))),
	)

	initAllRoutess(s)

	reflection.Register(s)

	fmt.Printf("Server is up on port %d\n", port)
	if err = s.Serve(list); err != nil {
		panic(err)
	}
}

func initAllRoutess(s *grpc.Server) {
	innergrpc.RegisterUsersServer(s, users.GetRoutes())
	innergrpc.RegisterAuthServer(s, auth.GetRoutes())
}
