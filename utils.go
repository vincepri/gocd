package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

var Cache map[string]time.Time
var CachePath string

func init() {
	Cache = make(map[string]time.Time)
	usr, _ := user.Current()
	CachePath = fmt.Sprintf("%s/.gocd", usr.HomeDir)
}

func LoadCache() {
	cache, _ := ioutil.ReadFile(CachePath)
	json.Unmarshal(cache, &Cache)
	b, _ := ioutil.ReadFile(CachePath + "_map")
	json.Unmarshal(b, &PathMap)
}

func StoreCache() {
	b, _ := json.Marshal(Cache)
	ioutil.WriteFile(CachePath, b, 0644)
	b, _ = json.Marshal(PathMap)
	ioutil.WriteFile(CachePath+"_map", b, 0644)
}

func VisitAndStore(fp string, fi os.FileInfo, err error) error {
	if !fi.IsDir() || strings.Contains(fp, "/.") || checkForUpdate(fp, fi.ModTime()) {
		return nil
	}
	Cache[fp] = time.Now().UTC()
	pathSplit := strings.Split(fp, "/")
	w := pathSplit[len(pathSplit)-1]
	PathMap[w] = append(PathMap[w], fp)
	return nil
}

func checkForUpdate(fp string, modTime time.Time) bool {
	if cached, ok := Cache[fp]; ok {
		if cached.Sub(modTime) < 0 {
			return true
		}
	}
	return false
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
