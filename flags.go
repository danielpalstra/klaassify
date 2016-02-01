/*
	This file contains the available CLI flags for handling ENC entries.
*/

package main

import "github.com/codegangsta/cli"

var entry Entry

// Flags contains a list of available arguments to use
var Flags = []cli.Flag{
	flagHostname,
	flagEnvironment,
	flagClass,
	flagProject,
	flagType,
	flagSize,
}

var flagHostname = cli.StringFlag{
	Name:  "hostname, host, f",
	Usage: "Node classifier of node to add to ENC entry",
	// Destination: &entry.Classes,
}

var flagEnvironment = cli.StringFlag{
	Name:        "environment, e",
	Usage:       "Environment for ENC entry",
	Destination: &entry.Environment,
}

var flagClass = cli.StringFlag{
	Name:  "class, c",
	Usage: "Node classifier of node to add to ENC entry",
	// Destination: &entry.Classes,
}

var flagProject = cli.StringFlag{
	Name:        "project, p",
	Usage:       "Classes to add to ENC entry",
	Destination: &entry.Parameters.Project,
}

var flagType = cli.StringFlag{
	Name:        "type, t",
	Usage:       "Type of node to add to ENC entry",
	Destination: &entry.Parameters.Type,
}

var flagSize = cli.StringFlag{
	Name:        "size, s",
	Usage:       "Size of node to add to ENC entry",
	Destination: &entry.Parameters.Size,
}
