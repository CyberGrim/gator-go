package main

import (
	"context"
	"database/sql"
	"log"

	// "strings"
	"time"

	"github.com/cybergrim/gator/internal/database"
	"github.com/google/uuid"
)

func scrapeFeeds(s *state) {
	nextFeed, nextFeedError := s.db.GetNextFeedToFetch(context.Background())
	if nextFeedError != nil {
		log.Println(nextFeedError)
		return
	}

	_, markError := s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if markError != nil {
		log.Println(markError)
		return
	}

	feedData, feedDataError := fetchFeed(context.Background(), nextFeed.Url)
	log.Printf("Feed %s collected, %v posts found", nextFeed.Name, len(feedData.Channel.Item))
	if feedDataError != nil {
		log.Println(feedDataError)
		return
	}

	for i := range feedData.Channel.Item {
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, feedData.Channel.Item[i].PubDate); err == nil {
			publishedAt = sql.NullTime{Time: t, Valid: true}
		}
		url := feedData.Channel.Item[i].Link
		if url == "" {
			url = feedData.Channel.Item[i].GUID
		}
		_, createError := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       feedData.Channel.Item[i].Title,
			Url:         url,
			Description: sql.NullString{String: feedData.Channel.Item[i].Description, Valid: true},
			PublishedAt: publishedAt,
			FeedID:      nextFeed.ID,
		})
		if createError != nil {
			log.Println("Create error:", createError)
			continue
		}
		// if createError != nil {
		// 	if strings.Contains(createError.Error(), "duplicate key value violates unique constraint") {
		// 		continue
		// 	} else {
		// 		log.Println(createError)
		// 		continue
		// 	}
		// } else {
		// 	log.Println("Post saved:", feedData.Channel.Item[i].Title)
		// }
	}
}
