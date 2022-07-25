package entities

import "time"

type User struct {
	ID        int        `gorm:"primary_key" json:"id"`
	Name      string     `db:"name" json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"_"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}
