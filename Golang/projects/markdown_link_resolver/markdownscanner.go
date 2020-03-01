package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

//This is unfinished^W bad and I feel bad

//Dani, December-2019

//The overall structure of this program follows:
//https://blog.golang.org/pipelines
//https://golang.org/doc/effective_go.html#concurrency

//From idiomatic go: functions should never panic

//COMMENTS EVERYWHERE
//TESTS!
//take repo as parameter
//take branch as parameter
//handle invalid branch and repos appropriately

//define basedir as where to clone! (anything other than /tmp for instance)
//take repositoryUrl as an arguement?
//if url doesn't begin by http, throw error!
//you need to add a way to ignore certain MarkdownLinks (minutes, slack MarkdownLinks, mailto, etc)

//to allow only certain types to be assigned, maybe make something of a const type somehow? https://blog.golang.org/constants
//statuses: OK, BROKEN, NOTAPPLICABLE, UNKNOWN
//type: FILE, HTTP, EMAIL

//EFFICIENCY IMPROVEMENT: Don't check the same MarkdownLink twice
//Formatting the JSON to one check per line would be neat for readability

type MarkdownLink struct {
	File        string
	Name        string
	Destination string
	Type        string
	Status      string
}

func (m *MarkdownLink) IsHTTP() bool {
	_, err := url.ParseRequestURI(m.Destination)
	if err == nil {
		return true
	}
	return false
}

func (m *MarkdownLink) IsFile() bool {
	m.Type = "File"
	if strings.HasPrefix(m.Destination, ".") || strings.HasPrefix(m.Destination, "/") {
		return true
	}
	return false
}

func (m *MarkdownLink) IsEmail() bool {
	if strings.HasPrefix(m.Destination, "mailto") {
		return true
	}
	return false
}

//Is this really necessary? These will be matched by files
func (m *MarkdownLink) IsMarkdownHeaderLink() bool {
	//I'm not too sure on the second condition. Probably maybe. Doublecheck or remove this comment
	if strings.HasPrefix(m.Destination, "#") || strings.Contains(m.Destination, "#") {
		return true
	}
	return false
}

