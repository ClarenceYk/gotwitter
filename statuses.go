package gotwitter

import (
	"encoding/json"
)

// UserTimeline fetch a timeline of a special user.
// See more at https://developer.twitter.com/en/docs/tweets/timelines/api-reference/get-statuses-user_timeline
func (app *Application) UserTimeline(param *UserTimelineParam) ([]*Tweet, error) {
	req, err := app.userTimelineReq(param)
	if err != nil {
		return nil, err
	}

	grc, err := app.doRequest(req)
	if err != nil {
		return nil, err
	}
	defer grc.Close()

	ts := make([]*Tweet, 0)
	if err := json.NewDecoder(grc).Decode(&ts); err != nil {
		return nil, err
	}

	return ts, nil
}
