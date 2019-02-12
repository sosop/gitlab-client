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

// 获取所属及可访问项目
func (gitlab *GitlabClient) ListAllProjects() ([]ProjectInfo, error) {
	own, err := gitlab.ListOwnProjects()
	if err != nil {
		return nil, err
	}
	access, err := gitlab.ListProjects()
	if err != nil {
		return nil, err
	}
	all := make([]ProjectInfo, 0, 1024)
	all = append(all, own...)
	all = append(all, access...)
	return all, nil
}

func generateProject(data []byte) ([]ProjectInfo, error) {
	projs := make([]ProjectInfo, 0, 1024)
	err := json.Unmarshal(data, &projs)
	if err != nil {
		return nil, err
	}
	return projs, nil
}