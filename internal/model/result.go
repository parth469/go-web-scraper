package model

type Event struct {
	Date      string `json:"date"`
	Time      string `json:"time"`
	Title     string `json:"title"`
	EventType string `json:"type"`
	Location  string `json:"location"`
}