//Don't forge to handle links in the docs that you may have to login into first (see if you can handle a forbidden or something)
//Double check all HTTP codes one by one to make sure you're handling them correctly
func (m *MarkdownLink) CheckHTTP() {
	m.Type = "HTTP"
	resp, err := http.Head(m.Destination)
	if err != nil {
		m.Status = "BROKEN"
		return
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 208 || resp.StatusCode >= 300 && resp.StatusCode <= 308 || resp.StatusCode == 405 {
		m.Status = "OK"
	} else {
		m.Status = "BROKEN"
	}
}

func (m *MarkdownLink) CheckFile() {
	mDestination := filepath.Dir(m.File) + string(os.PathSeparator) + m.Destination

	//Still can't check things like: /app_management/secrets_and_configmaps.md#secrets-from-files (yet!)
	if strings.HasPrefix(m.Destination, "#") || strings.Contains(m.Destination, "#") {
		m.Status = "NOT IMPLEMENTED"
		return
	}

	_, err := os.Stat(mDestination)
	if os.IsNotExist(err) {
		m.Status = "BROKEN"
	} else {
		m.Status = "OK"
	}

	//
	////you need to handle this error stop being lazy yo
	//matched, _ := regexp.Match(`[A-Za-z0-9_]+\..+`, []byte(l.Destination))
	//
	//if matched {
	//	l.Type = "FILE"
	//	lDestination := filepath.Dir(l.File) + string(os.PathSeparator) + l.Destination
	//	if _, err := os.Stat(lDestination); os.IsNotExist(err) {
	//		l.Status = "BROKEN"
	//	} else {
	//		l.Status = "OK"
	//	}
	//	out <- l
	//	continue
	//}
}

func (m *MarkdownLink) CheckEmail() {
	m.Type = "EMAIL"
	m.Status = "NOT IMPLEMENTED"
}

func (m *MarkdownLink) CheckLink() {
	switch {
	case m.IsHTTP():
		m.CheckHTTP()
	case m.IsFile():
		m.CheckFile()
	case m.IsEmail():
	//case m.IsMarkdownHeaderLink():
	default:
		m.Type = "UNKNOWN"
		m.Status = "NOT IMPLEMENTED"
	}
}

//If ran on a Unix system:
//If this function receives /tmp it will return /tmp/
//If it receives /tmp/ it will return /tmp/
func CheckAndAddPathSeparatorSuffix(fsPath string) string {
	if !strings.HasSuffix(fsPath, string(os.PathSeparator)) {
		fsPath = fsPath + string(os.PathSeparator)
		return fsPath
	} else {
		return fsPath
	}
}

func PrintAndPanic(s string, err error) {
	fmt.Print(s)
	panic(err)
}

//returns a string containing the filesystem path for the cloned directory on disk
//hey is GitRepository a type?
//only works with http(s). Where's the SSH support?
//is returning a string the right thing to do?

//divide the contents of this function into two subfunctions:
//CloneGitRepository
//UpdateGitRepository

func GetGitRepository(repositoryUrl, tmpDir string) (string, error) {
	tmpDir = CheckAndAddPathSeparatorSuffix(tmpDir)
	url, err := url.ParseRequestURI("http://github.com/kubernetes/kubectl")
	if err != nil {
		return "", err
	}

	repoFilesystemPath := tmpDir + url.Path
	_, fsErr := os.Stat(repoFilesystemPath)

	if os.IsNotExist(fsErr) {
		log.Println("Cloning", repositoryUrl)
		_, err := git.PlainClone(repoFilesystemPath, false, &git.CloneOptions{
			URL:      repositoryUrl,
			Progress: os.Stdout,
		})
		if err != nil {
			return "", err
		}
	} else if fsErr != nil {
		return repoFilesystemPath, fsErr
	}

	repository, err := git.PlainOpen(repoFilesystemPath)
	if err != nil {
		return "", err
	}

	workTree, err := repository.Worktree()
	if err != nil {
		return "", err
	}

	log.Println("Pulling", repositoryUrl)
	err = workTree.Pull(&git.PullOptions{RemoteName: "origin"})
	if err == git.NoErrAlreadyUpToDate {
		return repoFilesystemPath, nil
	} else {
		return "", err
	}
}

func GetMarkdownFilepaths(repoFilesystemPath string) []string {
	var MarkdownFilepaths []string

	err := filepath.Walk(repoFilesystemPath, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), ".md") {
			MarkdownFilepaths = append(MarkdownFilepaths, path)
		}
		return err
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", err)
	}
	return MarkdownFilepaths
}

////leave some big comments explaining the regexes in this function
//this can still be improved, the regex could go on their own functions and run in parallel
func GetMarkdownLinksFromFiles(filePaths []string) ([]MarkdownLink, error) {

	var markdownLinks []MarkdownLink

	for _, f := range filePaths {
		fileContents, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}

		//GetInlineLinks
		//regex for footnote style MarkdownLinks
		re := regexp.MustCompile(`(\[.+\])\s*:\s*(.+)`)
		for _, matchedMarkdownLink := range re.FindAllStringSubmatch(string(fileContents), -1) {
			mdLink := MarkdownLink{
				File:        f,
				Name:        matchedMarkdownLink[1],
				Destination: matchedMarkdownLink[2],
			}
			markdownLinks = append(markdownLinks, mdLink)
		}

		//GetFootnoteLinks
		//regex for inline style links
		re = regexp.MustCompile(`(\[.+?\])((\()(.+?)(\)))`)
		for _, matchedMarkdownLink := range re.FindAllStringSubmatch(string(fileContents), -1) {
			mdLink := MarkdownLink{
				File:        f,
				Name:        matchedMarkdownLink[1],
				Destination: matchedMarkdownLink[4],
			}
			markdownLinks = append(markdownLinks, mdLink)
		}
	}

	return markdownLinks, nil
}

