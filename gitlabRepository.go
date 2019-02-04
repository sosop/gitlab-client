package gitlabClient

import (
	"os/exec"
)

// 依赖客户端git
func (gitlab *GitlabClient) GetRepo(project ProjectInfo) ([]byte, error) {
	return exec.Command("git clone -b " + project.Branch + " " + project.GitAddr + " /tmp/" + project.ProjectName).Output()
}
