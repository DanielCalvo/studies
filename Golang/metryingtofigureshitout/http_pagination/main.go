package main

/*
lets use pagination to get all kubernetes issues, im interested in just getting every single page using pagination!

curl -L \
  -H "Accept: application/vnd.github+json" \
  -H "X-GitHub-Api-Version: 2026-03-10" \
  https://api.github.com/repos/kubernetes/kubernetes/issues


curl -L -H "Accept: application/vnd.github+json" -H "X-GitHub-Api-Version: 2026-03-10" https://api.github.com/repos/kubernetes/kubernetes/issues
curl -I -H "Accept: application/vnd.github+json" -H "X-GitHub-Api-Version: 2026-03-10" https://api.github.com/repos/kubernetes/kubernetes/issues

*/

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
//learnings:
- it can be useful to debug the program to see exactly what data you're getting from a third party API to validate that it behaves how you think it does and to validate your program is working the way you think it is!
- ask codex how to test adding functionality to the program once it gets big? dunno if you split it into small things or just write a function and then test it or something
*/

func main() {
	url := "https://api.github.com/repos/kubernetes/kubernetes/issues"
	//url := "https://api.github.com/repositories/20580498/issues?after=Y3Vyc29yOnYyOpLPAAABVoh0pOjOCjII8A%3D%3D&per_page=30&page=86" //let me start on page 86 to test the last tthing

	res, err := makeRequest(url)

	if err != nil {
		log.Fatal(err)
	}
	//reqHeader := res.Header.Get("Link")

	//we dont know for how many pages we have to loop for upfront, so looping until we get no more pages is an acceptable approach
	//the seen feature keeps us from trusting then functions that get the next page too much and stops the program in case we try to visit a page twice, breaking an infinite loop
	//We could add more and more functionality here to increase the safety of this function but codex shared a really interesting insight:
	//- learning code asks “how does this work?”
	//- production code asks “what is the safest thing I can reasonably maintain?”
	seen := map[string]bool{}

	for {
		fmt.Println("working with", url)
		url = getNextPageFromHeader(res.Header.Get("Link"))
		if url == "" {
			break
		}
		if seen[url] {
			log.Fatal("pagination loop detected")
		}
		seen[url] = true

		res, err = makeRequest(url)
		if err != nil {
			log.Fatal(err)
		}
	}

	//while header contains next or last, get next page
	//otherwise stop

}

// i need to study some dependency injection to figure out how to test this properly
func getIssuePages(url string) []string {
	var result []string
	seen := map[string]bool{}

	for {
		res, err := makeRequest(url)

		fmt.Println("working with", url)
		url = getNextPageFromHeader(res.Header.Get("Link"))
		if url == "" {
			break
		}
		if seen[url] {
			log.Fatal("pagination loop detected")
		}
		seen[url] = true

		res, err = makeRequest(url)
		if err != nil {
			log.Fatal(err)
		}
	}

	return result
}

// If this was in a production system I think we should use a package with a well established parser but for this learning exercise let's just go with this
// parsing by comma is ok for a learning-level solution
// when to split a function: is this logic one idea, or two ideas? in here we're trying to do only one thing: get next page from header, and its easy to read (for now) so its best to keep it as is
// small functions are not automatically better and can make the flow difficult to follow!
func getNextPageFromHeader(s string) string { // i dont like that this returns empty string and does not tell you why, maybe an error would be better?

	if !strings.Contains(s, "rel=\"next\"") {
		return ""
	}

	parts := strings.Split(s, ",")
	for _, part := range parts {
		if strings.Contains(part, "rel=\"next\"") {
			s = part
		}
	}

	start := strings.Index(s, "<")
	end := strings.Index(s, ">")

	//if start or end are not found in the string
	if start == -1 || end == -1 {
		return ""
	}
	//if the gt/lt signs are out of order (lt has to show first in the string, ex: <hello> and not >hello<
	if start > end {
		return ""
	}
	return s[start+1 : end]
	//how about the next or last?
	//how do you check if its a valid url?
	//what do you get on the last page?
}

func makeRequest(url string) (http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return http.Response{}, err
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("GH_TOKEN")) //pls remove this from the code before pushing to git
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return http.Response{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return http.Response{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return http.Response{}, err
	}
	//lets return a response for now and i'll figure out how to refine this later
	return *res, nil
}
