package main

import (
	"fmt"
	"net/http"
)

func GreetingServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, Greeting("Code.education Rocks!"))
}