package gitlabClient

import "testing"

func TestGitLabClient_ListMyProject(t *testing.T) {
	data, err := gitlab.ListOwnProject()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
