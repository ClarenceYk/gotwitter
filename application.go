package gotwitter

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

// Application holds the information of your twitter application.
type Application struct {
	ConsumerKey    string
	ConsumerSecret string
	BearerToken    string
	debugLevel     int
}

// NewApplication creates a new application.
// If you pass 2 parameters to this function that
// means you passed your consumer key and consumer secret to
// this funcation(consumer key followed by consumer secret).
// You also can pass noting to this funcation and it will obtain the keys from
// keys.conf file. The format of keys.conf file is like:
// 		consumerKey=<your consumer key>
// 		consumerSecret=<your consumer screte>
func NewApplication(debug int, keys ...string) (*Application, error) {
	var (
		consumerKey    string
		consumerSecret string
	)

	if len(keys) == 2 {
		consumerKey = keys[0]
		consumerSecret = keys[1]
	} else {
		// Read content of .conf file
		content, err := ioutil.ReadFile("keys.conf")
		if err != nil {
			return nil, err
		}

		keyMap := make(map[string]string)
		br := bufio.NewReader(bytes.NewBuffer(content))
		for {
			line, err := br.ReadString('\n')
			line = strings.Trim(line, "\n") // remove '\n'
			if debug > 1 {
				fmt.Printf("NewApplication(): read line from .conf flie <---> %s\n", line)
			}

			if strings.Contains(line, "=") {
				kv := strings.Split(line, "=")
				keyMap[kv[0]] = kv[1]
			}

			if err != nil {
				break
			}
		}
		consumerKey = keyMap["consumerKey"]
		consumerSecret = keyMap["consumerSecret"]
	}

	return &Application{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		debugLevel:     debug,
	}, nil
}
