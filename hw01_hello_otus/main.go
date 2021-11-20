package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	reverseMessage := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(reverseMessage)
}
