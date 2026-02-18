package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/seiobata/gator/internal/config"
	"github.com/seiobata/gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("could not open database connection: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	s := &state{
		cfg: &cfg,
		db:  dbQueries,
	}

	cmds := commands{
		names: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerResetUsers)
	cmds.register("users", handlerGetAllUsers)
	cmds.register("agg", handlerGetFeed)

	words := os.Args
	if len(words) < 2 {
		log.Fatal("not enough arguments provided")
	}

	cmd := command{
		name: words[1],
		args: words[2:],
	}

	if err = cmds.run(s, cmd); err != nil {
		log.Fatalf("problem executing command: %v", err)
	}
}
