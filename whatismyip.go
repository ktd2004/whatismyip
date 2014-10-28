package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
)

var port = flag.Int("port", 8080, "Port to use, defaults to 8080")

func handler(w http.ResponseWriter, r *http.Request) {
	ipAddress, _, _ := net.SplitHostPort(r.RemoteAddr)
	fmt.Fprintf(w, "%s", ipAddress)
}

func main() {
	flag.Parse()

	hostAndPort := fmt.Sprintf(":%d", *port)

	http.HandleFunc("/", handler)
	http.ListenAndServe(hostAndPort, nil)
}
