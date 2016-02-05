package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/acsellers/inflections"
	"github.com/codegangsta/cli"
	"github.com/coreos/go-semver/semver"
	"github.com/fatih/color"
	De "github.com/tj/go-debug"
)

var debug = De.Debug("governator:main")

type templateVars struct {
	ProjectName string
	PROJECTNAME string
}

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

	vars := templateVars{
		ProjectName: name,
		PROJECTNAME: strings.ToUpper(inflections.Underscore(name)),
	}

	fmt.Printf("name: %v", name)
	for _, assetName := range AssetNames() {
		fmt.Printf("assetName: %v\n", assetName)
		asset := parseTemplate(vars, assetName)
		fmt.Println(asset)
	}
}

func parseTemplate(vars templateVars, assetName string) string {
	asset := MustAsset(assetName)
	assetStr := string(asset)
	assetTemplate, err := template.New(assetName).Parse(assetStr)
	PanicIfError("template.Parse failed", err)

	buffer := bytes.NewBufferString("")
	assetTemplate.Execute(buffer, vars)

	return buffer.String()
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

// PanicIfError prints error and dies if error is non nil
func PanicIfError(msg string, err error) {
	if err == nil {
		return
	}

	log.Panicf("ERROR(%v):\n\n%v", msg, err)
}
