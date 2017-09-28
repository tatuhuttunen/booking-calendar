package main

import (
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/tatuhuttunen/booking-calendar/pb/users"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	users []*users.User
}

func (server) GetUser(context.Context, *users.GetUserRequest) (*users.User, error) {
	panic("implement me")
}

func (s server) ListUsers(context.Context, *users.ListUsersRequest) (*users.ListUsersResponse, error) {
	res := new(users.ListUsersResponse)
	res.Users = append(res.Users, s.users...)
	res.NextPageToken = "users tokeni"
	return res, nil
}

func (s *server) CreateUser(ctx context.Context, in *users.CreateUserRequest) (*users.User, error) {
	s.users = append(s.users, in.User)
	return in.User, nil
}

func (server) UpdateUser(context.Context, *users.UpdateUserRequest) (*users.User, error) {
	panic("implement me")
}

func (server) DeleteUser(context.Context, *users.DeleteUserRequest) (*empty.Empty, error) {
	panic("implement me")
}

func main() {
	port := 8080
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	users.RegisterUsersServer(srv, &server{make([]*users.User, 0)})
	err = srv.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
