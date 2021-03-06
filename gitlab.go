package gitlabClient

import (
	"net/http"
	"io/ioutil"
	"github.com/json-iterator/go"
	"net/url"
	"strings"
)

var (
	GitInfo *gitlabInfo = nil
	GitlabClients = make(map[string]*GitlabClient, 1024)
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type gitlabInfo struct {
	client 			*http.Client
	Address 		string			`json:"address"`
}

type GitlabClient struct {
	*gitlabInfo
	PrivateToken 	string			`json:"privateToken"`
}

type ProjectInfo struct {
	Branch			string	`json:"branch"`
	GitAddr			string	`json:"ssh_url_to_repo"`
	ProjectName		string	`json:"name"`
	ProjectID 		int		`json:"id"`
	FullName		string	`json:"path_with_namespace"`
	GitlabClient
}

func NewProject(gitAddr, projectName string, projectID int) ProjectInfo {
	return ProjectInfo{GitAddr: gitAddr, ProjectName: projectName, ProjectID: projectID}
}

func NewGitlabInfo(client *http.Client, address string) *gitlabInfo {
	if client == nil {
		client = http.DefaultClient
	}
	return &gitlabInfo{client: client, Address: address}
}

func PushGitlabClient(privateToken string) {
	GitlabClients[privateToken] = NewGitlabClient(privateToken)
}


func NewGitlabClient(privateToken string) *GitlabClient {
	if GitInfo ==  nil {
		return nil
	}
	return &GitlabClient{gitlabInfo: GitInfo, PrivateToken: privateToken}
}

func (gitlab *GitlabClient) get(uri string, headers map[string]string) ([]byte, error) {
	return request(gitlab, "GET", uri, headers, nil)
}

func (gitlab *GitlabClient) post(uri string, headers map[string]string, params url.Values) ([]byte, error) {
	return request(gitlab, "POST", uri, headers, params)
}

func request(gitlab *GitlabClient,  method, uri string, headers map[string]string, params url.Values) ([]byte, error) {

	var req *http.Request
	var err error

	if strings.ToUpper(method) == "POST" {
		req, err = http.NewRequest(method, gitlab.Address + uri, strings.NewReader(params.Encode()))
	} else {
		req, err = http.NewRequest(method, gitlab.Address + uri, nil)
	}

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
