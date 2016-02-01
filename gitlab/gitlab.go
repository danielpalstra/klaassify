package gitlab

import (
	"fmt"
	"log"

	glc "github.com/xanzy/go-gitlab"
)

// ENCBackend describes the connection properties used for the backend containing the ENC
type ENCBackend struct {
	Token      string
	URL        string
	Namespace  string
	Repository string
}

// AddToGitlab adds a new entry to Gitlab
func AddToGitlab(fileName string, content []byte, backend ENCBackend) {

	git := glc.NewClient(nil, backend.Token)
	git.SetBaseURL(backend.URL)

	// Create a new repository file
	cf := &glc.CreateFileOptions{
		FilePath:      fileName,
		BranchName:    "master",
		Encoding:      "text",
		Content:       string(content),
		CommitMessage: "Adding ENC entry " + fileName,
	}
	file, _, err := git.RepositoryFiles.CreateFile(backend.Namespace+"/"+backend.Repository, cf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
}

// RemoveFromGitlab removes a new entry to Gitlab
func RemoveFromGitlab(fileName string, backend ENCBackend) {

	git := glc.NewClient(nil, backend.Token)
	git.SetBaseURL(backend.URL)

	fo := &glc.DeleteFileOptions{
		FilePath:      fileName,
		BranchName:    "master",
		CommitMessage: "Removing ENC entry " + fileName,
	}

	_, _, err := git.RepositoryFiles.DeleteFile(backend.Namespace+"/"+backend.Repository, fo)
	if err != nil {
		log.Fatal(err)
	}
}
