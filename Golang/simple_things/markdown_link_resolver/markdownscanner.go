package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

//This is unfinished^W bad and I feel bad
//FIX LINE 105 PLZ

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
//take repoUrl as an arguement?
//if url doesn't begin by http, throw error!
//you need to add a way to ignore certain links (minutes, slack links, mailto, etc)

//to allow only certain types to be assigned, maybe make something of a const type somehow? https://blog.golang.org/constants
//statuses: OK, BROKEN, NOTAPPLICABLE, UNKNOWN
//type: FILE, HTTP, EMAIL

//EFFICIENCY IMPROVEMENT: Don't check the same link twice
//Formatting the JSON to one check per line would be neat for readability

type Link struct {
	SrcFile     string
	Name        string
	Destination string
	Type        string
	Status      string
}

type Config struct {
	TmpDir string
}

var config = Config{
	TmpDir: "/tmp/mdscanner",
}

//returns a string containing the filesystem path for the cloned directory on disk
func GetRepo(repoUrl, tmpDir string) string {
	log.Println(repoUrl, tmpDir)

	if !strings.HasPrefix(repoUrl, "http") {
		log.Fatalln(repoUrl, "does not begin with http, does not appear valid URL")
	}

	//will split an URL and give you the last element
	_, repoName := path.Split(repoUrl)
	repoPath := tmpDir + "/repos" + string(os.PathSeparator) + repoName

	//If the repository does not exist on the local filesystem, we clone it
	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		log.Println(tmpDir + string(os.PathSeparator) + repoName + " does not exist, cloning")
		_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
			URL:      repoUrl,
			Progress: os.Stdout,
		})
		if err != nil {
			log.Fatal(err)
		}

		//if the repository already exists in the local filesystem, we pull the latest changes
	} else if _, err := os.Stat(repoPath); err == nil {
		log.Println("Repository already exists on filesystem, pulling latest changes")

		r, err := git.PlainOpen(repoPath)
		if err != nil {
			log.Fatal(err)
		}

		w, err := r.Worktree()
		if err != nil {
			log.Fatal(err)
		}

		err = w.Pull(&git.PullOptions{RemoteName: "origin"})

		//The pull function above returns an error if the repository is already up to date
		//We're not concerned with that error, but throw fatal if we get another error

		//This gives a nil pointer oftentimes and crashes, redo this
		if err.Error() == "already up-to-date" {
			log.Println("repo up to date, all gucci fam")
		} else if err != nil {
			log.Fatal(err)
		}
	}
	return repoPath
}

func GetMarkdownFiles(repoPath string) <-chan string {

	out := make(chan string)

	go func() {
		err := filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			}

			if strings.HasSuffix(info.Name(), ".md") {
				out <- path
			}
			return err
		})

		if err != nil {
			fmt.Printf("error walking the path %q: %v\n", err)
		}
		close(out)
	}()

	return out
}

//leave some big comments explaining the regexes in this function
func GetLinksFromFile(filePath <-chan string) <-chan Link {

	out := make(chan Link)

	go func() {
		for f := range filePath {
			data, err := ioutil.ReadFile(f)
			if err != nil {
				log.Panic(err)
			}

			//regex for footnote style links
			re := regexp.MustCompile(`(\[.+\])\s*:\s*(.+)`)

			for _, matchedLink := range re.FindAllStringSubmatch(string(data), -1) {
				l := Link{
					SrcFile:     f,
					Name:        matchedLink[1],
					Destination: matchedLink[2],
				}
				out <- l
			}

			//regex for inline links
			//https://stackoverflow.com/questions/6208367/regex-to-match-stuff-between-parentheses
			re = regexp.MustCompile(`(\[.+?\])((\()(.+?)(\)))`)

			for _, matchedLink := range re.FindAllStringSubmatch(string(data), -1) {
				l := Link{
					SrcFile:     f,
					Name:        matchedLink[1],
					Destination: matchedLink[4],
				}
				out <- l
			}
		}
		close(out)
	}()
	return out
}

