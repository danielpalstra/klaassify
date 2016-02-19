/*
	This file contains the available CLI flags for handling ENC entries.
*/

package cli

import (
	"github.com/codegangsta/cli"
	"github.com/danielpalstra/klaassify/enc"
)

var entry enc.Entry

// Flags contains a list of available arguments to use
var Flags = []cli.Flag{
	flagHostname,
	flagEnvironment,
	flagClass,
	flagProject,
	flagType,
	flagSize,
	flagGitlabToken,
	flagRepositoryURL,
	flagRepositoryName,
	flagRepositoryNamespace,
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

var flagRole = cli.StringFlag{
	Name:        "role, r",
	Usage:       "Size of node to add to ENC entry",
	Destination: &entry.Parameters.Role,
}

var flagGitlabToken = cli.StringFlag{
	Name:   "token",
	Usage:  "Gitlab API key",
	EnvVar: "GITLAB_API_KEY",
}

var flagRepositoryURL = cli.StringFlag{
	Name:   "url",
	Usage:  "Git remote used by ENC",
	EnvVar: "GITLAB_API_URL",
}

var flagRepositoryName = cli.StringFlag{
	Name:   "repository",
	Usage:  "Git repository used by ENC",
	EnvVar: "ENC_GIT_REPOSITORY",
}

var flagRepositoryNamespace = cli.StringFlag{
	Name:   "namespace",
	Usage:  "Gitlab namespace for ENC git repository",
	EnvVar: "ENC_GIT_NAMESPACE",
}
