package gotwitter

import (
	"testing"
)

func TestNewApplication(t *testing.T) {
	app, _ := NewApplication(2, "test1", "test2")
	t.Log(app)
	app, _ = NewApplication(2, "test1", "test2", "test3")
	t.Log(app)
	app, _ = NewApplication(2)
	t.Log(app)
}
