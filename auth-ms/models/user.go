package models

import "time"

type User struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	PeopleID     int       `json:"people_id"`
	Username     string    `json:"username" gorm:"size:100;unique"`
	Status       string    `json:"status" gorm:"type:status_enum"`
	LastActive   time.Time `json:"last_active"`
	LastActionBy int       `json:"last_action_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
