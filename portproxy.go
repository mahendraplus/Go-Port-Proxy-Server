package main

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "strconv"
    "strings"
)

func main() {
    // Create a new HTTP server on port 8022
    http.HandleFunc("/", proxyHandler)
    log.Println("Starting server on :2701")
    err := http.ListenAndServe(":2701", nil)
    if err != nil {
        log.Fatalf("ListenAndServe: %v", err)
    }
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the port from the URL path
    pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
    if len(pathParts) == 0 {
        http.Error(w, "Invalid request path", http.StatusBadRequest)
        return
    }

    portStr := pathParts[0]
    port, err := strconv.Atoi(portStr)
    if err != nil || port < 0 || port > 65535 {
        http.Error(w, "Invalid port", http.StatusBadRequest)
        return
    }

    // Construct the destination URL
    targetURL := &url.URL{
        Scheme: "http",
        Host:   "127.0.0.1:" + portStr,
        Path:   "/" + strings.Join(pathParts[1:], "/"),
    }

    // Proxy the request to the target URL
    proxy := httputil.NewSingleHostReverseProxy(targetURL)
    proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
        log.Printf("proxy error: %v", err)
        http.Error(w, "Proxy error", http.StatusBadGateway)
    }

    r.URL.Path = targetURL.Path
    r.Host = targetURL.Host
    proxy.ServeHTTP(w, r)
} // ⚡😘
