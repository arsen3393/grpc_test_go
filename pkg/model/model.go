package model

import (
	"time"
)

type Anomalies struct {
	ID        int       `gorm:"primary_key;"`
	SessionID string    `gorm:"not null"`
	Frequency float64   `gorm:"not null"`
	Timestamp time.Time `gorm:"not null"`
}
