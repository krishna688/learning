package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	responseSize("https://www.example.com")
	responseSize("https://www.geeksforgeeks.org/linked-list-data-structure/")
	responseSize("https://training.play-with-docker.com/")
	responseSize("https://leetcode.com/discuss/post/665604/important-and-useful-links-from-all-over-ocy8/")
	responseSize("https://www.udemy.com/")
}

func responseSize(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error : %s \n", err)
	}

	defer resp.Body.Close()

	by, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error %s \n", err)
	}

	log.Println(url[:20], ":", len(by))
}
