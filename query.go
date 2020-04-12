package trakttv

import (
	"net/url"
	"strconv"
	"strings"
)

// SearchQuery represents a search query
type SearchQuery struct {
	Query   string
	Type    Type
	Year    int
	Ratings string
	Field   Field
}

// ExtendedInfo represents the params you can add to a query to get more
// detailed informations
type ExtendedInfo string

// Available ExtendedInfo
const (
	ExtendedInfoMin      ExtendedInfo = "min"
	ExtendedInfoImages   ExtendedInfo = "images"
	ExtendedInfoFull     ExtendedInfo = "full"
	ExtendedInfoEpisodes ExtendedInfo = "episodes"
	ExtendedInfoMetadata ExtendedInfo = "metadata"
)

// Pagination represents the pagination options you can add to a query
type Pagination struct {
	Page  int
	Limit int
}

// QueryOption represents the query options
type QueryOption struct {
	ExtendedInfos []ExtendedInfo
	Pagination    Pagination
}

func (qo *QueryOption) extendedInfos() []string {
	infos := []string{}
	for _, ei := range qo.ExtendedInfos {
		infos = append(infos, string(ei))
	}
	return infos
}

// NewQuery returns a new query from a query option
func NewQuery(qo QueryOption) *Query {
	return &Query{
		QueryOption: qo,
	}
}

// Query represents a query
type Query struct {
	sq *SearchQuery
	QueryOption
}

func (q *Query) urlValues() *url.Values {
	v := &url.Values{}

	if len(q.ExtendedInfos) > 0 {
		v.Add("extended", strings.Join(q.extendedInfos(), ","))
	}

	if q.Pagination.Page != 0 {
		v.Add("page", strconv.Itoa(q.Pagination.Page))
	}

	if q.Pagination.Limit != 0 {
		v.Add("limit", strconv.Itoa(q.Pagination.Limit))
	}

	if q.sq == nil {
		return v
	}

	if q.sq.Query != "" {
		v.Add("query", q.sq.Query)
	}

	if q.sq.Year != 0 {
		v.Add("year", strconv.Itoa(q.sq.Year))
	}

	if q.sq.Ratings != "" {
		v.Add("ratings", q.sq.Ratings)
	}
	if q.sq.Field != "" {
		v.Add("fields", string(q.sq.Field))
	}

	return v
}
