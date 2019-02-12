package gitlabClient

import "testing"

func TestGitLabClient_ListOwnProjects(t *testing.T) {
	projs, err := gitlab.ListOwnProjects()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(projs)
}

func TestGitLabClient_ListProjects(t *testing.T) {
	projs, err := gitlab.ListProjects()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(projs)
}

func TestGitLabClient_ListAllProjects(t *testing.T) {
	projs, err := gitlab.ListAllProjects()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(projs)
}