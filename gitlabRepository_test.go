package gitlabClient

import "testing"

func TestGitLabClient_GetRepo(t *testing.T) {
	data, err := gitlab.GetRepo(NewProject("", "", 108))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}