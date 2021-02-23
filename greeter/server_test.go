package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestGreetingServer(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	GreetingServer(response, request)

	got := response.Body.String()
	want := "<b>Code.education Rocks!</b>"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}