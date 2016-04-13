// http4 sample modified to solve exercises + misc. enhancements.
//

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func main() {
	db := database{db: map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/setprice", db.setprice)

	portnum := 8000
	if len(os.Args) > 1 {
		portnum, _ = strconv.Atoi(os.Args[1])
	}
	log.Printf("Going to listen on port %d\n", portnum)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portnum), nil))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

// For ex. 7.11 adding a mutex to guard the map.
type database struct {
	mu sync.Mutex // guards db
	db map[string]dollars
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	log.Printf("list %v", req)
	db.mu.Lock()
	defer db.mu.Unlock()
	for item, price := range db.db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	log.Printf("price %v", req)
	item := req.URL.Query().Get("item")
	db.mu.Lock()
	defer db.mu.Unlock()
	if price, ok := db.db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

// Exercise 7.11
func (db database) setprice(w http.ResponseWriter, req *http.Request) {
	log.Printf("setprice %v, %v", req, req.URL.Query())
	item := req.URL.Query().Get("item")
	newprice, _ := strconv.ParseFloat(req.URL.Query().Get("newprice"), 32)
	db.mu.Lock()
	defer db.mu.Unlock()
	db.db[item] = dollars(newprice)
	fmt.Fprintf(w, "%s=%f\n", item, newprice)
}
