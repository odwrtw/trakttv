package trakttv

// Images represents the images linked to a video
type Images struct {
	Fanart   []string `json:"fanart"`
	Poster   []string `json:"poster"`
	Logo     []string `json:"logo"`
	Clearart []string `json:"clearart"`
	Banner   []string `json:"banner"`
	Thumb    []string `json:"thumb"`
}
