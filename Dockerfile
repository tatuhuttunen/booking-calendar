FROM golang:1.9
COPY . /go/src/github.com/tatuhuttunen/booking-calendar
WORKDIR /go/src/github.com/tatuhuttunen/booking-calendar
RUN cd cmd/meetings && go build .
RUN cd cmd/users && go build .
RUN cd cmd/api && go build .
