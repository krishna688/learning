package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()

	urls := []string{
		"https://www.apple.com/",
		"https://www.microsoft.com/",
		"https://www.ibm.com/",
		"https://www.dell.com/",
	}

	sites, err := sitesInfo(urls)
	if err != nil {
		return
	}

	for url, info := range sites {
		fmt.Printf("%s: %+v\n", url, info)
	}

	fmt.Println(time.Since(start))

}
