package main

import (
	"fmt"
	"log"

	"github.com/seiobata/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}
	fmt.Printf("Read config:\n%+v\n", cfg)

	if err = cfg.SetUser("name"); err != nil {
		log.Fatalf("could not set user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("could not read config: %v", err)
	}
	fmt.Printf("Read config again:\n%+v\n", cfg)
}
