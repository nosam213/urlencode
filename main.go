package main

import (
	"net/url"
	"os"
)

func UrlEncode(x string) string {
	var z string = url.QueryEscape(x)
	return z
}

func main() {
  if len(os.Args) > 1 {
		println(UrlEncode(os.Args[1]))
	} else {
    println("No input.")
		os.Exit(0)
	}
}