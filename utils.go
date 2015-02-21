package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/codegangsta/cli"
)

func VisitAndStore(fp string, fi os.FileInfo, err error) error {
	if !fi.IsDir() || strings.Contains(fp, "/.") {
		return nil
	}
	pathSplit := strings.Split(fp, "/")
	w := pathSplit[len(pathSplit)-1]
	PathMap[w] = append(PathMap[w], fp)
	return nil
}

func VisitDir(fp string, fi os.FileInfo, err error) error {
	if !fi.IsDir() || strings.Contains(fp, "/.") {
		return nil
	}
	path := strings.Replace(fp, SrcPath, "", 1)
	fmt.Println(path)
	pathSplit := strings.Split(path, "/")
	fmt.Println(pathSplit[len(pathSplit)-1])
	return nil
}

func BashComplete(c *cli.Context) {
	if len(c.Args()) > 1 {
		return
	}
	filepath.Walk(SrcPath, VisitDir)
}

func DirectoryExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	panic(err)
}
