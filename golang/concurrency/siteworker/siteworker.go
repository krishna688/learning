package main

import (
	"io"
	"net/http"
	"time"
)

type Info struct {
	Url        string
	StatusCode int
	Delay      time.Duration
	err        error
}

type workerInfo struct {
	Info
}

func sitesInfo(urls []string) (map[string]Info, error) {

	sites := map[string]Info{}

	c := make(chan Info)
	for _, url := range urls {
		go func(u string) {
			info, err := siteInfo(u)
			info.Url = u
			info.err = err

			c <- info

		}(url)

	}

	for range urls {

		info := <-c
		sites[info.Url] = info
	}

	return sites, nil
}

func siteInfo(url string) (info Info, err error) {

	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if _, err = io.Copy(io.Discard, res.Body); err != nil {
		return
	}

	info.StatusCode = res.StatusCode
	info.Delay = time.Since(start)

	return
}
