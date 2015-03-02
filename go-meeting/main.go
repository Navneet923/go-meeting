
// This is a very minimal file to invoke our server. This allows
// us to produce a binary which is named "go-meeting" as opposed to
// "server" which is generic and misleading
package main

import (
    "fmt"
    "github.com/sabhiram/go-meeting/server"
    "net/http"
)

func main() {
    server.SetupRoutes()

    fmt.Printf("Starting the meeting server...\n")
    http.ListenAndServe(":3000", nil)
}
