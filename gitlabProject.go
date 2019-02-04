package gitlabClient


// 根据token判断所属项目
func (gitlab *GitlabClient) ListOwnProject() ([]byte, error) {
	data, err := gitlab.get("/api/v3/projects/owned", nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}