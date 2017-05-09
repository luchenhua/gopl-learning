package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", " "
	s += strings.Join(os.Args[1:], sep)
	fmt.Println(s)
}
