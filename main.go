package main

import (
	"bootfeeds/internal/config"
	"bootfeeds/internal/handlers"
	"log"
	"os"
)

func main() {

	cfg := config.Read()
	s := config.State{Cfg: cfg}

	var cmds config.Commands
	cmds = config.Commands{
		Commands: map[string]func(*config.State, config.Command) error{},
	}

	cmds.Register("login", handlers.LoginHandler)

	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments.")
		return
	}

	name := os.Args[1]
	args := os.Args[2:]

	cmds.Run(&s, config.Command{Name: name, Cmd: args})

}
