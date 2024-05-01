package models

import "time"

type User struct {
	UserID    uint      `json:"userID" bun:"id,pk,autoincrement"`
	Name      string    `json:"name" bun:"name,notnull"`
	Height    string    `json:"height" bun:"height"`
	BirthDate time.Time `json:"birthDate" bun:"date"`
}
