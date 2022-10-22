package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	l := 0
	if len(os.Args) > 1 {
		for _, v := range strings.Split(strings.TrimSpace(os.Args[1]), " ") {
			if v != "" {
				l++
			}
		}
	}
	fmt.Printf("%v", l)

}
