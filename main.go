package main

import (
	"os"
)

func main()  {
	app := parseArgs()
	app.Version = "v0.0.1"
	app.Run(os.Args)
}