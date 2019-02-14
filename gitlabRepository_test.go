package gitlabClient

import "testing"

func TestGitLabClient_GetRepo(t *testing.T) {
	project := ProjectInfo{Branch: "master", GitAddr: "git@gitlab.lixin360.com:drone/myproject.git", ProjectName: "myproject"}
	data, err := gitlab.GetRepo(project)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}