package main

type Event struct {
	Name        string   `json:"name"`
	Date        string   `json:"date"`
	Description string   `json:"description"`
	Link        string   `json:"link"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"` // used to filter
}
