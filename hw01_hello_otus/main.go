package main

import (
	"fmt"

	"github.com/agrison/go-commons-lang/stringUtils"
)

func main() {
	str := "Hello, OTUS!"
	reversed := stringUtils.Reverse(str)
	fmt.Println(reversed)
}
