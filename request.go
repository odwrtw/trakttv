package trakttv

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

// request performs a GET request to the given URL, appending query options,
// and decodes the JSON response into result.
func (t *TraktTv) request(URL string, q *Query, result interface{}) error {
	// Add the query options to the URL
	u, err := url.Parse(URL)
	if err != nil {
		return err
	}
	u.RawQuery = q.urlValues().Encode()
	URL = u.String()

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("trakt-api-key", t.Key)
	req.Header.Add("trakt-api-version", strconv.Itoa(t.Version))

	resp, err := t.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	return json.NewDecoder(resp.Body).Decode(result)
}
