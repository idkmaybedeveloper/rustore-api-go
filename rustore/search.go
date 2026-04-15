package rustore

import (
	"fmt"
	"net/url"
)

func GetSearchSuggestions(query string) (SearchSuggestResponse, error) {
	return apiGet[SearchSuggestResponse](fmt.Sprintf("/search/suggest?query=%s", url.QueryEscape(query)))
}

func SearchApps(query string, pageNumber, pageSize int) (SearchResponse, error) {
	params := url.Values{}
	params.Set("query", query)
	params.Set("pageNumber", fmt.Sprintf("%d", pageNumber))
	params.Set("pageSize", fmt.Sprintf("%d", pageSize))
	return apiGet[SearchResponse]("/applicationData/apps?" + params.Encode())
}
