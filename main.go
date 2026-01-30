package main

import (
	"log"
	"os"

	"github.com/seiobata/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}
	s := &state{
		cfg: &cfg,
	}
	cmds := commands{
		names: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

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
