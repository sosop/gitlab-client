package gitlabClient

// 获取所属项目
func (gitlab *GitlabClient) ListOwnProjects() ([]ProjectInfo, error) {
	data, err := gitlab.get("/api/v3/projects/owned", nil)
	if err != nil {
		return nil, err
	}
	return generateProject(data)
}

// 获取可访问项目
func (gitlab *GitlabClient) ListProjects() ([]ProjectInfo, error) {
	data, err := gitlab.get("/api/v3/projects", nil)
	if err != nil {
		return nil, err
	}
	return generateProject(data)
}

func generateProject(data []byte) ([]ProjectInfo, error) {
	projs := make([]ProjectInfo, 0, 1024)
	err := json.Unmarshal(data, &projs)
	if err != nil {
		return nil, err
	}
	return projs, nil
}