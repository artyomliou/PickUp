package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	cmdPkg "pick-up/backend/internal/command"
	"pick-up/backend/internal/httpapi"
	"syscall"
	"time"
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
		err = startHttpApi()
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}

// This function cannot be a command, otherwise cause import cycle
func startHttpApi() error {
	router := httpapi.SetupRouter()
	srv := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Shutdown server if it received these signals...
	// https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// Shutdown in 2 seconds or less
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// catching ctx.Done(). timeout of 2 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 2 seconds.")
	}
	log.Println("Server exiting")

	return nil
}
