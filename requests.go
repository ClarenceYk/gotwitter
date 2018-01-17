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

func (app *Application) getRequest(baseURL string, param interface{}) (*http.Request, error) {
	v, err := query.Values(param)
	if err != nil {
		return nil, err
	}

	urlParamStr := v.Encode()
	if urlParamStr != "" {
		urlParamStr = "?" + urlParamStr
	}
	qstr := baseURL + urlParamStr
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

func (app *Application) doRequest(req *http.Request, result interface{}) error {
	resp, err := app.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return app.requestError(resp)
	}

	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}

	var reader io.Reader
	if app.debugLevel > 1 {
		reader = debugReader{gzipReader}
	} else {
		reader = normalReader{gzipReader}
	}

	if err := json.NewDecoder(reader).Decode(result); err != nil {
		return err
	}

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