//This function is an absolute disaster
//Let's fix it later: https://about.sourcegraph.com/go/idiomatic-go
//Could you also do this with switch case?
//clean code has good advice on how to fix this too
//exclude the following strings: Changelogs, minutes
//func CheckMarkdownLink(markdownLink <-chan MarkdownLink, repoFilesystemPath string, workerNum int) chan MarkdownLink {
//	out := make(chan markdownLink)
//	var wg sync.WaitGroup
//	wg.Add(workerNum)
//
//	go func() {
//		for w := 0; w < workerNum; w++ {
//			go func() {
//				defer wg.Done()
//				for l := range MarkdownLink {
//					if strings.Contains(strings.ToLower(l.File), "changelog") || strings.Contains(strings.ToLower(l.File), "minute") || strings.Contains(strings.ToLower(l.File), "meeting") || strings.Contains(strings.ToLower(l.File), "release") {
//						continue
//					}
//
//					//you need to handle this error stop being lazy yo
//					matched, _ := regexp.Match(`[A-Za-z0-9_]+\..+`, []byte(l.Destination))
//
//					if matched {
//						l.Type = "FILE"
//						lDestination := filepath.Dir(l.File) + string(os.PathSeparator) + l.Destination
//						if _, err := os.Stat(lDestination); os.IsNotExist(err) {
//							l.Status = "BROKEN"
//						} else {
//							l.Status = "OK"
//						}
//						out <- l
//						continue
//					}
//
//					l.Type = "UNKNOWN"
//					l.Status = "NOTAPPLICABLE"
//					out <- l
//				}
//			}()
//		}
//		wg.Wait()
//		close(out)
//	}()
//	return out
//}			continue
//					}
//
//					//how about a timeout?
//					if strings.HasPrefix(l.Destination, "http") {
//						l.Type = "HTTP"
//						resp, err := http.Head(l.Destination)
//						if err != nil {
//							l.Status = "BROKEN"
//							out <- l
//							continue
//						}
//
//						if resp.StatusCode == 200 || resp.StatusCode == 301 {
//							l.Status = "OK"
//						} else {
//							l.Status = "BROKEN"
//						}
//						out <- l
//						continue
//					}
//
//					if strings.Contains(l.Destination, "#") {
//						continue
//					}
//
//					if strings.HasPrefix(l.Destination, ".") {
//						l.Type = "FILE"
//						lDestination := filepath.Dir(l.File) + string(os.PathSeparator) + l.Destination
//						if _, err := os.Stat(lDestination); os.IsNotExist(err) {
//							l.Status = "BROKEN"
//						} else {
//							l.Status = "OK"
//						}
//						out <- l
//						continue
//					}
//
//					if strings.HasPrefix(l.Destination, "/") {
//						l.Type = "FILE"
//						lDestination := repoFilesystemPath + l.Destination
//						if _, err := os.Stat(lDestination); os.IsNotExist(err) {
//							l.Status = "BROKEN"
//						}
//						l.Status = "OK"
//						out <- l
//						continue
//					}
//
//					//you need to handle this error stop being lazy yo
//					matched, _ := regexp.Match(`[A-Za-z0-9_]+\..+`, []byte(l.Destination))
//
//					if matched {
//						l.Type = "FILE"
//						lDestination := filepath.Dir(l.File) + string(os.PathSeparator) + l.Destination
//						if _, err := os.Stat(lDestination); os.IsNotExist(err) {
//							l.Status = "BROKEN"
//						} else {
//							l.Status = "OK"
//						}
//						out <- l
//						continue
//					}
//
//					l.Type = "UNKNOWN"
//					l.Status = "NOTAPPLICABLE"
//					out <- l
//				}
//			}()
//		}
//		wg.Wait()
//		close(out)
//	}()
//	return out
//}

////plz rename this function
//func ContainsDestination(ll []MarkdownLink, l MarkdownLink) bool {
//	for _, a := range ll {
//		if a.Destination == l.Destination {
//			return true
//		}
//	}
//	return false
//}
//
//func handler(w http.ResponseWriter, req *http.Request) {
//	///tmp/mdscanner/results/kubernetes/kubelet/
//
//	jsonFile, err := os.Open(config.tmpDir + "/results" + req.URL.Path + "/report.json")
//	if err != nil {
//		fmt.Fprintf(w, "Could not open json at: "+config.tmpDir+"/results"+req.URL.Path+"/report.json")
//	}
//	byteValue, _ := ioutil.ReadAll(jsonFile)
//
//	fmt.Fprintf(w, string(byteValue))
//
//}

//Implement logging!

