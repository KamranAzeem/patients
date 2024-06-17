package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", MainServer)

    server_port := ":4444"

    http.ListenAndServe(server_port, nil)
}

func MainServer(w http.ResponseWriter, r *http.Request) {
    // The following lines are used to log to console, not on the web UI
    fmt.Println("Secret(s) and config files(s) were loaded.")
    server_port:=":4444"
    fmt.Printf("Starting application at port: %s", server_port)
    // Following is the main application functionality,
    //   which displays something on the web UI.
    // Keep changing the name of city to mimic change in the code.
    fmt.Fprintf(w, os.Getenv("GREETING") + " Patients! We are located in: Bodo")

}


