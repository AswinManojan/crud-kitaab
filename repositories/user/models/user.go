package models

import "time"

type User struct {
	UserID    uint      `json:"userID" bun:"user_id,pk,autoincrement"`
	Name      string    `json:"name"`
	Height    string    `json:"height"`
	BirthDate time.Time `json:"birthDate"`
}
