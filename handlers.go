package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/seiobata/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("no username provided\n")
	}
	name := cmd.args[0]
	if _, err := s.db.GetUser(context.Background(), name); err != nil {
		return fmt.Errorf("Name %v does not exist in database\n", name)
	}
	if err := s.cfg.SetUser(name); err != nil {
		return err
	}
	fmt.Printf("User set to %v\n", name)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("no username provided\n")
	}
	name := cmd.args[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("problem calling CreateUser: %v", err)
	}
	if err = s.cfg.SetUser(user.Name); err != nil {
		return fmt.Errorf("problem calling SetUser: %v", err)
	}
	fmt.Printf("User %v was created on %v\n", user.Name, user.CreatedAt)
	return nil
}