//This function has several anti-patterns. Rob, look away now...
//Let's fix it later: https://about.sourcegraph.com/go/idiomatic-go
//Could you also do this with switch case?
//exclude the following strings: Changelogs, minutes
func CheckLink(link <-chan Link, repoPath string, workerNum int) chan Link {
	out := make(chan Link)
	var wg sync.WaitGroup
	wg.Add(workerNum)

	go func() {
		for w := 0; w < workerNum; w++ {
			go func() {
				defer wg.Done()
				for l := range link {
					if strings.Contains(strings.ToLower(l.SrcFile), "changelog") || strings.Contains(strings.ToLower(l.SrcFile), "minute") || strings.Contains(strings.ToLower(l.SrcFile), "meeting") || strings.Contains(strings.ToLower(l.SrcFile), "release") {
						continue
					}

					if strings.HasPrefix(l.Destination, "mailto") {
						continue
					}

					if strings.HasPrefix(l.Destination, "#") {
						continue
					}

					//how about a timeout?
					if strings.HasPrefix(l.Destination, "http") {
						l.Type = "HTTP"
						resp, err := http.Head(l.Destination)
						if err != nil {
							l.Status = "BROKEN"
							out <- l
							continue
						}

						if resp.StatusCode == 200 || resp.StatusCode == 301 {
							l.Status = "OK"
						} else {
							l.Status = "BROKEN"
						}
						out <- l
						continue
					}

					if strings.Contains(l.Destination, "#") {
						continue
					}

					if strings.HasPrefix(l.Destination, ".") {
						l.Type = "FILE"
						lDestination := filepath.Dir(l.SrcFile) + string(os.PathSeparator) + l.Destination
						if _, err := os.Stat(lDestination); os.IsNotExist(err) {
							l.Status = "BROKEN"
						} else {
							l.Status = "OK"
						}
						out <- l
						continue
					}

					if strings.HasPrefix(l.Destination, "/") {
						l.Type = "FILE"
						lDestination := repoPath + l.Destination
						if _, err := os.Stat(lDestination); os.IsNotExist(err) {
							l.Status = "BROKEN"
						}
						l.Status = "OK"
						out <- l
						continue
					}

					//you need to handle this error stop being lazy yo
					matched, _ := regexp.Match(`[A-Za-z0-9_]+\..+`, []byte(l.Destination))

					if matched {
						l.Type = "FILE"
						lDestination := filepath.Dir(l.SrcFile) + string(os.PathSeparator) + l.Destination
						if _, err := os.Stat(lDestination); os.IsNotExist(err) {
							l.Status = "BROKEN"
						} else {
							l.Status = "OK"
						}
						out <- l
						continue
					}

					l.Type = "UNKNOWN"
					l.Status = "NOTAPPLICABLE"
					out <- l
				}
			}()
		}
		wg.Wait()
		close(out)
	}()
	return out
}

//plz rename this function
func ContainsDestination(ll []Link, l Link) bool {
	for _, a := range ll {
		if a.Destination == l.Destination {
			return true
		}
	}
	return false
}

func handler(w http.ResponseWriter, req *http.Request) {
	///tmp/mdscanner/results/kubernetes/kubelet/

	jsonFile, err := os.Open(config.TmpDir + "/results" + req.URL.Path + "/report.json")
	if err != nil {
		fmt.Fprintf(w, "Could not open json at: "+config.TmpDir+"/results"+req.URL.Path+"/report.json")
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	fmt.Fprintf(w, string(byteValue))

}

func main() {

	log.SetOutput(os.Stdout)
	projectRepos := []string{

		"https://github.com/kubernetes/kubelet",
		//"https://github.com/kubernetes/kops",
		//"https://github.com/kubernetes/kubeadm",
	}

	for _, projectRepo := range projectRepos {
		repoPath := GetRepo(projectRepo, config.TmpDir)
		markdownFiles := GetMarkdownFiles(repoPath)
		links := GetLinksFromFile(markdownFiles)
		//results := CheckLink(links, repoPath, 10)

		//trying to read from a channel and appending it to a slice tries to close an already closed channel
		var ll []Link
		//Only append duplicate links for broken links
		for r := range CheckLink(links, repoPath, 10) {
			log.Println(r)

			//if it's broken, you always want to append
			if r.Status == "BROKEN" {
				ll = append(ll, r)
				//otherwise only append if the link is not on the list already (aka no duplicates)
			} else if !ContainsDestination(ll, r) {
				ll = append(ll, r)
			}
		}

		u, err := url.Parse(projectRepo)
		if err != nil {
			log.Fatal(err)
		}

		//u.EscapedPath() returns /kubernetes/kubectl

		jsonSavePath := config.TmpDir + "/results" + u.EscapedPath()

		err = os.MkdirAll(jsonSavePath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		file, _ := json.MarshalIndent(ll, "", "")
		_ = ioutil.WriteFile(jsonSavePath+"/report.json", file, 0644)
		log.Println("Report saved at: ", jsonSavePath+"/report.json")

	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
