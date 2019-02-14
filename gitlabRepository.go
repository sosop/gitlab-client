package gitlabClient

import (
	"os/exec"
	"bytes"
)

// 依赖客户端git
func (gitlab *GitlabClient) GetRepo(project ProjectInfo) (data []byte, err error) {
	var buf bytes.Buffer
	buf.WriteString("delete dir /tmp/" + project.ProjectName + "\r\n")
	exec.Command("rm",  "-rf", "/tmp/" + project.ProjectName).Output()

	buf.WriteString("starting git clone -b " + project.Branch + " " + project.GitAddr + "\r\n")
	data, err = exec.Command("git",  "clone", "-b", project.Branch, project.GitAddr, "/tmp/" + project.ProjectName).Output()
	if err != nil {
		return
	}
	buf.Write(data)
	buf.WriteString("success to clone project")
	data = buf.Bytes()
	return
}
