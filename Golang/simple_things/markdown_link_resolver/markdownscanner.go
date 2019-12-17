package main

import (
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//take repo as parameter
//take branch as parameter
//handle invalid branch and repos appropriately

func main() {

	//define basedir as where to clone! (anything other than /tmp for instance)
	//take repoUrl as an arguement?
	//if url doesn't begin by http, throw error!

	tmpDirectory := "/tmp/"

	repoUrl := "https://github.com/kubernetes/community"
	repoPath := "/tmp/community"

	if !strings.HasPrefix(repoUrl, "http") {
		log.Fatalln(repoUrl, "does not begin with http, does not appear valid URL")
	}

	dir, file := path.Split(repoUrl)

	fmt.Println(repoUrl + " - " + repoPath)
	fmt.Println(dir + " - " + file)

	os.Exit(0)

	//If the repository does not exist on the local filesystem, we clone it
	if _, err := os.Stat("/tmp/community"); os.IsNotExist(err) {
		log.Println("/tmp/community does not exist, cloning")
		_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
			URL:      repoUrl,
			Progress: os.Stdout,
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	//if the repository already exists in the local filesystem, we pull the latest changes
	if _, err := os.Stat(repoPath); err == nil {
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

		if err.Error() == "already up-to-date" {
			log.Println("repo up to date, all gucci fam")
		} else if err != nil {
			log.Fatal(err)
		}
	}

	err := filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if strings.HasSuffix(info.Name(), ".md") {
			fmt.Println(path + info.Name())
		}
		return nil
	})

	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", err)
		return
	}

	//you have to set the repo url here
	log.Println("Finished!")

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
