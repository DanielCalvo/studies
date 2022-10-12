package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
	"io/ioutil"
	"net/http"
	"time"
)

/*
Looks like this works:
curl \
  -H 'Accept: application/vnd.github.v3.raw' \
  -O \
  -L https://api.github.com/repos/DanielCalvo/markdownscanner/contents/main.go
But the URL above is different -- we either have to transform the URL we see in the browser to a raw url or an api one

API: https://api.github.com/repos/DanielCalvo/markdownscanner/contents/main.go
URL: https://github.com/DanielCalvo/markdownscanner/blob/master/main.go
RAW: https://raw.githubusercontent.com/DanielCalvo/markdownscanner/master/main.go
*/

type GithubProjectApiResponse []struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"owner"`
	HTMLURL          string    `json:"html_url"`
	Description      string    `json:"description"`
	Fork             bool      `json:"fork"`
	URL              string    `json:"url"`
	ForksURL         string    `json:"forks_url"`
	KeysURL          string    `json:"keys_url"`
	CollaboratorsURL string    `json:"collaborators_url"`
	TeamsURL         string    `json:"teams_url"`
	HooksURL         string    `json:"hooks_url"`
	IssueEventsURL   string    `json:"issue_events_url"`
	EventsURL        string    `json:"events_url"`
	AssigneesURL     string    `json:"assignees_url"`
	BranchesURL      string    `json:"branches_url"`
	TagsURL          string    `json:"tags_url"`
	BlobsURL         string    `json:"blobs_url"`
	GitTagsURL       string    `json:"git_tags_url"`
	GitRefsURL       string    `json:"git_refs_url"`
	TreesURL         string    `json:"trees_url"`
	StatusesURL      string    `json:"statuses_url"`
	LanguagesURL     string    `json:"languages_url"`
	StargazersURL    string    `json:"stargazers_url"`
	ContributorsURL  string    `json:"contributors_url"`
	SubscribersURL   string    `json:"subscribers_url"`
	SubscriptionURL  string    `json:"subscription_url"`
	CommitsURL       string    `json:"commits_url"`
	GitCommitsURL    string    `json:"git_commits_url"`
	CommentsURL      string    `json:"comments_url"`
	IssueCommentURL  string    `json:"issue_comment_url"`
	ContentsURL      string    `json:"contents_url"`
	CompareURL       string    `json:"compare_url"`
	MergesURL        string    `json:"merges_url"`
	ArchiveURL       string    `json:"archive_url"`
	DownloadsURL     string    `json:"downloads_url"`
	IssuesURL        string    `json:"issues_url"`
	PullsURL         string    `json:"pulls_url"`
	MilestonesURL    string    `json:"milestones_url"`
	NotificationsURL string    `json:"notifications_url"`
	LabelsURL        string    `json:"labels_url"`
	ReleasesURL      string    `json:"releases_url"`
	DeploymentsURL   string    `json:"deployments_url"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	PushedAt         time.Time `json:"pushed_at"`
	GitURL           string    `json:"git_url"`
	SSHURL           string    `json:"ssh_url"`
	CloneURL         string    `json:"clone_url"`
	SvnURL           string    `json:"svn_url"`
	Homepage         string    `json:"homepage"`
	Size             int       `json:"size"`
	StargazersCount  int       `json:"stargazers_count"`
	WatchersCount    int       `json:"watchers_count"`
	Language         string    `json:"language"`
	HasIssues        bool      `json:"has_issues"`
	HasProjects      bool      `json:"has_projects"`
	HasDownloads     bool      `json:"has_downloads"`
	HasWiki          bool      `json:"has_wiki"`
	HasPages         bool      `json:"has_pages"`
	ForksCount       int       `json:"forks_count"`
	MirrorURL        string    `json:"mirror_url"`
	Archived         bool      `json:"archived"`
	Disabled         bool      `json:"disabled"`
	OpenIssuesCount  int       `json:"open_issues_count"`
	License          struct {
		Key    string `json:"key"`
		Name   string `json:"name"`
		SpdxID string `json:"spdx_id"`
		URL    string `json:"url"`
		NodeID string `json:"node_id"`
	} `json:"license"`
	Forks         int    `json:"forks"`
	OpenIssues    int    `json:"open_issues"`
	Watchers      int    `json:"watchers"`
	DefaultBranch string `json:"default_branch"`
	Permissions   struct {
		Admin bool `json:"admin"`
		Push  bool `json:"push"`
		Pull  bool `json:"pull"`
	} `json:"permissions"`
}

func main() {
	fmt.Println("sup")

	GetRawUrl("https://github.com/DanielCalvo/markdownscanner/blob/master/cmd/root.go")
	//https://raw.githubusercontent.com/DanielCalvo/markdownscanner/master/cmd/root.go

	client := github.NewClient(nil)
	//orgs, _, _ := client.Organizations.List(context.Background(), "torvalds", nil)
	//fmt.Println(orgs)

	//fmt.Println(client.Repositories.List(context.Background(), "DanielCalvo", nil))

	//func (s *RepositoriesService) GetReadme(ctx context.Context, owner, repo string, opt *RepositoryContentGetOptions) (*RepositoryContent, *Response, error) {
	repocontent, _, _ := client.Repositories.GetReadme(context.Background(), "DanielCalvo", "markdownscanner", nil)
	fmt.Println(repocontent.GetContent())

	//https://github.com/DanielCalvo/markdownscanner/blob/master/internal/config/config_test.go
	io_reader, err := client.Repositories.DownloadContents(context.Background(), "DanielCalvo", "markdownscanner", "internal/config/config_test.go", nil)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(io_reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	//Let's get a list of repos form a github project. Let's say, lets get all the repos for the kubernetes project

	req, err := http.NewRequest("GET", "https://api.github.com/orgs/kubernetes/repos?per_page=2000", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/vnd.github.inertia-preview+json")

	httpClient := &http.Client{Timeout: time.Second * 10}

	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	ghResponse := GithubProjectApiResponse{}

	err = json.Unmarshal(body, &ghResponse)
	if err != nil {
		panic(err)
	}

	for _, repo := range ghResponse {
		fmt.Println(repo.HTMLURL)
	}

}

func GetRawUrl(url string) string {

	//Step 1: Replace s/github.com/raw.githubusercontent.com/
	//Step 2: remove "blob/" from the url
	//Step 3: Make a HTTP request and it should work

	return ""
}
