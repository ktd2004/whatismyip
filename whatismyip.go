package main

import (
	"fmt"
	"net"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	ipAddress, _, _ := net.SplitHostPort(r.RemoteAddr)
	fmt.Fprintf(w, "%s", ipAddress)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
