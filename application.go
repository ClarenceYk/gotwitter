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
	DebugLevel     int
	httpClient     *http.Client
}

// NewApplication creates a new application.
func NewApplication(debug int, key, secret, name string) (*Application, error) {
	app := &Application{
		ConsumerKey:    key,
		ConsumerSecret: secret,
		Name:           name,
		DebugLevel:     debug,
		httpClient:     &http.Client{},
	}

	if err := app.Authorize(); err != nil {
		return nil, err
	}

	return app, nil
}

// NewApplicationWithToken creates a new application.
func NewApplicationWithToken(debug int, key, secret, name, token string) (*Application, error) {
	return &Application{
		ConsumerKey:    key,
		ConsumerSecret: secret,
		Name:           name,
		BearerToken:    token,
		DebugLevel:     debug,
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
	if app.DebugLevel > 1 {
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

func (app *Application) tokenCredentials() (base64Encoded string) {
	encodedCK := url.QueryEscape(app.ConsumerKey)
	encodedCS := url.QueryEscape(app.ConsumerSecret)
	bearer := encodedCK + ":" + encodedCS
	base64Encoded = base64.StdEncoding.EncodeToString([]byte(bearer))
	if app.DebugLevel > 1 {
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

	addHeadersForAuthRequest(req, app.Name, app.tokenCredentials())

	return req, nil
}
