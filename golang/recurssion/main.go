package main

import (
	"fmt"
	"reflect"
)

func main() {
	input := "krishna prasad"

	fmt.Println(reflect.TypeOf(input[1:]))
}
