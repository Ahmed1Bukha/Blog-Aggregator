package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ahmed1bukha/Blog-Aggregator/internal/database"
)

func scrapeFeeds(s *state) error{
	nextFeed,err := s.db.GetNextFeedToFetch(context.Background())
	if err !=nil{
		return err
	}
	paramsMarked := database.MarkFeedFetchedParams{
		ID:         nextFeed.ID,
		UpdatedAt:  time.Now(),
		LastFetched: sql.NullTime{Time: time.Now(), Valid: true},
	}
	err = s.db.MarkFeedFetched(context.Background(), paramsMarked)
	if err != nil {
		return err
	}
	fmt.Println(nextFeed)
	return  nil
}