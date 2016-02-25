/*
	This file contains the available CLI commands that are supported
*/

package cli

import (
	"errors"
	"log"

	"github.com/codegangsta/cli"
	"github.com/danielpalstra/klaassify/enc"
	"github.com/danielpalstra/klaassify/gitlab"
)

// Commands -  List of commands that are available
var Commands = []cli.Command{
	CommandAddEntry,
	CommandRemoveEntry,
}

// CommandAddEntry adds a new host to the ENC.
var CommandAddEntry = cli.Command{
	Name:        "add",
	Usage:       "klaassify add [options]",
	Description: `Command to add new host to the ENC`,
	Action:      addEntry,
	Flags:       Flags,
}

// CommandRemoveEntry removes a host from the ENC
var CommandRemoveEntry = cli.Command{
	Name:        "remove",
	Usage:       "klaassify remove [options]",
	Description: `Command to remove a host from the ENC`,
	Action:      removeEntry,
	Flags:       Flags,
}

// addEntry adds the new entry to the ENC
func addEntry(c *cli.Context) {
	err := checkArguments(c)
	if err != nil {
		log.Fatal(err)
	}

	if c.String("environment") == "" {
		err = errors.New("environment flag cannot be empty")
		log.Fatal(err)
	}

	if c.String("class") == "" {
		err = errors.New("class flag cannot be empty")
		log.Fatal(err)
	}

	if c.String("size") == "" {
		err = errors.New("size flag cannot be empty")
		log.Fatal(err)
	}

	if c.String("project") == "" {
		err = errors.New("project flag cannot be empty")
		log.Fatal(err)
	}

	b := gitlab.ENCBackend{c.String("token"), c.String("url"), c.String("namespace"), c.String("repository")}

	enc.AddENCEntry(c.String("hostname"), c.String("class"), entry, b, c.Bool("force"))
}

func removeEntry(c *cli.Context) {
	err := checkArguments(c)
	if err != nil {
		log.Fatal(err)
	}

	b := gitlab.ENCBackend{c.String("token"), c.String("url"), c.String("namespace"), c.String("repository")}

	enc.RemoveENCEntry(c.String("hostname"), b, c.Bool("force"))
}

func checkArguments(c *cli.Context) (err error) {

	if c.String("token") == "" {
		err = errors.New("token flag or environment GITLAB_API_KEY variable cannot be empty")
	}

	if c.String("url") == "" {
		err = errors.New("url flag or environment variable GITLAB_API_URL cannot be empty")
	}

	if c.String("repository") == "" {
		err = errors.New("repository flag or environment variable ENC_GIT_REPOSITORY cannot be empty")
	}

	if c.String("namespace") == "" {
		err = errors.New("namespace flag or environment ENC_GIT_NAMESPACE variable cannot be empty")
	}

	if c.String("hostname") == "" {
		err = errors.New("hostname flag cannot be empty")
	}
	return err
}
