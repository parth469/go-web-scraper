package model

import "time"

type Event struct {
	StartTime time.Time `json:"start"`
	EndTime   time.Time `json:"end"`
	Title     string    `json:"title"`
	EventType string    `json:"type"`
	Location  string    `json:"location"`
}

type Poster struct {
	StartTime time.Time `json:"start"`
	EndTime   time.Time `json:"end"`
	Title     string    `json:"title"`
	Presenter string    `json:"presenter"`
	PosterId  string    `json:"poster_id"`
	TopicArea string    `json:"topic_area"`
	Abstract  string    `json:"abstract"`
}
