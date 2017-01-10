package main

import (
	"flag"
	"fmt"
	"os"

	app "github.com/kirsle/follow-sync/src"
)

func main() {
	wait := flag.Int("wait", 60, "How many seconds to pause between unfollows.")
	version := flag.Bool("version", false, "Show the application version number.")
	flag.Parse()

	if *version {
		fmt.Println(app.Version)
		os.Exit(0)
	}

	app := app.New()
	app.Wait = *wait
	app.Run()
}
