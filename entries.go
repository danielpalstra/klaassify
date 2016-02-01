package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

// Entry contains the structure to add to the ENC
type Entry struct {
	Environment string     `yaml:"environment,omitempty"`
	Classes     []string   `yaml:"classes,omitempty"`
	Parameters  Parameters `yaml:"parameters,omitempty"`
}

// Parameters contain the parameters a node is classified with
type Parameters struct {
	Project string `yaml:"project,omitempty"`
	Type    string `yaml:"type,omitempty"`
	Size    string `yaml:"size,omitempty"`
}

// AddENCEntry adds a new enc entry file to the repository
func AddENCEntry(hostName string, class string, entry Entry) {
	// TODO class should be directly injected in the entry array
	b := []string{class}
	entry.Classes = b

	// Marshal to yaml
	enc, err := yaml.Marshal(entry)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	writeToFile(enc, hostName)
}

func writeToFile(source []byte, fileName string) {
	// Build yaml header and footer
	header := []byte("---\n")
	footer := []byte("\n\n")
	f := append(append(header, source...), footer...)

	fileName += ".yaml"

	err := ioutil.WriteFile(fileName, f, 0644)
	if err != nil {
		log.Fatalf("Error wring %s to file: %s\n", source, err)
	}
	updateGitRepository(fileName)
}

func updateGitRepository(file string) {

	// Clone the git repo

	// Check if the working directory is a git repository
	err := exeCmd("git rev-parse --is-inside-work-tree")
	if err != nil {
		log.Fatalf("Working directory is not a git repository: %s \n", err)
	}

	// Commit the changes
	err = exeCmd("git add .")
	if err != nil {
		log.Fatalf("Error when adding files to git")
	}

	err = exeCmd("git commit -m 'Added new enc entry for " + file + "'")
	if err != nil {
		log.Fatalf("Error when committing new ENC entry: %s", err)
	}

	// Push it

}

func exeCmd(i string) error {

	// splitting head => g++ parts => rest of the command
	fmt.Println(i)
	parts := strings.Fields(i)
	head := parts[0]
	parts = parts[1:len(parts)]

	// result, err := exec.Command(head, parts...).Output()
	cmd := exec.Command(head, parts...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())

	}
	return err
}
