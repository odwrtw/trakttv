package trakttv

import "time"

// Show holds the show informations
type Show struct {
	Title      string    `json:"title"`
	Year       int       `json:"year"`
	IDs        IDs       `json:"ids"`
	Rating     float64   `json:"rating"`
	Votes      int       `json:"votes"`
	FirstAired time.Time `json:"first_aired"`
	Overview   string    `json:"overview"`
	Status     string    `json:"status"`
	Images     Images    `json:"images"`
}

// Season holds the season informations
type Season struct {
	Number int `json:"number"`
	IDs    IDs `json:"ids"`
}

// Episode holds the episode informations
type Episode struct {
	Season int    `json:"season"`
	Number int    `json:"number"`
	Title  string `json:"title"`
	IDs    IDs    `json:"ids"`
	Images Images `json:"images"`
}
