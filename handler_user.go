package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mpetkov228/gator/internal/database"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	username := cmd.Args[0]

	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User logged in successfully!")
	return nil
}

func handleRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]
	id := uuid.New()
	createdAt := time.Now()
	updatedAt := time.Now()
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("error creating user: %v", err)
	}

	s.cfg.SetUser(user.Name)

	fmt.Println("New user was created!")
	fmt.Printf("%v\n", user)
	return nil
}
