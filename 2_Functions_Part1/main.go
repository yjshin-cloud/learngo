package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("nico", "lynn", "dal", "marl", "flynn")

	totalLenght, upperName := lenAndUpper("nico")
	fmt.Println(totalLenght, upperName)
}
