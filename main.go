package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-semver/semver"
	"github.com/fatih/color"
	De "github.com/tj/go-debug"
)

var debug = De.Debug("governator:main")

func main() {
	app := cli.NewApp()
	app.Name = "go-go"
	app.Version = version()
	app.ArgsUsage = "<project-name>"
	app.Action = run
	app.Flags = []cli.Flag{}
	app.Run(os.Args)
}

func run(context *cli.Context) {
	debug("run")
	name := getOpts(context)

	fmt.Println("name", name)
}

func getOpts(context *cli.Context) string {
	name := context.Args().First()

	if name == "" {
		cli.ShowAppHelp(context)
		color.Red("  Missing required arg <project-name>")
		os.Exit(1)
	}

	return name
}

func version() string {
	version, err := semver.NewVersion(VERSION)
	if err != nil {
		errorMessage := fmt.Sprintf("Error with version number: %v", VERSION)
		log.Panicln(errorMessage, err.Error())
	}
	return version.String()
}
