package main

//
//import (
//	"net/http"
//	"os"
//	"path/filepath"
//	"regexp"
//	"strings"
//	"sync"
//)
//
//func CheckMarkdownLink_backup(MarkdownLink <-chan MarkdownLink, repoFilesystemPath string, workerNum int) chan MarkdownLink {
//	out := make(chan MarkdownLink)
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
//					if strings.HasPrefix(l.Destination, "mailto") {
//						continue
//					}
//
//					if strings.HasPrefix(l.Destination, "#") {
//						continue
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
