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
