/*
	This file contains the available CLI commands that are supported
*/

package main

import (
	"errors"

	"github.com/codegangsta/cli"
)

// Commands -  List of commands that are available
var Commands = []cli.Command{
	CommandAddEntry,
}

// CommandAddEntry adds a new entry to the ENC.
var CommandAddEntry = cli.Command{
	Name:        "add",
	Usage:       "enc add [options]",
	Description: `Command to add new entry to ENC`,
	Action:      addEntry,
	Flags:       Flags,
}

// addEntry adds the new entry to the ENC
func addEntry(c *cli.Context) {

	err := checkArguments(c)
	if err != nil {
		panic(err)
	}
	AddENCEntry(c.String("hostname"), c.String("class"), entry)

}

func checkArguments(c *cli.Context) (err error) {

	if c.String("environment") == "" {
		err = errors.New("environment flag cannot be empty")
	}

	return err
}
