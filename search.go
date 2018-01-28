package gotwitter

// StandardSearchTweets returns a collection of relevant Tweets matching a specified query.
// See more at https://developer.twitter.com/en/docs/tweets/search/api-reference/get-search-tweets
func (app *Application) StandardSearchTweets(param RequestParameters) (results *SearchResult, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	results = &SearchResult{}
	err = app.doRequest(req, &results)

	return
}
