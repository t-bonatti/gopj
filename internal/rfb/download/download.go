package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"
	"time"

	"github.com/apex/log"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

func DownloadFiles(urls []string) (err error) {

	var wg sync.WaitGroup
	p := mpb.New(mpb.WithWaitGroup(&wg))
	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			err = downloadFile("/Users/bonatti/Downloads/tmp/", url, p)
			if err != nil {
				log.Info(fmt.Sprint("Error on download file: ", url, "|", err))
			}

		}(url)
	}

	p.Wait()
	return
}

func downloadFile(folder string, url string, p *mpb.Progress) error {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	filename := path.Base(r.URL.Path)
	filepath := folder + filename
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}

	client := http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(r)
	if err != nil {
		out.Close()
		return err
	}
	defer resp.Body.Close()

	fsize, _ := strconv.Atoi(resp.Header.Get("Content-Length"))

	bar := buildBar(p, filename, int64(fsize))
	proxyReader := bar.ProxyReader(resp.Body)
	defer proxyReader.Close()

	if _, err = io.Copy(out, proxyReader); err != nil {
		out.Close()
		return err
	}

	out.Close()

	if err = os.Rename(filepath+".tmp", filepath); err != nil {
		return err
	}
	return nil
}

func buildBar(p *mpb.Progress, name string, total int64) *mpb.Bar {
	return p.AddBar(total,
		mpb.PrependDecorators(
			decor.Name(name, decor.WC{W: len(name) + 1, C: decor.DidentRight}),
			decor.CountersKibiByte("% .2f / % .2f"),
		),
		mpb.AppendDecorators(
			decor.Percentage(decor.WC{W: 5}),
		),
	)
}
