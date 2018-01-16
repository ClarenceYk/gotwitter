package gotwitter

import (
	"bufio"
	"fmt"
	"io"
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

func addHeadersForAuthRequest(req *http.Request, appName, credentials string) {
	addHeadersForAllRequest(req, appName)
	req.Header.Add("Authorization", "Basic "+credentials)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
}

func addHeadersForRetrieveRequest(req *http.Request, appName, token string) {
	addHeadersForAllRequest(req, appName)
	req.Header.Add("Authorization", "Bearer "+token)
}

func addHeadersForAllRequest(req *http.Request, appName string) {
	req.Header.Add("User-Agent", appName)
	req.Header.Add("Accept-Encoding", "gzip")
}

type debugReader struct {
	r io.Reader
}

func (drc debugReader) Read(buff []byte) (int, error) {
	n, err := drc.r.Read(buff)
	if err != nil {
		return n, err
	}
	fmt.Printf("[DEBUG 2] Buffer[0:%d] <---> %s\n", n, string(buff[:n]))
	return n, nil
}

type normalReader struct {
	r io.Reader
}

func (nrc normalReader) Read(buff []byte) (int, error) {
	n, err := nrc.r.Read(buff)
	if err != nil {
		return n, err
	}
	return n, nil
}
