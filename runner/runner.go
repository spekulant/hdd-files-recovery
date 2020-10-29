package runner

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

var filesScanned int64 = 0

// Run triggers the recursive cloning process given a root path
func Run(root string) error {
	v1, err := readConfig()
	must(err)

	var context = runconf{
		FilterOut: v1.GetStringSlice("filterOut"),
		LookFor:   v1.GetStringSlice("lookFor"),
		rootPath:  root,
	}

	iterate(context)
	return nil
}

func iterate(context runconf) {
	filepath.Walk(context.rootPath, func(path string, info os.FileInfo, err error) error {
		dontPanicOn(err)
		must(copyOrPass(path, context, info))
		return nil
	})
}

func copyOrPass(path string, context runconf, info os.FileInfo) error {
	monitorProgress()

	if fileIsToBeFilteredOut(path, context) {
		return nil
	}

	if !isInteresting(filepath.Ext(path), context) {
		return nil
	}

	log.Printf("\nFound an interesting file, processing: %s", path)

	outputFileName := fmt.Sprintf("clone/%v-%s", info.Size(), info.Name())
	if thisFileIsCopiedCorrect(outputFileName, info.Size()) {
		log.Printf("file already copied, skipping: %s", info.Name())
		return nil
	}

	inputFile := fetchFile(path, 0)
	if inputFile == nil {
		log.Printf("Too many attempts at reading %s, skipping", path)
		return nil
	}
	must(saveFile(outputFileName, inputFile))

	log.Printf("file successfully copied to: %s", outputFileName)
	return nil
}

func monitorProgress() {
	filesScanned++
	if filesScanned%1000 == 0 {
		log.Printf("%v files scanned", filesScanned)
	}
}

func must(err error) {
	if err != nil {
		log.Fatal().Msg(err.Error())
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
