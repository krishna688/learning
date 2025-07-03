package main

import (
	"fmt"
	"image"
	"log"
)

func main() {
	rect := image.Rect(0, 0, 50, 50)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println(scaleDir(".", "edited/.", rect))
}
