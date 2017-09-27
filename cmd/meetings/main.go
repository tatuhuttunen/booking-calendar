package main

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/tatuhuttunen/booking-calendar/pb/meetings"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	meetings []*meetings.Meeting
}

func (s server) GetMeeting(context.Context, *meetings.GetMeetingRequest) (*meetings.Meeting, error) {
	panic("implement me")
}

func (s server) ListMeetings(context.Context, *meetings.ListMeetingsRequest) (*meetings.ListMeetingsResponse, error) {
	res := new(meetings.ListMeetingsResponse)

	for _, meeting := range s.meetings {
		res.Meetings = append(res.Meetings, meeting)
	}
	res.NextPageToken = "meeting tokeni"
	return res, nil
}

func (s server) CreateMeeting(ctx context.Context, in *meetings.CreateMeetingRequest) (*meetings.Meeting, error) {
	s.meetings = append(s.meetings, in.Meeting)
	return in.Meeting, nil
}

func (s server) UpdateMeeting(context.Context, *meetings.UpdateMeetingRequest) (*meetings.Meeting, error) {
	panic("implement me")
}

func (s server) DeleteMeeting(context.Context, *meetings.DeleteMeetingRequest) (*empty.Empty, error) {
	panic("implement me")
}

func main() {
	var port int = 8080
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	meetings.RegisterMeetingsServer(srv, &server{make([]*meetings.Meeting, 0)})
	srv.Serve(lis)
}
