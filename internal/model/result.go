package model

import "time"

type Event struct {
	StartTime time.Time `json:"start"`
	EndTime   time.Time `json:"end"`
	Title     string `json:"title"`
	EventType string `json:"type"`
	Location  string `json:"location"`
}
