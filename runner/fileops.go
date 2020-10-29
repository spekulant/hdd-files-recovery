package runner

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

func thisFileIsCopiedCorrect(fname string, size int64) bool {
	log.Printf("Checking if file has already been copied")
	fstat, err := os.Stat(fname)
	if err != nil {
		return false
	}
	if fstat.Size() != size {
		log.Printf("local and remote file sizes dont match, removing and trying again: %s", fstat.Name())
		must(os.Remove(fname))
		return false
	}
	return true
}

func fetchFile(fpath string, attempt int) []byte {
	log.Printf("Attempting to fetch file: %v", attempt)
	// TODO: figure out a way to retry fetching the file when the HDD freezes and the below hangs indefinitely
	inputFile, err := ioutil.ReadFile(fpath)

	if err != nil {
		time.Sleep(time.Duration(500 * attempt))
		if attempt > 15 {
			log.Printf("Unable to load %s", fpath)
			return nil
		}
		return fetchFile(fpath, attempt+1)
	}
	return inputFile
}

func saveFile(fpath string, data []byte) error {
	log.Printf("Saving file locally")
	err := ioutil.WriteFile(fpath, data, 0644)

	if err != nil {
		log.Printf("Unable to save %s", fpath)
	}
	return nil
}
