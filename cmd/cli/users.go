package main

import (
	"errors"
	"fmt"

	"github.com/tatuhuttunen/booking-calendar/pb/users"
	menu "github.com/turret-io/go-menu/menu"
	"golang.org/x/net/context"
)

func listUsers(args ...string) error {
	if !validArgs(2, args...) {
		return errors.New("invalid arguments")
	}

	pageSize, err := intArg(args[0])
	if err != nil {
		return errors.New("Invalid page size")
	}

	res, err := cls.usersClient.ListUsers(context.Background(), &users.ListUsersRequest{
		PageSize:  pageSize,
		PageToken: args[1],
	})
	fmt.Println("users")
	fmt.Println(res.Users)
	fmt.Println("token")
	fmt.Println(res.NextPageToken)
	fmt.Println(err)

	return nil
}

func createUsers(args ...string) error {
	if !validArgs(3, args...) {
		return errors.New("invalid arguments")
	}

	fmt.Println(cls.usersClient.CreateUser(context.Background(), &users.CreateUserRequest{
		User: &users.User{
			Email: args[0],
			Name:  args[1],
			Phone: args[2],
		},
	}))
	return nil
}

func usersCli(args ...string) error {
	commandOptions := []menu.CommandOption{
		menu.CommandOption{
			Command:     "list",
			Description: "pageSize pageToken",
			Function:    listUsers,
		},
		menu.CommandOption{
			Command:     "create",
			Description: "email name phone",
			Function:    createUsers,
		},
	}

	menuOptions := menu.NewMenuOptions("users cli> ", 0, "")

	menu := menu.NewMenu(commandOptions, menuOptions)
	menu.Start()
	return nil
}
