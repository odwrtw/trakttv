package trakttv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// Test shit
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

	req.Header.Add("Content-type", "application/json")
	req.Header.Add("trakt-api-key", t.Key)
	req.Header.Add("trakt-api-version", "2")

	resp, err := t.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if os.Getenv("DEBUG") == "1" {
		fmt.Println("debug mode")
		fmt.Println("===============")
		fmt.Printf("URL: %q\n", URL)
		fmt.Println("===============")
		fmt.Printf("Status : %q", resp.Status)

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(body))
		fmt.Println("===============")

		return json.Unmarshal(body, result)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}
