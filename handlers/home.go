package handlers

import (
	"fmt"
	"net/http"
)

// home is a simple HTTP handler function which writes a response.

func home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello! your request was processed")
}