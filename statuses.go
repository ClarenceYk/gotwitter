package gotwitter

// UserTimeline fetch a timeline of a special user.
// See more at https://developer.twitter.com/en/docs/tweets/timelines/api-reference/get-statuses-user_timeline
func (app *Application) UserTimeline(param RequestParameters) (ts []*Tweet, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	ts = make([]*Tweet, 0)
	err = app.doRequest(req, &ts)

	return
}

// StatusesShow returns single tweet, specified by the id parameter.
// See more at https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/get-statuses-show-id
func (app *Application) StatusesShow(param RequestParameters) (tweet *Tweet, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	tweet = &Tweet{}
	err = app.doRequest(req, &tweet)

	return
}

// StatusesLookup returns fully-hydrated Tweet objects for up to 100 Tweets per request.
// See more at https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/get-statuses-lookup
func (app *Application) StatusesLookup(param RequestParameters) (ts []*Tweet, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	ts = make([]*Tweet, 0)
	err = app.doRequest(req, &ts)

	return
}

// StatusesRetweets returns a collection of the 100 most recent retweets of the Tweet specified by the id parameter.
// See more at https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/get-statuses-retweets-id
func (app *Application) StatusesRetweets(param RequestParameters) (ts []*Tweet, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	ts = make([]*Tweet, 0)
	err = app.doRequest(req, &ts)

	return
}

// StatusesRetweeters returns a collection of up to 100 user IDs belonging to users who have retweeted the Tweet specified by the id parameter.
// See more at https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/get-statuses-retweeters-ids
func (app *Application) StatusesRetweeters(param RequestParameters) (rts *Retweeters, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	rts = &Retweeters{}
	err = app.doRequest(req, &rts)

	return
}

// StatusesFavorites returns the 20 most recent Tweets liked by the authenticating or specified user.
// See more at https://developer.twitter.com/en/docs/tweets/post-and-engage/api-reference/get-favorites-list
func (app *Application) StatusesFavorites(param RequestParameters) (ts []*Tweet, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	ts = make([]*Tweet, 0)
	err = app.doRequest(req, &ts)

	return
}
