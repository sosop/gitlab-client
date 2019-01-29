package gitlabClient

import "testing"


func TestGitLabClient_ListBranch(t *testing.T) {
	data, err := gitlab.ListBranch("54")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func TestGitLabClient_GetBranch(t *testing.T) {
	data, err := gitlab.GetBranch("392", "master")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}