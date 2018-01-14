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
	followerIDsBaseURL = "https://api.twitter.com/1.1/followers/ids.json?"
)

// FollowerIDs retrieval follower ids.
// See more at https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-ids
func (app *Application) FollowerIDs(param *FollowerIDsParam) (*FollowerIDs, error) {
	// NOTICE: stringify must be false
	param.StringifyIDs = false

	req, err := app.followerIDsReq(param)
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

	ids := FollowerIDs{}
	if app.debugLevel > 1 {
		if buff, err := ioutil.ReadAll(gzipReader); err == nil {
			fmt.Printf("[DEBUG 2]*Application.FollowerIDs() buff <---> %s\n", string(buff))
			if err := json.NewDecoder(bytes.NewBuffer(buff)).Decode(&ids); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		if err := json.NewDecoder(gzipReader).Decode(&ids); err != nil {
			return nil, err
		}
	}

	return &ids, nil
}

func (app *Application) followerIDsReq(param *FollowerIDsParam) (*http.Request, error) {
	v, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	qstr := followerIDsBaseURL + v.Encode()
	req, err := http.NewRequest("GET", qstr, nil)
	if err != nil {
		return nil, err
	}

	addHeadersForRetrieveRequest(req, app)

	return req, nil
}
