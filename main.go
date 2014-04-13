package main

import (
	"./rakefile"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	rakefile.IsFile("rakefile/rakefile.go")
	cli.NewApp().Run(os.Args)
}
