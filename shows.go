package trakttv

import (
	"fmt"
	"time"
)

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
	Number        int        `json:"number"`
	IDs           IDs        `json:"ids"`
	Rating        float64    `json:"rating"`
	Votes         int        `json:"votes"`
	EpisodeCount  int        `json:"episode_count"`
	AiredEpisodes int        `json:"aired_episodes"`
	Title         string     `json:"title"`
	Overview      string     `json:"overview"`
	FirstAired    time.Time  `json:"first_aired"`
	Network       string     `json:"network"`
	Episodes      []*Episode `json:"episodes"`
}

// Episode holds the episode informations
type Episode struct {
	Season     int       `json:"season"`
	Number     int       `json:"number"`
	Title      string    `json:"title"`
	IDs        IDs       `json:"ids"`
	Overview   string    `json:"overview"`
	Rating     float64   `json:"rating"`
	Votes      int       `json:"votes"`
	Runtime    int       `json:"runtime"`
	FirstAired time.Time `json:"first_aired"`
}

// TrendingShow represents a trending show with the number of trakt.tv user
// currently watching it
type TrendingShow struct {
	Watchers int  `json:"watchers"`
	Show     Show `json:"show"`
}

// AnticipatedShow represents an anticipated show with the number of trakt.tv
// user having the show in their wishlist
type AnticipatedShow struct {
	ListCount int  `json:"list_count"`
	Show      Show `json:"show"`
}

// PlayedShow represents the show and its played count
type PlayedShow struct {
	WatcherCount   int  `json:"watcher_count"`
	PlayCount      int  `json:"play_count"`
	CollectedCount int  `json:"collected_count"`
	Show           Show `json:"show"`
}

// PopularShows returns the popular shows
func (t *TraktTv) PopularShows(qo QueryOption) ([]*Show, error) {

	url := fmt.Sprintf("%s/%s", t.Endpoint, "shows/popular")

	var shows []*Show
	if err := t.request(url, NewQuery(qo), &shows); err != nil {
		return nil, err
	}

	return shows, nil
}

// TrendingShows returns the trending shows
func (t *TraktTv) TrendingShows(qo QueryOption) ([]*TrendingShow, error) {

	url := fmt.Sprintf("%s/%s", t.Endpoint, "shows/trending")

	var shows []*TrendingShow
	if err := t.request(url, NewQuery(qo), &shows); err != nil {
		return nil, err
	}

	return shows, nil
}

// AnticipatedShows returns the trending shows
func (t *TraktTv) AnticipatedShows(qo QueryOption) ([]*AnticipatedShow, error) {

	url := fmt.Sprintf("%s/%s", t.Endpoint, "shows/anticipated")

	var shows []*AnticipatedShow
	if err := t.request(url, NewQuery(qo), &shows); err != nil {
		return nil, err
	}

	return shows, nil
}

// PlayedShows returns the mosts played shows
func (t *TraktTv) PlayedShows(qo QueryOption) ([]*PlayedShow, error) {

	url := fmt.Sprintf("%s/%s", t.Endpoint, "shows/played")

	var shows []*PlayedShow
	if err := t.request(url, NewQuery(qo), &shows); err != nil {
		return nil, err
	}

	return shows, nil
}

// WatchedShows returns the mosts played shows by unique users
func (t *TraktTv) WatchedShows(qo QueryOption) ([]*PlayedShow, error) {

	url := fmt.Sprintf("%s/%s", t.Endpoint, "shows/watched")

	var shows []*PlayedShow
	if err := t.request(url, NewQuery(qo), &shows); err != nil {
		return nil, err
	}

	return shows, nil
}

// CollectedShows returns the mosts collected shows
func (t *TraktTv) CollectedShows(qo QueryOption) ([]*PlayedShow, error) {

	url := fmt.Sprintf("%s/%s", t.Endpoint, "shows/collected")

	var shows []*PlayedShow
	if err := t.request(url, NewQuery(qo), &shows); err != nil {
		return nil, err
	}

	return shows, nil
}

// GetShowSeasons returns the seasons of a show
func (t *TraktTv) GetShowSeasons(id string, qo QueryOption) ([]*Season, error) {
	url := fmt.Sprintf("%s/shows/%s/seasons", t.Endpoint, id)

	var seasons []*Season
	if err := t.request(url, NewQuery(qo), &seasons); err != nil {
		return nil, err
	}

	return seasons, nil
}
