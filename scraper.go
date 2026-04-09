package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/deepakraj-p28/RSS_Aggregator/internal/database"
	"github.com/google/uuid"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequests time.Duration) {
	log.Printf("Started Scraping on %d goroutines every %s duration \n", concurrency, timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Error while fetching feeds")
			continue
		}

		waitGrp := &sync.WaitGroup{}
		for _, feed := range feeds {
			waitGrp.Add(1)
			go scrapeFeed(waitGrp, db, feed)
		}
		waitGrp.Wait()
	}
}

func scrapeFeed(wg *sync.WaitGroup, db *database.Queries, dbFeed database.Feed) {
	defer wg.Done()

	feed, err := db.MarkFeedAsFetched(context.Background(), dbFeed.ID)
	if err != nil {
		log.Println("Error which marking the feed as done")
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = db.CreatePost(context.Background(), database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		Title:       rssFeed.Channel.Title,
		Description: getSqlNullString(rssFeed.Channel.Description),
		Language:    getSqlNullString(rssFeed.Channel.Language),
		Publishedat: convertStringToTime(rssFeed.Channel.PubDate),
		Link:        rssFeed.Channel.Link,
		FeedID:      feed.ID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
		} else {
			log.Println("Could not create a new post:", err)
		}
	}

	log.Printf("%v Feeds collected ", len(rssFeed.Channel.Item))
}

func getSqlNullString(item string) sql.NullString {
	itemNullString := sql.NullString{}

	if item != "" {
		itemNullString.String = item
		itemNullString.Valid = true
	}

	return itemNullString
}

func convertStringToTime(timestamp string) time.Time {
	timeVal, err := time.Parse(time.RFC1123, timestamp)
	if err != nil {
		log.Println("Error while parsing time, returning default")
		return time.Time{}
	}

	return timeVal
}
