package gotwitter

import (
	"bufio"
	"fmt"
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
			fmt.Printf("NewApplication(): read line from .conf flie <---> %s\n", line)
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
