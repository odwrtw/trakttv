package trakttv

import "net/http"

// Constants to remove later
const (
	ProductionEndpoint = "https://api-v2launch.trakt.tv"
	StagingEndpoint    = "http://api.staging.trakt.tv"
)

// TraktTv client
type TraktTv struct {
	HTTPClient *http.Client
	Endpoint   string
	Version    int
	Key        string
}

// Type represents the types of data the API could return
type Type string

// Available types
const (
	TypeMovie   Type = "movie"
	TypeShow         = "show"
	TypeEpisode      = "episode"
	TypePerson       = "person"
	TypeList         = "list"
)

// New returns a new tvrage client
func New(key string) *TraktTv {
	return &TraktTv{
		HTTPClient: http.DefaultClient,
		Version:    2,
		Endpoint:   ProductionEndpoint,
		Key:        key,
	}
}

// IDs holds the media object ID
type IDs struct {
	Trakt  int    `json:"trakt"`
	Slug   string `json:"slug"`
	ImDB   string `json:"imdb"`
	TmDB   int    `json:"tmdb"`
	TvRage int    `json:"tvrage"`
	TvDB   int    `json:"tvdb"`
}
