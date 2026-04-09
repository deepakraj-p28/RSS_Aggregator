package models

import (
	"time"

	"github.com/deepakraj-p28/RSS_Aggregator/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apikey"`
}

func DatabaseUserToUser(DBUser database.User) User {
	return User{
		ID:        DBUser.ID,
		CreatedAt: DBUser.CreatedAt,
		UpdatedAt: DBUser.UpdatedAt,
		Name:      DBUser.Name,
		ApiKey:    DBUser.ApiKey,
	}
}

func DatabaseUsersToUsers(DBUsers []database.User) []User {
	users := []User{}
	for _, dbUser := range DBUsers {
		users = append(users, DatabaseUserToUser(dbUser))
	}

	return users
}
