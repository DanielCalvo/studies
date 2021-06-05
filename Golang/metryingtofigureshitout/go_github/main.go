package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"io/ioutil"
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

	io_reader, err := client.Repositories.DownloadContents(context.Background(), "DanielCalvo", "markdownscanner", "internal/config/config_test.go", nil)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(io_reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

}

func GetRawUrl(url string) string {

	//Step 1: Replace s/github.com/raw.githubusercontent.com/
	//Step 2: remove "blob/" from the url
	//Step 3: Make a HTTP request and it should work

	return ""
}