func CheckMarkdownLinksWorker(mdLinkIn <-chan MarkdownLink, workerNum int) <-chan MarkdownLink {
	mdLinkOut := make(chan MarkdownLink)
	var wg sync.WaitGroup

	wg.Add(workerNum)
	go func() {
		for i := 0; i < workerNum; i++ {
			go func() {
				defer wg.Done()
				for mdLink := range mdLinkIn {
					mdLink.CheckLink()
					mdLinkOut <- mdLink
				}
			}()
		}
		wg.Wait()
		close(mdLinkOut)
	}()
	return mdLinkOut
}

//Name things properly in this function
//Leave a comment or two in here
func CheckMarkdownLinks(mdLinks []MarkdownLink) {
	linkChan := make(chan MarkdownLink)
	go func() {
		for _, link := range mdLinks {
			linkChan <- link
		}
		close(linkChan)
	}()

	myLinks := CheckMarkdownLinksWorker(linkChan, 30)

	for n := range myLinks {
		if n.Type == "File" {
			fmt.Println(n)
		}
	}
}

//launch 4 workers
//go worker(chan link, waitgroup)

//for _, link := range mdLinks {
//	//put link in channel
//}

//wait for workers to finish

//switch {
//case link.isHTTP():
//	fmt.Println(link.Destination, " is a HTTP link")
//default:
//	fmt.Println(link.Destination, "is NOT a HTTP link")
//}

//
//link.Type = "HTTP"
//resp, err := http.Head(link.Destination)
//if err != nil {
//	link.Status = "BROKEN"
//}
//
//if resp.StatusCode == 200 || resp.StatusCode == 301 {
//	link.Status = "OK"
//} else {
//	link.Status = "BROKEN"
//}

//Think about your queueing package, do a queue.Start() at the beginning and have other things send things to it maybe!
//Don't forge to find a way to implement a list of things you want to ignore: if strings.Contains(strings.ToLower(l.File), "changelog") || strings.Contains(strings.ToLower(l.File), "minute") || strings.Contains(strings.ToLower(l.File), "meeting") || strings.Contains(strings.ToLower(l.File), "release") {
func main() {

	log.SetOutput(os.Stdout)
	//GitRepository := "https://github.com/kubernetes/kubectl"
	//tmpDir := "/tmp"

	//gitRepoFilesystemPath, err := GetGitRepository(GitRepository, tmpDir)
	//if err != nil {
	//	PrintAndPanic("Error running GetGitRepository:", err)
	//}

	//markdownFilepaths := GetMarkdownFilepaths(gitRepoFilesystemPath)
	markdownFilepaths := GetMarkdownFilepaths("/tmp/kubernetes/kubectl/")

	var markdownLinks []MarkdownLink

	markdownLinks, err := GetMarkdownLinksFromFiles(markdownFilepaths)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(markdownLinks)

	CheckMarkdownLinks(markdownLinks)

}

//for _, projectRepo := range projectRepos {
//	repoFilesystemPath := GetGitRepository(projectRepo, config.tmpDir)
//	MarkdownFilepaths := GetMarkdownFilepaths(repoFilesystemPath)
//	MarkdownLinks := GetMarkdownLinksFromFile(MarkdownFilepaths)
//	//results := CheckMarkdownLink(MarkdownLinks, repoFilesystemPath, 10)
//
//	//trying to read from a channel and appending it to a slice tries to close an already closed channel
//	var ll []MarkdownLink
//	//Only append duplicate MarkdownLinks for broken MarkdownLinks
//	for r := range CheckMarkdownLink(MarkdownLinks, repoFilesystemPath, 10) {
//		log.Println(r)
//
//		//if it's broken, you always want to append
//		if r.Status == "BROKEN" {
//			ll = append(ll, r)
//			//otherwise only append if the MarkdownLink is not on the list already (aka no duplicates)
//		} else if !ContainsDestination(ll, r) {
//			ll = append(ll, r)
//		}
//	}
//
//	u, err := url.Parse(projectRepo)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//u.EscapedPath() returns /kubernetes/kubectl
//
//	jsonSavePath := config.tmpDir + "/results" + u.EscapedPath()
//
//	err = os.MkdirAll(jsonSavePath, os.ModePerm)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	file, _ := json.MarshalIndent(ll, "", "")
//	_ = ioutil.WriteFile(jsonSavePath+"/report.json", file, 0644)
//	log.Println("Report saved at: ", jsonSavePath+"/report.json")
//
//}
//
//http.HandleFunc("/", handler)
//log.Fatal(http.ListenAndServe(":8080", nil))
