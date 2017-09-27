FROM golang:1.9
COPY . /go/src/github.com/tatuhuttunen/booking-calendar
WORKDIR /go/src/github.com/tatuhuttunen/booking-calendar
RUN go get google.golang.org/grpc \
    && go get github.com/golang/protobuf/proto \
    && go get github.com/turret-io/go-menu/menu
RUN cd cmd/meetings && go build .
RUN cd cmd/users && go build .
RUN cd cmd/cli && go build .