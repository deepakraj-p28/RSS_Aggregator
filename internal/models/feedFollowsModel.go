package models

import (
	"time"

	"github.com/deepakraj-p28/RSS_Aggregator/internal/database"
	"github.com/google/uuid"
)

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func DatabaseFeedFollowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}

}

func DatabaseFeedFollowsToFeedFollows(dbFeedFollow []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}

	for _, feedFollow := range dbFeedFollow {
		feedFollows = append(feedFollows, DatabaseFeedFollowToFeedFollow(feedFollow))
	}

	return feedFollows
}
