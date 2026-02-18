package main

import (
	"context"
	"fmt"
)

func handlerGetFeed(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("problem fetching feed: %v", err)
	}
	fmt.Printf("%+v\n", feed)
	return nil
}
