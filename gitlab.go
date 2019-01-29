package gitlabClient

import (
	"net/http"
	"io/ioutil"
)

var GitInfo *gitlabInfo = nil

type gitlabInfo struct {
	client 			*http.Client
	Address 		string
}

type GitLabClient struct {
	*gitlabInfo
	PrivateToken 	string
	Username		string
	Password		string
	Branch			string
	WebURL			string
	ProjectName		string
	ProjectID 		int
}


func NewGitlabInfo(client *http.Client, address string) *gitlabInfo {
	if client == nil {
		client = http.DefaultClient
	}
	return &gitlabInfo{client: client, Address: address}
}


func NewGitLabClient(privateToken, username, password string) *GitLabClient {
	if GitInfo ==  nil {
		return nil
	}
	return &GitLabClient{gitlabInfo: GitInfo, PrivateToken: privateToken, Username: username, Password: password}
}

func (gitlab *GitLabClient) get(uri string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", gitlab.Address + uri, nil)
	if err != nil {
		return nil, err
	}
	// 添加请求头部信息
	if headers != nil && len(headers) > 0 {
		for k, v := range headers {
			req.Header.Add(k, v)
		}
	}
	// 增加校验私有token
	req.Header.Add("PRIVATE-TOKEN", gitlab.PrivateToken)

	resp, err := gitlab.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}