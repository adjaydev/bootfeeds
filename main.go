package main

import (
	"bootfeeds/internal/config"
	"bootfeeds/internal/database"
	"bootfeeds/internal/handlers"
	"database/sql"

	// "database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	cfg := config.Read()

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal("Error connecting to the database.")
		return
	}

	dbQueries := database.New(db)

	s := config.State{Cfg: cfg, DB: dbQueries}

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
