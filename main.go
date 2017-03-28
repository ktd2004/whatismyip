package main

import (
	"flag"
	"fmt"
	"log"
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
	err := http.ListenAndServe(hostAndPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
