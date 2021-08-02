package main

type Profile struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Race        string `json:"race"`
	Age         string    `json:"age"`
	University  string `json:"university"`
	Major       string `json:"major"`
	Minor       string `json:"minor"`
	Gradyear    string    `json:"grad_year"`
	Description string `json:"description"`
}
