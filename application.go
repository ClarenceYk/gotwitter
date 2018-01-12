package gotwitter

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Application holds the information of your twitter application.
type Application struct {
	ConsumerKey    string
	ConsumerSecret string
	Name           string
	BearerToken    string
	debugLevel     int
}

// NewApplication creates a new application.
// If you pass 2 parameters to *params* that means you
// passed your consumer key and consumer secret to
// this funcation(consumer key followed by consumer secret).
//
// If you pass 3 paramters to *params* that meas you passed
// your consumer key, consumer secret and application name to
// this function(consumer key followed by consumer secret and name).
//
// You also can pass noting to this funcation and it will obtain the keys from
// keys.conf file. The format of keys.conf file is like:
//   consumerKey=<your consumer key>
//   consumerSecret=<your consumer screte>
//   name=<your application name> /*optional*/
func NewApplication(debug int, params ...string) (*Application, error) {
	var (
		consumerKey    string
		consumerSecret string
		name           string
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

		if cfg, ok := configs["consumerKey"]; ok {
			consumerKey = cfg
		}
		if cfg, ok := configs["consumerSecret"]; ok {
			consumerSecret = cfg
		}
		if cfg, ok := configs["name"]; ok {
			name = cfg
		}
	}

	return &Application{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		Name:           name,
		debugLevel:     debug,
	}, nil
}

// Authorize this application.
// See more at
// https://developer.twitter.com/en/docs/basics/authentication/overview/application-only
func (app *Application) Authorize() error {
	httpClient := &http.Client{}
	req, err := app.request()
	if err != nil {
		return err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

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
	req.Header.Add("User-Agent", app.Name)
	req.Header.Add("Authorization", "Basic "+app.tokenCredentials())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Accept-Encoding", "gzip")

	return req, nil
}

func getConfig(debug int, filename string) (map[string]string, error) {
	configs := make(map[string]string)

	// Read content of .conf file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	br := bufio.NewReader(file)
	for {
		line, err := br.ReadString('\n')
		line = strings.Trim(line, "\n") // remove '\n'
		if debug > 1 {
			fmt.Printf("[DEBUG 2]NewApplication(): read line from .conf flie <---> %s\n", line)
		}

		if strings.Contains(line, "=") {
			kv := strings.Split(line, "=")
			configs[kv[0]] = kv[1]
		}

		if err != nil {
			break
		}
	}
	return configs, nil
}
