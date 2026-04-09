package models

import (
	"time"

	"github.com/deepakraj-p28/RSS_Aggregator/internal/database"
	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func DatabaseFeedToFeed(DBFeed database.Feed) Feed {
	return Feed{
		ID:        DBFeed.ID,
		CreatedAt: DBFeed.CreatedAt,
		UpdatedAt: DBFeed.UpdatedAt,
		Name:      DBFeed.Name,
		Url:       DBFeed.Url,
		UserID:    DBFeed.UserID,
	}
}

func DatabaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, DatabaseFeedToFeed(dbFeed))
	}

	return feeds
}
