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

//COMMENTS EVERYWHERE
//TESTS!
//take repo as parameter
//take branch as parameter
//handle invalid branch and repos appropriately

//define basedir as where to clone! (anything other than /tmp for instance)
//take repoUrl as an arguement?
//if url doesn't begin by http, throw error!
//you need to add a way to ignore certain links (minutes, slack links, mailto, etc)

type Link struct {
	FilePath        string
	LinkName        string
	LinkDestination string
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
		log.Println("/tmp/community does not exist, cloning")
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

func GetMarkdownFiles(repoPath string) []string {

	var markdownFiles []string

	err := filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
		}

		if strings.HasSuffix(info.Name(), ".md") {
			markdownFiles = append(markdownFiles, path)
		}
		return err
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", err)
	}
	return markdownFiles
}

//leave some big comments explaining the regexes in this function
func GetLinksFromFile(filePath string) []Link {

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic(err)
	}

	var links []Link

	//regex for footnote style links
	re := regexp.MustCompile(`(\[.+\])\s*:\s*(.+)`)

	for _, matchedLink := range re.FindAllStringSubmatch(string(data), -1) {
		l := Link{
			FilePath:        filePath,
			LinkName:        matchedLink[1],
			LinkDestination: matchedLink[2],
		}
		links = append(links, l)
	}

	//regex for inline links
	//https://stackoverflow.com/questions/6208367/regex-to-match-stuff-between-parentheses
	re = regexp.MustCompile(`(\[.+?\])((\()(.+?)(\)))`)

	for _, matchedLink := range re.FindAllStringSubmatch(string(data), -1) {
		l := Link{
			FilePath:        filePath,
			LinkName:        matchedLink[1],
			LinkDestination: matchedLink[4],
		}
		links = append(links, l)
	}

	return links
}

func CheckLink(link Link) {

	//if strings.HasPrefix(url,"http"){
	//	resp, err := http.Get(url)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	log.Println(resp.StatusCode)
	//}

	//Find a way to check local filesystem links
	//Find a way to check http links
	//Implement parallel logic
	//Implement items to skip (mailto, links to headers)

	if strings.HasPrefix(link.LinkDestination, "mailto") {
		return
	}

	if strings.HasPrefix(link.LinkDestination, "#") {
		return
	}

	//idea:
	//make channel
	//go func
	//push into channel

	//go func
	//more go funcs for parallelism
	//retrieve from channel
	//wait with waitgroups

	if strings.HasPrefix(link.LinkDestination, "http") {
		resp, err := http.Head(link.LinkDestination)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(link)
		log.Println(resp.StatusCode)
	}

	//	log.Println(link)

	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Println(string(body))
}

func main() {

	log.SetOutput(os.Stdout)

	repoPath := GetRepo("https://github.com/kubernetes/community", "/tmp")
	markdownFiles := GetMarkdownFiles(repoPath)
	//
	for _, v := range markdownFiles {
		for _, v1 := range GetLinksFromFile(v) {

			CheckLink(v1)

		}
	}

}

//if directory exists and is repo
//git checkout

//if directory

//if path

//
//if err != nil {
//	fmt.Println(err)
//}
//
//err := filepath.Walk("/tmp/go-git", func(path string, info os.FileInfo, err error) error {
//	if err != nil {
//		fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
//		return err
//	}
//
//	if strings.HasSuffix(info.Name(), ".md"){
//
//		getLinkFromMD(path)
//	}
//
//	return nil
//
//})
//if err != nil {
//	fmt.Printf("error walking the path %q: %v\n", err)
//	return
//}
