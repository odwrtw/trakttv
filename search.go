package trakttv

import "fmt"

// SearchMovieByID searches a movie by its trakttv ID / slug or by its IMDB ID
func (t *TraktTv) SearchMovieByID(id string, qo QueryOption) (*Movie, error) {
	url := fmt.Sprintf("%s/%s/%s", t.Endpoint, "movies", id)

	var movie *Movie
	if err := t.request(url, NewQuery(qo), &movie); err != nil {
		return nil, err
	}

	return movie, nil
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
	url := fmt.Sprintf("%s/search", t.Endpoint)

	q := NewQuery(qo)
	q.sq = &sq

	var res []*SearchResult
	if err := t.request(url, q, &res); err != nil {
		return nil, err
	}

	return res, nil
}
