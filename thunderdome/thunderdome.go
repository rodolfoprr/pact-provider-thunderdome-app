package thunderdome

import (
	"fmt"
	"net/http"
)

// Start Thunderdome!
func Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/fighter/1", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")
		fmt.Fprintf(w, `{"Name":"Mel Gibson"}`)
	})

	go http.ListenAndServe(":8000", mux)
}
