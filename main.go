package main

import (
	"fmt"
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
	app.Action = Action
	app.HideHelp = true
	app.EnableBashCompletion = true
	app.BashComplete = BashComplete
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "shellinit",
		},
	}
	app.Run(os.Args)
}

func Action(c *cli.Context) {
	if c.Bool("shellinit") {
		fmt.Println(SrcPath + "github.com/vinceprignano/gocd/shellinit")
		return
	}
	if len(c.Args()) > 1 {
		fmt.Println(c.Args())
		return
	}

	// Calling just gocd goes into src
	if len(c.Args()) == 0 {
		fmt.Println(SrcPath)
		return
	}

	// Walk on the GOPATH
	filepath.Walk(SrcPath, VisitAndStore)

	// Get the chosen parameter
	chosen := c.Args()[0]

	// If the argument is a path inside the GOPATH
	if strings.Contains(chosen, "/") {
		path := SrcPath + chosen
		DirectoryExists(path)
		fmt.Println(path)
		return
	}

	// If the argument contains a selector
	if strings.Contains(chosen, ":") {
		split := strings.Split(chosen, ":")
		path := split[0]
		DirectoryExists(path)

		// Get the index
		index, _ := strconv.Atoi(split[1])
		if index > len(PathMap[path])-1 {
			fmt.Printf("Path with index %d not found for %s", index, path)
			os.Exit(1)
		}
		fmt.Println(PathMap[path][index])
		return
	}

	// If the chosen key has multiple paths associated
	// print the possible options
	if len(PathMap[chosen]) > 1 {
		fmt.Printf("Found multiple folders with this name:\n")
		for i, m := range PathMap[chosen] {
			fmt.Printf("\t%d: %s\n", i, m)
		}
		os.Exit(1)
		return
	}

	// Check the directory and print the path
	DirectoryExists(PathMap[chosen][0])
	fmt.Println(PathMap[chosen][0])
}
