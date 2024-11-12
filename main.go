package main

import (
	"os"
	"github.com/mdxabu/ipscout/web"
	"github.com/mdxabu/ipscout/cmd"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "geo" {
		go web.StartWebServer()
	}
	cmd.Execute()
}