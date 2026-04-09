package models

import (
	"time"

	"github.com/deepakraj-p28/RSS_Aggregator/internal/database"
	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID `json:"post_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Language    *string   `json:"language"`
	Publishedat time.Time `json:"published_at"`
	Link        string    `json:"post_url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func DatabasePostToPost(dbPost database.Post) Post {
	var description *string
	var language *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}
	if dbPost.Language.Valid {
		language = &dbPost.Language.String
	}

	return Post{
		ID:          dbPost.ID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: description,
		Language:    language,
		Publishedat: dbPost.Publishedat,
		Link:        dbPost.Link,
		FeedID:      dbPost.FeedID,
	}
}

func DatabasePostsToPosts(dbPosts []database.Post) []Post {
	posts := []Post{}
	for _, post := range dbPosts {
		posts = append(posts, DatabasePostToPost(post))
	}
	return posts
}
