package gotwitter

import (
	"testing"
)

func TestNewApplication(t *testing.T) {
	app, _ := NewApplication(0)
	err := app.Authorize()
	if err != nil {
		t.Log(err)
	}
	t.Log(app)
}
