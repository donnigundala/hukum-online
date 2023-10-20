package models

import (
	"time"
)

type Gathering struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	CreatorID   uint      `json:"creator_id"`
	TypeID      uint      `json:"type_id"`
	ScheduledAt time.Time `json:"scheduled_at"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
