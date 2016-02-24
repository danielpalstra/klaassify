package enc

import (
	"io/ioutil"
	"log"

	"github.com/danielpalstra/klaassify/gitlab"
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
	Role    string `yaml:"role,omitempty"`
}

// AddENCEntry adds a new ENC host entry to the repository
func AddENCEntry(hostName string, class string, entry Entry, backend gitlab.ENCBackend, force bool) {
	// TODO class should be directly injected in the entry array
	b := []string{class}
	entry.Classes = b

	// Marshal to yaml
	enc, err := yaml.Marshal(entry)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fileName, err := writeToFile(enc, hostName)
	if err != nil {
		log.Fatal(err)
	}

	// TODO implement error handling for
	gitlab.AddToGitlab(fileName, enc, backend, force)
}

// RemoveENCEntry removes an existing host from the ENC.
func RemoveENCEntry(hostName string, backend gitlab.ENCBackend, force bool) {
	gitlab.RemoveFromGitlab(hostName+".yaml", backend, force)
}

func writeToFile(source []byte, fileName string) (fn string, err error) {
	// Build yaml header and footer
	header := []byte("---\n")
	footer := []byte("\n\n")
	f := append(append(header, source...), footer...)

	fileName += ".yaml"

	err = ioutil.WriteFile(fileName, f, 0644)
	if err != nil {
		log.Fatalf("Error wring %s to file: %s\n", source, err)
	}

	return fileName, err
}
