package gotwitter

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/google/go-querystring/query"
)

const (
	userTimelineBaseURL   = "https://api.twitter.com/1.1/statuses/user_timeline.json"
	followerIDsBaseURL    = "https://api.twitter.com/1.1/followers/ids.json"
	followersListBaseURL  = "https://api.twitter.com/1.1/followers/list.json"
	friendIDsBaseURL      = "https://api.twitter.com/1.1/friends/ids.json"
	friendsListBaseURL    = "https://api.twitter.com/1.1/friends/list.json"
	showFriendshipBaseURL = "https://api.twitter.com/1.1/friendships/show.json"
)

func (app *Application) userTimelineReq(param *UserTimelineParam) (*http.Request, error) {
	return app.getRequest(userTimelineBaseURL, param)
}

func (app *Application) followerIDsReq(param *FollowerIDsParam) (*http.Request, error) {
	return app.getRequest(followerIDsBaseURL, param)
}

func (app *Application) followersListReq(param *FollowersListParam) (*http.Request, error) {
	return app.getRequest(followersListBaseURL, param)
}

func (app *Application) friendIDsReq(param *FriendIDsParam) (*http.Request, error) {
	return app.getRequest(friendIDsBaseURL, param)
}

func (app *Application) friendsListReq(param *FriendsListParam) (*http.Request, error) {
	return app.getRequest(friendsListBaseURL, param)
}

func (app *Application) showFriendshipReq(param *ShowFriendshipParam) (*http.Request, error) {
	return app.getRequest(showFriendshipBaseURL, param)
}

func (app *Application) getRequest(baseURL string, param interface{}) (*http.Request, error) {
	v, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	qstr := baseURL + "?" + v.Encode()
	if app.debugLevel > 0 {
		fmt.Printf("[DEBUG 1]getRequest() query <---> %s\n", qstr)
	}

	req, err := http.NewRequest("GET", qstr, nil)
	if err != nil {
		return nil, err
	}

	addHeadersForRetrieveRequest(req, app.Name, app.BearerToken)

	return req, nil
}

func (app *Application) doRequest(req *http.Request) (io.ReadCloser, error) {
	resp, err := app.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, app.requestError(resp)
	}

	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}

	if app.debugLevel > 1 {
		return debugReadCloser{gzipReader, resp.Body}, nil
	}

	return normalReadCloser{gzipReader, resp.Body}, nil
}

// requestError handle the unexpected status.
func (app *Application) requestError(resp *http.Response) error {
	defer resp.Body.Close()
	errors := struct {
		Errs []struct {
			Code    int    `json:"code"`
			Label   string `json:"label"`
			Message string `json:"message"`
		} `json:"errors"`
	}{}

	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	if app.debugLevel > 1 {
		if buff, err := ioutil.ReadAll(gzipReader); err == nil {
			fmt.Printf("[DEBUG 2]*Application.requestError() buff <---> %s\n", string(buff))
			if err := json.NewDecoder(bytes.NewBuffer(buff)).Decode(&errors); err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		if err := json.NewDecoder(gzipReader).Decode(&errors); err != nil {
			return err
		}
	}

	errStr := ""
	for i, err := range errors.Errs {
		errStr = fmt.Sprintf("err_%d=[code]%d;[label]%s;[msg]%s %s",
			i,
			err.Code,
			err.Label,
			err.Message,
			errStr,
		)
	}
	return fmt.Errorf("%s", errStr)
}
