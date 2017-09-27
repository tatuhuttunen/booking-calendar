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
		return errors.New("Invalid page size")
	}

	res, err := cls.meetingsClient.ListMeetings(context.Background(), &meetings.ListMeetingsRequest{
		args[0],
		pageSize,
		args[2],
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
		"",
		&meetings.Meeting{
			"parent",
			"id",
			&meetings.Time{
				"datetime",
				"tz",
			},
			&meetings.Time{
				"datetime",
				"tz",
			},
			"title",
			"desc",
			"location",
		},
	}))
	return nil
}

func meetingsCli(args ...string) error {
	commandOptions := []menu.CommandOption{
		menu.CommandOption{"list", "parent pageSize pageToken", listMeetings},
		menu.CommandOption{"create", "", createMeeting},
	}

	menuOptions := menu.NewMenuOptions("meetings cli> ", 0, "")

	menu := menu.NewMenu(commandOptions, menuOptions)
	menu.Start()
	return nil
}
