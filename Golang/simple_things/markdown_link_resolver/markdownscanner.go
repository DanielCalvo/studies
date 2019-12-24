package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

//This is unfinished
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

type Link struct {
	SrcFile     string
	Name        string
	Destination string
	Type        string
	Status      string
}

//returns a string containing the filesystem path for the cloned directory on disk
func GetRepo(repoUrl, tmpDir string) string {
	fmt.Println(repoUrl, tmpDir)

	if !strings.HasPrefix(repoUrl, "http") {
		log.Fatalln(repoUrl, "does not begin with http, does not appear valid URL")
	}

	//will split an URL and give you the last element
	_, repoName := path.Split(repoUrl)
	repoPath := tmpDir + string(os.PathSeparator) + repoName

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

//idea:
//make channel
//go func
//push into channel

//go func
//more go funcs for parallelism
//retrieve from channel
//wait with waitgroups

//This function has several anti-patterns. Rob, look away now...
//Let's fix it later: https://about.sourcegraph.com/go/idiomatic-go
//Could you also do this with switch case?
func CheckLink(link <-chan Link, repoPath string, routineNum int) <-chan Link {
	out := make(chan Link)
	go func() {
		for l := range link {
			if strings.HasPrefix(l.Destination, "mailto") {
				l.Type = "EMAIL"
				l.Status = "NOTAPPLICABLE"
				out <- l
				continue
			}

			if strings.HasPrefix(l.Destination, "#") {
				l.Type = "HEADER"
				l.Status = "NOTAPPLICABLE"
				out <- l
				continue
			}

			if strings.HasPrefix(l.Destination, "http") {
				l.Type = "HTTP"
				resp, err := http.Head(l.Destination)
				if err != nil {
					log.Fatalln(err)
				}

				if resp.StatusCode == 200 {
					l.Status = "OK"
				} else {
					l.Status = "BROKEN"
				}
				out <- l
				continue
			}

			if strings.Contains(l.Destination, "#") {
				//You also need to handle ls with destinations such as:
				// /contributors/devel/README.md#setting-up-your-dev-environment-coding-and-debugging
				//Check the file, but not the header (for now)
				l.Type = "HEADER"
				l.Status = "NOTAPPLICABLE"
				out <- l
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
		close(out)
	}()

	return out
}

func main() {

	log.SetOutput(os.Stdout)

	repoPath := GetRepo("https://github.com/kubernetes/kubectl", "/tmp")
	markdownFiles := GetMarkdownFiles(repoPath)
	links := GetLinksFromFile(markdownFiles)
	results := CheckLink(links, repoPath, 50)

	var ll []Link
	//trying to read from a channel and appending it to a slice tries to close an already closed channel
	for r := range results {
		log.Println(r)
		ll = append(ll, r)

	}

	log.Println(ll)
	//
	//file, _ := json.MarshalIndent(ll, "", " ")
	//_ = ioutil.WriteFile("/tmp/test1.json", file, 0644)

}
