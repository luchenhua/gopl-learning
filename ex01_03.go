package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//TODO: write a performance test for normal "+" vs "string.Join" after finish ch1.6 & ch11.4
	s, sep := "", " "
	s += strings.Join(os.Args[1:], sep)
	fmt.Println(s)
}
