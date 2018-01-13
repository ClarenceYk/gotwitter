package gotwitter

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	userTimelineBaseURL = "https://api.twitter.com/1.1/statuses/user_timeline.json?"
)

// UserTimeline fetch a timeline of a special user.
// See more at https://developer.twitter.com/en/docs/tweets/timelines/api-reference/get-statuses-user_timeline
func (app *Application) UserTimeline(param *UserTimelineParam) ([]*Tweet, error) {
	req, err := app.userTimelineReq(param)
	if err != nil {
		return nil, err
	}

	resp, err := app.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, app.requestError(resp)
	}

	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}

	ts := make([]*Tweet, 0)
	if app.debugLevel > 1 {
		if buff, err := ioutil.ReadAll(gzipReader); err == nil {
			fmt.Printf("[DEBUG 2]*Application.UserTimeline() buff <---> %s\n", string(buff))
			if err := json.NewDecoder(bytes.NewBuffer(buff)).Decode(&ts); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		if err := json.NewDecoder(gzipReader).Decode(&ts); err != nil {
			return nil, err
		}
	}

	return ts, nil
}

func (app *Application) userTimelineReq(param *UserTimelineParam) (*http.Request, error) {
	v, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	qstr := userTimelineBaseURL + v.Encode()
	req, err := http.NewRequest("GET", qstr, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", app.Name)
	req.Header.Add("Authorization", "Bearer "+app.BearerToken)
	req.Header.Add("Accept-Encoding", "gzip")

	return req, nil
}
