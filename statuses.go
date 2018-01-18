package gotwitter

import "net/http"

// UserTimeline fetch a timeline of a special user.
// See more at https://developer.twitter.com/en/docs/tweets/timelines/api-reference/get-statuses-user_timeline
func (app *Application) UserTimeline(param *UserTimelineParam) (ts []*Tweet, err error) {
	req, err := app.userTimelineReq(param)
	if err != nil {
		return
	}

	ts = make([]*Tweet, 0)
	err = app.doRequest(req, &ts)

	return
}

func (app *Application) userTimelineReq(param *UserTimelineParam) (*http.Request, error) {
	return app.getRequest(
		"GET",
		"https://api.twitter.com/1.1/statuses/user_timeline.json",
		param,
	)
}
