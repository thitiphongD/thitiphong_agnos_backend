package models

import "time"

type Logger struct {
	ID         uint `gorm:"primarykey"`
	Request    string
	Response   string
	ClientIP   string
	Method     string
	Path       string
	Proto      string
	StatusCode int
	Latency    string
	UserAgent  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
