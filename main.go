package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var filesScanned int64 = 0

func main() {
	logger := log.New(os.Stderr, "", 0)
	rootPath := os.Args[1]

	must(run(rootPath, logger))
}

func run(root string, logger *log.Logger) error {
	v1, err := readConfig()
	must(err)

	var context = runconf{
		FilterOut: v1.GetStringSlice("filterOut"),
		LookFor:   v1.GetStringSlice("lookFor"),
		rootPath:  root,
	}

	iterate(context, logger)
	return nil
}

func iterate(context runconf, logger *log.Logger) {
	filepath.Walk(context.rootPath, func(path string, info os.FileInfo, err error) error {
		dontPanicOn(err)
		must(copyOrPass(path, context, info, logger))
		return nil
	})
}

func copyOrPass(path string, context runconf, info os.FileInfo, logger *log.Logger) error {
	monitorProgress(logger)

	if fileIsToBeFilteredOut(path, context) {
		return nil
	}

	if !isInteresting(filepath.Ext(path), context) {
		return nil
	}

	logger.Printf("\nFound an interesting file, processing: %s", path)

	outputFileName := fmt.Sprintf("clone/%v-%s", info.Size(), info.Name())
	if thisFileIsCopiedCorrect(outputFileName, info.Size(), logger) {
		logger.Printf("file already copied, skipping: %s", info.Name())
		return nil
	}

	inputFile := fetchFile(path, 0, logger)
	if inputFile == nil {
		logger.Printf("Too many attempts at reading %s, skipping", path)
		return nil
	}
	must(saveFile(outputFileName, inputFile, logger))

	logger.Printf("file successfully copied to: %s", outputFileName)
	return nil
}

func monitorProgress(logger *log.Logger) {
	filesScanned++
	if filesScanned%1000 == 0 {
		logger.Printf("%v files scanned", filesScanned)
	}
}

func must(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func dontPanicOn(err error) {
	if err != nil {
		log.Print(err.Error())
	}
}

func readConfig() (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType("json")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	return v, err
}

type runconf struct {
	LookFor   []string
	FilterOut []string
	rootPath  string
}
