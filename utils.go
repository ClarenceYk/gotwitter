package gotwitter

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

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
		line = strings.Trim(line, " \n") // remove '\n'
		if debug > 1 {
			fmt.Printf("[DEBUG 2]getConfig(): read line from .conf flie <---> %s\n", line)
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

func addHeadersForAuthRequest(req *http.Request, app *Application) {
	addHeadersForAllRequest(req, app)
	req.Header.Add("Authorization", "Basic "+app.tokenCredentials())
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
}

func addHeadersForRetrieveRequest(req *http.Request, app *Application) {
	addHeadersForAllRequest(req, app)
	req.Header.Add("Authorization", "Bearer "+app.BearerToken)
}

func addHeadersForAllRequest(req *http.Request, app *Application) {
	req.Header.Add("User-Agent", app.Name)
	req.Header.Add("Accept-Encoding", "gzip")
}
