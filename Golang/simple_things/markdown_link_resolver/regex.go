package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

//leave comprehensive comments
//create tests
//you still need to support inline links

type Link struct {
	FilePath        string
	LinkName        string
	LinkDestination string
}

func GetLinksFromFile(filePath string) []Link {
	//You need to also handle inline links!
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic(err)
	}

	//leave some big comment explaining the regex
	re := regexp.MustCompile(`(\[.+\])\s*:\s*(.+)`)

	var links []Link

	for _, matchedLink := range re.FindAllStringSubmatch(string(data), -1) {
		l := Link{
			FilePath:        filePath,
			LinkName:        matchedLink[1],
			LinkDestination: matchedLink[2],
		}
		links = append(links, l)
	}
	return links
}

func main() {

	stuff := GetLinksFromFile("/home/daniel/tmp/community/community-membership.md")

	for _, v := range stuff {
		fmt.Println(v)
	}

	//maybe you don't even need link_name and link_addr?

	//Explain this for loop well:
	//[code reviews]: /contributors/guide/collab.md
	//[code reviews]
	// /contributors/guide/collab.md

	//result := make(map[string]string)
	//
	//for i, name := range re.SubexpNames() {
	//	if i != 0 && name != "" {
	//		result[name] = match[i]
	//	}
	//}

	//fmt.Printf("by name: %s %s\n", result["link_name"], result["link_addr"])

}
