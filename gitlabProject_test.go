package gitlabClient

import "testing"

func TestGitLabClient_ListOwnProjects(t *testing.T) {
	data, err := gitlab.ListOwnProjects()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func TestGitLabClient_ListProjects(t *testing.T) {
	data, err := gitlab.ListProjects()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}