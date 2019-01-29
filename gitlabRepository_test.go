package gitlabClient

import "testing"

func TestGitLabClient_GetRepo(t *testing.T) {
	data, err := gitlab.GetRepo()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}