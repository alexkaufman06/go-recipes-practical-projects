package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

func downloadSize(url string) (int, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("bad status: %d %s", resp.StatusCode, resp.Status)
	}

	return strconv.Atoi(resp.Header.Get("Content-Length"))
}

func downloadsSize(urls []string) (int, error) {
	var dlSize int64
	var g errgroup.Group

	for _, url := range urls {
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			log.Print(url)
			size, err := downloadSize(url)
			if err != nil {
				return err
			}

			atomic.AddInt64(&dlSize, int64(size))
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return 0, err
	}

	return int(dlSize), nil
}

func gen2020URLs() []string {
	var urls []string
	//urlTemplate := "https://s3.amazonaws.com/nyc-tlc/trip+data/%s_tripdata_2020-%02d.csv"
	//for _, vendor := range []string{"yellow", "green"} {
	//	for month := 1; month <= 12; month++ {
	//		url := fmt.Sprintf(urlTemplate, vendor, month)
	//		urls = append(urls, url)
	//	}
	//}
	urls = append(urls, "https://stats.govt.nz/assets/Uploads/Annual-enterprise-survey/Annual-enterprise-survey-2021-financial-year-provisional/Download-data/annual-enterprise-survey-2021-financial-year-provisional-size-bands-csv.csv")
	urls = append(urls, "https://stats.govt.nz/assets/Uploads/Business-employment-data/Business-employment-data-September-2022-quarter/Download-data/Business-employment-data-september-2022-quarter-csv.zip")
	urls = append(urls, "https://stats.govt.nz/assets/Uploads/Annual-enterprise-survey/Annual-enterprise-survey-2021-financial-year-provisional/Download-data/annual-enterprise-survey-2021-financial-year-provisional-csv.csv")
	return urls
}

func main() {
	urls := gen2020URLs()
	size, err := downloadsSize(urls)
	if err != nil {
		// We will get a 403 error trying to hit s3 due to access issues
		log.Fatal(err)
	}

	sizeMB := float64(size) / (1 << 20)
	fmt.Printf("size = %.2fMB\n", sizeMB)
}
