package main

import (
	"flag"
	"fmt"

	"github.com/Naithar01/dc_cli_crawler/crawler"
)

func main() {
	var count int
	var reverse bool

	// -n 5 -r 이런 식으로
	// 	go build cli-app.go
	// ./cli-app -n 5 -r
	flag.IntVar(&count, "n", 10, "number of items to show")
	flag.BoolVar(&reverse, "r", false, "show results in reverse order")

	flag.Parse()

	data := crawler.Page()

	if reverse {
		for i := len(data) - 1; i >= 0 && count > 0; i-- {
			fmt.Println(data[i])
			count--
		}
	} else {
		for i := 0; i < len(data) && count > 0; i++ {
			fmt.Println(data[i])
			count--
		}
	}
}
