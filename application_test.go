package gotwitter

import (
	"testing"
)

func TestNewApplication(t *testing.T) {
	app, _ := NewApplication(3)
	app.Authorize()
	t.Log(app)
}
