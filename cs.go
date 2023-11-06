package main

import (
	"fmt"
	"unittesteExample/split"
)

func main() {
	s := split.Split("a:bdsd:c:d:e", ":")
	fmt.Println("data:", s)
}
