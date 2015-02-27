package main

import (
    "fmt"
    "regexp"
    "net/http"
) 

// Define HTTP route structure for the server
type Route struct {
    pattern *regexp.Regexp
    method  string
    handler http.Handler
}




