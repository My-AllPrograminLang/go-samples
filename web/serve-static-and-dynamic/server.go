// Serves both static and dynamic content with a Go http server.
//
// Similar to the web/http-ajax.go sample, except that the HTML/CSS are served
// from the filesystem ("static"), whereas the dynamic server is done inline in
// Go.

// TODO: move CSS into its own dir too
// TODO: make it a valid HTML5 file

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// handler to cater AJAX requests
func handlerGetTime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"))
}

func main() {
	http.HandleFunc("/gettime", handlerGetTime)
	http.Handle("/", http.FileServer(http.Dir("public/html")))
	http.Handle("/css", http.FileServer(http.Dir("public/css")))
	log.Fatal(http.ListenAndServe(":9999", nil))
}
