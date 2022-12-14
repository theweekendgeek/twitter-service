package twitter

import (
	"doescher.ninja/twitter-service/config"
	"fmt"
	"strings"
)

// QueryOptions configure query parameters for timeline requests
type QueryOptions struct {
	Max             int    // maximum number of tweets to request
	SinceID         string // get tweets newer than this id
	PaginationToken string // paginate the timeline
}

func buildQueryParams(options QueryOptions) string {
	var queryParams []string
	var maxQuery int

	if options.Max != 0 {
		maxQuery = options.Max
	} else {
		maxQuery = config.Const().MaxTweets
	}
	queryParams = append(queryParams, "max_results="+fmt.Sprintf("%d", maxQuery))

	if options.SinceID != "" {
		queryParams = append(queryParams, "since_id="+options.SinceID)
	}

	if options.PaginationToken != "" {
		queryParams = append(queryParams, "pagination_token="+options.PaginationToken)
	}

	queryString := strings.Join(queryParams, "&")
	if len(queryString) > 0 {
		queryString = "?" + queryString
	}
	return queryString
}

func buildTimelineURL(options QueryOptions) string {
	return config.Const().EndpointTimelineByID + buildQueryParams(options)
}
