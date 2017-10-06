package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/tatuhuttunen/booking-calendar/pb/meetings"
	"github.com/tatuhuttunen/booking-calendar/pb/users"
	"path"
)

var (
	usersEndpoint    = flag.String("users_endpoint", "users:8080", "users endpoint")
	meetingsEndpoint = flag.String("meetings_endpoint", "meetings:8080", "meetings endpoint")
)

func newGateway(ctx context.Context, opts ...runtime.ServeMuxOption) (http.Handler, error) {
	mux := runtime.NewServeMux(opts...)
	dialOpts := []grpc.DialOption{grpc.WithInsecure()}

	err := users.RegisterUsersHandlerFromEndpoint(ctx, mux, *usersEndpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	err = meetings.RegisterMeetingsHandlerFromEndpoint(ctx, mux, *meetingsEndpoint, dialOpts)
	if err != nil {
		return nil, err
	}
	return mux, nil
}

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	p := path.Join("cmd", "api", "swagger.json")
	http.ServeFile(w, r, p)
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()
	mux.HandleFunc("/doc/", serveSwagger)

	gw, err := newGateway(
		ctx,
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{OrigName: true, EmitDefaults: true}))

	if err != nil {
		return err
	}
	mux.Handle("/", gw)

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
