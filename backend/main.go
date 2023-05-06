package main

import (
	"flag"
	"fmt"
	"os"
	cmdPkg "the-video-project/backend/internal/command"
	"the-video-project/backend/internal/httpapi"
)

var cmd string
var subcmd string

func init() {
	flag.StringVar(&cmd, "command", "server", "想執行的指令 (server, migrate...)")
	flag.StringVar(&subcmd, "subcommand", "", "想執行的次級指令 (up, down...)")
}

func main() {
	flag.Parse()
	var cmdResolved cmdPkg.Command
	switch cmd {
	case "migrate":
		cmdResolved = cmdPkg.MigrateCommand{Subcommand: subcmd}
	case "seed":
		cmdResolved = cmdPkg.SeedCommand{}
	}

	var err error
	if cmdResolved != nil {
		err = cmdResolved.Run()
	} else {
		err = startHttpApi(5000)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

// This function cannot be a command, otherwise cause import cycle
func startHttpApi(port uint) error {
	router := httpapi.SetupRouter()
	return router.Run(fmt.Sprintf(":%d", port))
}
