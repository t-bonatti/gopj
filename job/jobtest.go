package job

import (
	"log"
	"t-bonatti/gopj/download"
	"t-bonatti/gopj/rfb"
	"time"

	"github.com/go-co-op/gocron"
)

func New() *gocron.Scheduler {
	s := gocron.NewScheduler(time.UTC)

	job, _ := s.Every(60).Seconds().Do(func() {
		pageParser := rfb.NewPageParser()
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

	return s
}
