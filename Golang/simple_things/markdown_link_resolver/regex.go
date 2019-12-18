package main

//leave comprehensive comments
//create tests
//you still need to support inline links

//type Link struct {
//	FilePath        string
//	LinkName        string
//	LinkDestination string
//}

//leave some big comments explaining the regexes in this function
//func GetLinksFromFile(filePath string) []Link {
//
//	data, err := ioutil.ReadFile(filePath)
//	if err != nil {
//		log.Panic(err)
//	}
//
//	var links []Link
//
//	//regex for footnote style links
//	re := regexp.MustCompile(`(\[.+\])\s*:\s*(.+)`)
//
//	for _, matchedLink := range re.FindAllStringSubmatch(string(data), -1) {
//		l := Link{
//			FilePath:        filePath,
//			LinkName:        matchedLink[1],
//			LinkDestination: matchedLink[2],
//		}
//		links = append(links, l)
//	}
//
//	//regex for inline links
//	//https://stackoverflow.com/questions/6208367/regex-to-match-stuff-between-parentheses
//	re = regexp.MustCompile(`(\[.+?\])(\(.+?\))`)
//
//	for _, matchedLink := range re.FindAllStringSubmatch(string(data), -1) {
//		l := Link{
//			FilePath:        filePath,
//			LinkName:        matchedLink[1],
//			LinkDestination: matchedLink[2],
//		}
//		links = append(links, l)
//	}
//
//	return links
//}

func main() {

}
