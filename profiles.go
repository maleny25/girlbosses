package main

type Profile struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	University  string `json:"university"`
	Major       string `json:"major"`
	Minor       string `json:"minor"`
	GradYear    int    `json:"grad_year"`
	Description string `json:"description"`
}
