package main

import (
	"bootfeeds/internal/config"
	"bootfeeds/internal/database"
	"bootfeeds/internal/handlers"
	"context"
	"database/sql"
	"fmt"

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
	cmds.Register("register", handlers.RegisterHandler)
	cmds.Register("reset", handlers.ResetHandler)
	cmds.Register("users", handlers.GetUsersHandler)
	cmds.Register("agg", handlers.FeedHandler)
	cmds.Register("addfeed", middlewareLoggedIn(handlers.FeedAddHandler))
	cmds.Register("feeds", handlers.FeedsHandler)
	cmds.Register("follow", middlewareLoggedIn(handlers.FeedFollowsHandler))
	cmds.Register("following", middlewareLoggedIn(handlers.FeedFollowingHandler))

	if len(os.Args) < 2 {
		log.Fatal("Not enough arguments.")
		return
	}

	name := os.Args[1]
	args := os.Args[2:]

	err = cmds.Run(&s, config.Command{Name: name, Cmd: args})
	if err != nil {
		log.Fatal(err)
	}

}

func middlewareLoggedIn(handler func(s *config.State, cmd config.Command, user database.User) error) func(*config.State, config.Command) error {
	return func(s *config.State, cmd config.Command) error {
		u, err := s.DB.GetUser(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("User not found.")
		}
		return handler(s, cmd, u)
	}
}
