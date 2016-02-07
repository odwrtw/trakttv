package trakttv

import (
	"net/url"
	"reflect"
	"testing"
)

func TestQuery(t *testing.T) {

	for _, mock := range []struct {
		query    *Query
		expected *url.Values
	}{
		{
			query: &Query{
				QueryOption: QueryOption{
					ExtendedInfos: []ExtendedInfo{
						ExtendedInfoImages,
						ExtendedInfoFull,
					},
					Pagination: Pagination{
						Page:  2,
						Limit: 100,
					},
				},
			},
			expected: &url.Values{
				"extended": []string{"images,full"},
				"page":     []string{"2"},
				"limit":    []string{"100"},
			},
		},
		{
			query: &Query{
				QueryOption: QueryOption{
					ExtendedInfos: []ExtendedInfo{
						ExtendedInfoMin,
					},
					Pagination: Pagination{
						Page: 1,
					},
				},
			},
			expected: &url.Values{
				"extended": []string{"min"},
				"page":     []string{"1"},
			},
		},
		{
			query: &Query{
				QueryOption: QueryOption{
					Pagination: Pagination{
						Limit: 10,
					},
				},
			},
			expected: &url.Values{
				"limit": []string{"10"},
			},
		},
	} {
		got := mock.query.urlValues()
		if !reflect.DeepEqual(mock.expected, got) {
			t.Errorf("invalid query option, got %#v, expected %#v", got, mock.expected)
		}
	}
}
