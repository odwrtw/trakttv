package trakttv

// baseImage represents the a generic image
type baseImage struct {
	Full   string `json:"full"`
	Medium string `json:"medium"`
	Thumb  string `json:"thumb"`
}

// baseImageFullOnly represents the a generic image with only the full option
type baseImageFullOnly struct {
	Full string `json:"full"`
}

// Images represents the images linked to a video
type Images struct {
	Poster     baseImage         `json:"poster"`
	Fanart     baseImage         `json:"fanart"`
	Screenshot baseImage         `json:"screenshot"`
	Headshot   baseImage         `json:"headshot"`
	Banner     baseImageFullOnly `json:"banner"`
	Logo       baseImageFullOnly `json:"logo"`
	Clearart   baseImageFullOnly `json:"clearart"`
	Thumb      baseImageFullOnly `json:"thumb"`
	Avatar     baseImageFullOnly `json:"avatar"`
}
