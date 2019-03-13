package main

import (
	"fmt"
	"github.com/jawher/mow.cli"
	log "github.com/c2nc/protoc-go-inject-tag/logger"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const (
	AppName = "protoc-go-inject-tag"
	AppVer = "1.0.1"
	AppDescr = "Protobuf custom tags and validation"
)

func init() {
	log.SetLogLevel(EnvOrDefault("LOG_LEVEL", "error"))
}

func processFile(inputFile, xxxTags string) {
	var xxxSkipSlice []string

	if len(xxxTags) > 0 {
		xxxSkipSlice = strings.Split(xxxTags, ",")
	}

	if len(inputFile) == 0 {
		log.Fatal("input file is mandatory")
	}

	areas, err := parseFile(inputFile, xxxSkipSlice)
	if err != nil {
		log.Fatal(err)
	}
	if err = writeFile(inputFile, areas); err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := cli.App(AppName, AppDescr)
	app.Version("V version", fmt.Sprintf("%s v%s", AppName, AppVer))

	var (
		input = app.StringOpt("I input", ".", "Input directory")
		xxxTags = app.StringOpt("S skip", "", "skip tags to inject on XXX fields")
	)

	app.Action = func() {
		content, err := ioutil.ReadDir(*input)
		if err != nil { log.Fatalf("Read file error %v", err) }

		for _, f := range content {
			fpath := path.Join(*input, f.Name())
			if path.Ext(fpath) == ".go" {
				processFile(fpath, *xxxTags)
			}
		}
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("%s error: %v", AppName, err)
	}
}
