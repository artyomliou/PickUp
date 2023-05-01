package main

import (
	"flag"
	"fmt"
	"os"
	cmdPkg "the-video-project/backend/internal/command"
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
	case "server":
		cmdResolved = cmdPkg.ServerCommand{Port: 5000}
	case "migrate":
		cmdResolved = cmdPkg.MigrateCommand{Subcommand: subcmd}
	case "seed":
		cmdResolved = cmdPkg.SeedCommand{}
	}
	if cmdResolved != nil {
		err := cmdResolved.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			os.Exit(0)
		}
	}
}
