package main

import (
	"github.com/pingc0y/URLFinder/cmd"
	"github.com/pingc0y/URLFinder/crawler"
	"io"
	"log"
)

func main() {
	log.SetOutput(io.Discard)
	cmd.Parse()
	crawler.Run()
}
