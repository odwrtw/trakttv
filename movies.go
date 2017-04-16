package trakttv

import "fmt"

// Movie holds the movie informations
type Movie struct {
	Title         string   `json:"title"`
	Year          int      `json:"year"`
	Images        Images   `json:"images"`
	IDs           IDs      `json:"ids"`
	Tagline       string   `json:"tagline"`
	Overview      string   `json:"overview"`
	Votes         int      `json:"votes"`
	Released      string   `json:"released"`
	Runtime       int      `json:"runtime"`
	Language      string   `json:"language"`
	Rating        float64  `json:"rating"`
	Genres        []string `json:"genres"`
	Certification string   `json:"certification"`
}

// TrendingMovie represents a trending movie with the number of trakt.tv user
// currently watching it
type TrendingMovie struct {
	Watchers int   `json:"watchers"`
	Movie    Movie `json:"movie"`
}

// AnticipatedMovie represents an anticipated movie with the number of trakt.tv
// user having the movie in their wishlist
type AnticipatedMovie struct {
	ListCount int   `json:"list_count"`
	Movie     Movie `json:"movie"`
}

// BoxOfficeMovie represents the movie and its revenue
type BoxOfficeMovie struct {
	Revenue int   `json:"revenue"`
	Movie   Movie `json:"movie"`
}

// PlayedMovie represents the movie and its played count
type PlayedMovie struct {
	WatcherCount   int   `json:"watcher_count"`
	PlayCount      int   `json:"play_count"`
	CollectedCount int   `json:"collected_count"`
	Movie          Movie `json:"movie"`
}

// PopularMovies returns the popular movies
func (t *TraktTv) PopularMovies(qo QueryOption) ([]*Movie, error) {

	url := fmt.Sprintf("%s/%s", t.Endpoint, "movies/popular")

	var movies []*Movie
	if err := t.request(url, NewQuery(qo), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

// TrendingMovies returns the trending movies
func (t *TraktTv) TrendingMovies(qo QueryOption) ([]*TrendingMovie, error) {

	url := fmt.Sprintf("%s/%s", t.Endpoint, "movies/trending")

	var movies []*TrendingMovie
	if err := t.request(url, NewQuery(qo), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

// AnticipatedMovies returns the trending movies
func (t *TraktTv) AnticipatedMovies(qo QueryOption) ([]*AnticipatedMovie, error) {

	url := fmt.Sprintf("%s/%s", t.Endpoint, "movies/anticipated")

	var movies []*AnticipatedMovie
	if err := t.request(url, NewQuery(qo), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

// BoxOfficeMovies returns the movies in the box office
func (t *TraktTv) BoxOfficeMovies(qo QueryOption) ([]*BoxOfficeMovie, error) {

	url := fmt.Sprintf("%s/%s", t.Endpoint, "movies/boxoffice")

	var movies []*BoxOfficeMovie
	if err := t.request(url, NewQuery(qo), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

// PlayedMovies returns the mosts played movies
func (t *TraktTv) PlayedMovies(qo QueryOption) ([]*PlayedMovie, error) {

	url := fmt.Sprintf("%s/%s", t.Endpoint, "movies/played")

	var movies []*PlayedMovie
	if err := t.request(url, NewQuery(qo), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

// WatchedMovies returns the mosts played movies by unique users
func (t *TraktTv) WatchedMovies(qo QueryOption) ([]*PlayedMovie, error) {

	url := fmt.Sprintf("%s/%s", t.Endpoint, "movies/watched")

	var movies []*PlayedMovie
	if err := t.request(url, NewQuery(qo), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}

// CollectedMovies returns the mosts collected movies
func (t *TraktTv) CollectedMovies(qo QueryOption) ([]*PlayedMovie, error) {

	url := fmt.Sprintf("%s/%s", t.Endpoint, "movies/collected")

	var movies []*PlayedMovie
	if err := t.request(url, NewQuery(qo), &movies); err != nil {
		return nil, err
	}

	return movies, nil
}
