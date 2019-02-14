package gitlabClient

import (
	"net/url"
)

// 获取所属项目
func (gitlab *GitlabClient) ListOwnProjects() ([]ProjectInfo, error) {
	data, err := gitlab.get("/api/v3/projects/owned", nil)
	if err != nil {
		return nil, err
	}
	return generateProject(data, gitlab.PrivateToken)
}

// 获取可访问项目
func (gitlab *GitlabClient) ListProjects() ([]ProjectInfo, error) {
	data, err := gitlab.get("/api/v3/projects", nil)
	if err != nil {
		return nil, err
	}
	return generateProject(data, gitlab.PrivateToken)
}

func generateProject(data []byte, privateToken string) ([]ProjectInfo, error) {
	projs := make([]ProjectInfo, 0, 1024)
	err := json.Unmarshal(data, &projs)
	if err != nil {
		return nil, err
	}
	realProjes := make([]ProjectInfo, 0, len(projs))
	for _, p := range projs {
		p.PrivateToken = privateToken
		realProjes = append(realProjes, p)
	}
	return realProjes, nil
}

func (gitlab *GitlabClient) CreateHook(projectId, callback string) ([]byte, error) {
	params := url.Values{}
	params.Set("push_events", "true")
	params.Set("url", callback)
	params.Set("id", projectId)
	return gitlab.post("/api/v3/projects/"+ projectId + "/hooks", nil, params)
}