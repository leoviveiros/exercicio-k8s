package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting("Code.education Rocks!")
	want := "<b>Code.education Rocks!</b>"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}