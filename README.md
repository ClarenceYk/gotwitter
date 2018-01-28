# GoTwitter - A light and simple to use Twitter SDK implemented using Go

*Notice: This project is under deveplopment only implement Application-only authentication after the last commit.*

# Getting started

## Install:

```
git clone https://github.com/ClarenceYk/gotwitter.git
```

## Usage:

```go
package main

import (
	"fmt"
	"log"

	"github.com/ClarenceYk/gotwitter"
)

func main() {
	app, err := gotwitter.NewApplication(
		0, // debug level
		"YourConsumerKey",
		"YourConsumerSecret",
		"YourAppName",
	)

	if err != nil {
		log.Fatalln(err)
	}

	param := &gotwitter.UserTimelineParam{
		ScreenName: "MaYukkee",
		Count:      5,
	}

	tweets, err := app.UserTimeline(param)
	if err != nil {
		log.Fatalln(err)
	} else {
		for _, tweet := range tweets {
			fmt.Println(tweet.Text)
			if tweet.QuotedStatus != nil {
				fmt.Println(tweet.QuotedStatus.Text)
			}
			if tweet.RetweetedStatus != nil {
				fmt.Println(tweet.RetweetedStatus.Text)
			}
		}
	}
}
```

# Examples

More examples is available in `application_test.go` file.

