package main

import (
	"context"
	"fmt"
	"log"
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
	if feedDataError != nil {
		log.Println(feedDataError)
		return
	}

	for i := range feedData.Channel.Item {
		fmt.Println(feedData.Channel.Item[i].Title)
	}
}
