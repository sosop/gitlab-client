package gitlabClient

import "os/exec"

// 依赖客户端git
func (gitlab *GitLabClient) GetRepo() ([]byte, error) {
	return exec.Command("git clone -b " + gitlab.Branch + "http://" + gitlab.Username + ":" + gitlab.Password + "@" + gitlab.WebURL + " /tmp/" + gitlab.ProjectName).Output()
}
