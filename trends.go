package gotwitter

// TrendsPlace returns the top 50 trending topics for a specific WOEID, if trending information is available for it.
// See more at https://developer.twitter.com/en/docs/trends/trends-for-location/api-reference/get-trends-place
func (app *Application) TrendsPlace(param RequestParameters) (ts []*Trend, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	ts = make([]*Trend, 0)
	err = app.doRequest(req, &ts)

	return
}

// TrendsAvailable returns the locations that Twitter has trending topic information for.
// See more at https://developer.twitter.com/en/docs/trends/locations-with-trending-topics/api-reference/get-trends-available
func (app *Application) TrendsAvailable() (ts []*TrendAvailable, err error) {
	req, err := app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/trends/available.json",
		struct{}{},
	)
	if err != nil {
		return
	}

	ts = make([]*TrendAvailable, 0)
	err = app.doRequest(req, &ts)

	return
}

// TrendsClosest returns the top 50 trending topics for a specific WOEID, if trending information is available for it.
// See more at https://developer.twitter.com/en/docs/trends/locations-with-trending-topics/api-reference/get-trends-closest
func (app *Application) TrendsClosest(param RequestParameters) (ts []*TrendClosest, err error) {
	req, err := param.GetRequest(app)
	if err != nil {
		return
	}

	ts = make([]*TrendClosest, 0)
	err = app.doRequest(req, &ts)

	return
}
