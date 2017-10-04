package main

import (
	"errors"
	"fmt"

	"github.com/tatuhuttunen/booking-calendar/pb/meetings"
	menu "github.com/turret-io/go-menu/menu"
	"golang.org/x/net/context"
)

func listMeetings(args ...string) error {
	if !validArgs(3, args...) {
		return errors.New("invalid arguments")
	}

	pageSize, err := intArg(args[1])
	if err != nil {
		return errors.New("invalid page size")
	}

	res, err := cls.meetingsClient.ListMeetings(context.Background(), &meetings.ListMeetingsRequest{
		Parent:    args[0],
		PageSize:  pageSize,
		PageToken: args[2],
	})

	fmt.Println("meetings:")
	fmt.Println(res.Meetings)
	fmt.Println("token")
	fmt.Println(res.NextPageToken)
	fmt.Println(err)
	return nil
}

func createMeeting(args ...string) error {
	fmt.Println(cls.meetingsClient.CreateMeeting(context.Background(), &meetings.CreateMeetingRequest{
		Parent: "",
		Meeting: &meetings.Meeting{
			Parent: "parent",
			Id:     "id",
			Start: &meetings.Time{
				DateTime: "datetime",
				TimeZone: "tz",
			},
			End: &meetings.Time{
				DateTime: "datetime",
				TimeZone: "tz",
			},
			Title:       "title",
			Description: "desc",
			Location:    "location",
		},
	}))
	return nil
}

func meetingsCli(args ...string) error {
	commandOptions := []menu.CommandOption{
		{
			Command:     "list",
			Description: "parent pageSize pageToken",
			Function:    listMeetings,
		},
		{
			Command:     "create",
			Description: "",
			Function:    createMeeting,
		},
	}

	menuOptions := menu.NewMenuOptions("meetings cli> ", 0)

	m := menu.NewMenu(commandOptions, menuOptions)
	m.Start()
	return nil
}
