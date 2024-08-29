package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
) // I CAN'T SLEEP...... 

func main() {
	// Start the proxy server on port 2701
	http.HandleFunc("/", proxyHandler)
	log.Println("Starting server on :2701")
	if err := http.ListenAndServe(":2701", nil); err != nil {
		log.Fatalf("ListenAndServe: %v", err)
	}
}

// proxyHandler handles incoming requests and forwards them to the appropriate port and path
func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// Split the request path to extract the port and the rest of the path
	pathParts := strings.SplitN(strings.TrimPrefix(r.URL.Path, "/"), "/", 2)
	if len(pathParts) < 1 {
		http.Error(w, "Invalid request path", http.StatusBadRequest)
		return
	}

	// The first part of the path is assumed to be the port number
	portStr := pathParts[0]
	restOfPath := "/"
	if len(pathParts) == 2 {
		restOfPath = "/" + pathParts[1]
	}

	// Construct the target URL for the reverse proxy
	targetURL := &url.URL{
		Scheme: "http",
		Host:   "127.0.0.1:" + portStr,
		Path:   restOfPath,
	}

	log.Printf("Proxying request to %s", targetURL.String())

	// Create a reverse proxy to forward the request
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = "127.0.0.1:" + portStr
		req.URL.Path = restOfPath
		req.Host = req.URL.Host
	}
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Proxy error: %v", err)
		http.Error(w, "Proxy error", http.StatusBadGateway)
	}

	// Forward the request to the target server
	proxy.ServeHTTP(w, r)
}
