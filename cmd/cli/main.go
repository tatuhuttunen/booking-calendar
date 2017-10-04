package main

import (
	"github.com/turret-io/go-menu/menu"

	"log"

	"strconv"

	"github.com/tatuhuttunen/booking-calendar/pb/meetings"
	"github.com/tatuhuttunen/booking-calendar/pb/users"
	"google.golang.org/grpc"
)

type clientsCollection struct {
	meetingsClient meetings.MeetingsClient
	usersClient    users.UsersClient
}

var (
	cls clientsCollection
)

func main() {
	var (
		meetingsAddr = "meetings:8080"
		clientsAddr  = "users:8080"
	)
	cls = clientsCollection{
		meetingsClient: meetings.NewMeetingsClient(mustDial(meetingsAddr)),
		usersClient:    users.NewUsersClient(mustDial(clientsAddr)),
	}

	commandOptions := []menu.CommandOption{
		{
			Command:     "meetings",
			Description: "manage meetings",
			Function:    meetingsCli,
		},
		{
			Command:     "users",
			Description: "manage users",
			Function:    usersCli,
		},
	}

	menuOptions := menu.NewMenuOptions("calendar cli> ", 0)

	m := menu.NewMenu(commandOptions, menuOptions)
	m.Start()
}

// mustDial ensures a tcp connection to specified address.
func mustDial(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
		panic(err)
	}
	return conn
}

func validArgs(count int, args ...string) bool {
	return len(args) == count
}

func intArg(arg string) (int32, error) {
	v, err := strconv.ParseInt(arg, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}
