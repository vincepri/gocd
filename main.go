package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/codegangsta/cli"
)

var SrcPath string
var PathMap map[string][]string

func main() {
	SrcPath = os.Getenv("GOPATH") + "/src/"
	PathMap = make(map[string][]string)
	app := cli.NewApp()
	app.Name = "gocd"
	app.Usage = "Quick cd into a GOPATH directory"
	app.Action = func(c *cli.Context) {
		if c.Bool("shellinit") {
			fmt.Println(SrcPath + "github.com/vinceprignano/gocd/shellinit")
			return
		}
		if len(c.Args()) > 1 {
			fmt.Println(c.Args())
			return
		}
		filepath.Walk(SrcPath, visitAndStore)
		chosen := c.Args()[0]
		if strings.Contains(chosen, "/") {
			fmt.Println(SrcPath + chosen)
			return
		}
		if strings.Contains(chosen, ":") {
			split := strings.Split(chosen, ":")
			index, _ := strconv.Atoi(split[1])
			path := split[0]
			fmt.Println(PathMap[path][index])
			return
		}
		if len(PathMap[chosen]) > 1 {
			fmt.Printf("Found multiple folders with this name:\n")
			for i, m := range PathMap[chosen] {
				fmt.Printf("\t%d: %s\n", i, m)
			}
			os.Exit(1)
			return
		}
		if err := os.Chdir(PathMap[chosen][0]); err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(PathMap[chosen][0])
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "shellinit",
		},
	}
	app.HideHelp = true
	app.EnableBashCompletion = true
	app.BashComplete = func(c *cli.Context) {
		if len(c.Args()) > 1 {
			return
		}
		filepath.Walk(SrcPath, visitDir)
	}
	app.Run(os.Args)
}

func visitAndStore(fp string, fi os.FileInfo, err error) error {
	if !fi.IsDir() || strings.Contains(fp, "/.") {
		return nil
	}
	pathSplit := strings.Split(fp, "/")
	w := pathSplit[len(pathSplit)-1]
	PathMap[w] = append(PathMap[w], fp)
	return nil
}

func visitDir(fp string, fi os.FileInfo, err error) error {
	if !fi.IsDir() || strings.Contains(fp, "/.") {
		return nil
	}
	path := strings.Replace(fp, SrcPath, "", 1)
	fmt.Println(path)
	pathSplit := strings.Split(path, "/")
	fmt.Println(pathSplit[len(pathSplit)-1])
	return nil
}
