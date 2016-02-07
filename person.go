package trakttv

// Person holds the person informations
type Person struct {
	Name   string `json:"name"`
	IDs    IDs    `json:"ids"`
	Images Images `json:"images"`
}
