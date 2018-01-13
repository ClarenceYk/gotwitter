package gotwitter

import (
	"testing"
)

func TestNewApplication(t *testing.T) {
	app, _ := NewApplication(3)
	err := app.Authorize()
	if err != nil {
		t.Log(err)
	}
	t.Log(app)
}
