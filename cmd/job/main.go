package main

import (
	"log"
	"t-bonatti/gopj/internal/rfb/download"
	"t-bonatti/gopj/internal/rfb/parser"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {

	s := gocron.NewScheduler(time.UTC)

	job, _ := s.Every(1).Seconds().Do(func() {
		pageParser := parser.NewRfbPageParser()
		rfbFiles, err := pageParser.GetRfbFiles()
		if err != nil {
			log.Fatal(err)
		}

		err = download.DownloadFiles(rfbFiles.FilesUrl)
		if err != nil {
			log.Fatal(err)
		}

	})
	job.SingletonMode()

	s.StartBlocking()
	defer s.Stop()
}
