package gotwitter

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Application holds the information of your twitter application.
type Application struct {
	ConsumerKey    string
	ConsumerSecret string
	Name           string
	BearerToken    string
	debugLevel     int
	httpClient     *http.Client
}

// NewApplication creates a new application.
// If you pass 2 parameters to *params* that means you
// passed your consumer key and consumer secret to
// this funcation(consumer key followed by consumer secret).
//
// If you pass 3 paramters to *params* that means you passed
// your consumer key, consumer secret and application name to
// this function(consumer key followed by consumer secret and name).
//
// You can pass noting to this funcation so that the function will
// obtain the infos from keys.conf file automatically.
// The format of keys.conf file is like:
//   consumerKey=<your consumer key>
//   consumerSecret=<your consumer screte>
//   name=<your application name> /*optional*/
//   token=<your authorized token> /*optional*/
func NewApplication(debug int, params ...string) (*Application, error) {
	var (
		consumerKey    string
		consumerSecret string
		name           string
		token          string
	)

	if len(params) == 2 {
		consumerKey = params[0]
		consumerSecret = params[1]
	} else if len(params) == 3 {
		consumerKey = params[0]
		consumerSecret = params[1]
		name = params[2]
	} else {
		configs, err := getConfig(debug, "keys.conf")
		if err != nil {
			return nil, err
		}

		for k, v := range configs {
			switch k {
			case "consumerKey":
				consumerKey = v
			case "consumerSecret":
				consumerSecret = v
			case "name":
				name = v
			case "token":
				token = v
			}
		}
	}

	return &Application{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		Name:           name,
		BearerToken:    token,
		debugLevel:     debug,
		httpClient:     &http.Client{},
	}, nil
}

// Authorize this application.
// See more at
// https://developer.twitter.com/en/docs/basics/authentication/overview/application-only
func (app *Application) Authorize() error {
	req, err := app.request()
	if err != nil {
		return err
	}
	resp, err := app.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return app.requestError(resp)
	}

	token := struct {
		TokenType   string `json:"token_type"`
		AccessToken string `json:"access_token"`
	}{}
	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return fmt.Errorf("error<*Application.Authorize> err=%s", err)
	}
	if app.debugLevel > 1 {
		if buff, err := ioutil.ReadAll(gzipReader); err == nil {
			fmt.Printf("[DEBUG 2]*Application.Authorize() buff <---> %s\n", string(buff))
			if err := json.NewDecoder(bytes.NewBuffer(buff)).Decode(&token); err != nil {
				return fmt.Errorf("error<*Application.Authorize> err=%s", err)
			}
		} else {
			return fmt.Errorf("error<*Application.Authorize> err=%s", err)
		}
	} else {
		if err := json.NewDecoder(gzipReader).Decode(&token); err != nil {
			return fmt.Errorf("error<*Application.Authorize> err=%s", err)
		}
	}
	app.BearerToken = token.AccessToken
	return nil
}

// requestError handle the unexpected status.
func (app *Application) requestError(resp *http.Response) error {
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
			fmt.Printf("[DEBUG 2]*Application.authorizeError() buff <---> %s\n", string(buff))
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

func (app *Application) tokenCredentials() (base64Encoded string) {
	encodedCK := url.QueryEscape(app.ConsumerKey)
	encodedCS := url.QueryEscape(app.ConsumerSecret)
	bearer := encodedCK + ":" + encodedCS
	base64Encoded = base64.StdEncoding.EncodeToString([]byte(bearer))
	if app.debugLevel > 1 {
		fmt.Printf("[DEBUG 2]*Application.tokenCredentials(): base64Encoded <---> %s\n",
			base64Encoded,
		)
	}
	return
}

func (app *Application) request() (*http.Request, error) {
	req, err := http.NewRequest("POST",
		"https://api.twitter.com/oauth2/token",
		strings.NewReader("grant_type=client_credentials"),
	)
	if err != nil {
		return nil, err
	}

	addHeadersForAuthRequest(req, app)

	return req, nil
}
