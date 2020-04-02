package trakttv

import (
	"errors"
	"fmt"
)

// SearchMovieByID searches a movie by its trakttv ID / slug or by its IMDB ID
func (t *TraktTv) SearchMovieByID(id string, qo QueryOption) (*Movie, error) {
	url := fmt.Sprintf("%s/%s/%s", t.Endpoint, "movies", id)

	var movie Movie
	if err := t.request(url, NewQuery(qo), &movie); err != nil {
		return nil, err
	}

	return &movie, nil
}

// SearchShowByID searches a show by its trakttv ID / slug or by its IMDB ID
func (t *TraktTv) SearchShowByID(id string, qo QueryOption) (*Show, error) {
	url := fmt.Sprintf("%s/%s/%s", t.Endpoint, "shows", id)

	var show Show
	if err := t.request(url, NewQuery(qo), &show); err != nil {
		return nil, err
	}

	return &show, nil
}

// SearchEpisode searches an episode by its trakttv ID / slug / IMDB ID,
// season and episode number
func (t *TraktTv) SearchEpisode(id string, season, episode int, qo QueryOption) (*Episode, error) {
	url := fmt.Sprintf("%s/shows/%s/seasons/%d/episodes/%d", t.Endpoint, id, season, episode)

	var e Episode
	if err := t.request(url, NewQuery(qo), &e); err != nil {
		return nil, err
	}

	return &e, nil
}

// SearchResult represents a generic search result
type SearchResult struct {
	Type   Type    `json:"type"`
	Score  float64 `json:"score"`
	Movie  *Movie  `json:"movie"`
	Show   *Show   `json:"show"`
	Person *Person `json:"person"`

	// TODO: deal with this shit later
	// When there is an episode there is also a show, but the show as the year
	// as a string, not as a int...
	// Episode *Episode `json:"episode"`
}

// Search searches for anything
func (t *TraktTv) Search(sq SearchQuery, qo QueryOption) ([]*SearchResult, error) {
	if sq.Type == "" || sq.Query == "" {
		return nil, errors.New("missing argument")
	}
	url := fmt.Sprintf("%s/search/%s", t.Endpoint, sq.Type)

	q := NewQuery(qo)
	q.sq = &sq

	var res []*SearchResult
	if err := t.request(url, q, &res); err != nil {
		return nil, err
	}

	return res, nil
}
