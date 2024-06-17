package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
)

func main() {
    http.HandleFunc("/", DisplayMainMessage)

    server_port := ":4444"

    // The following lines are used to log to console, not on the web UI
    log.Println("Secret(s) and config files(s) were loaded.")

    log.Printf("Starting application at port: %s \n", server_port)

    http.ListenAndServe(server_port, nil)

}

func DisplayMainMessage(w http.ResponseWriter, r *http.Request) {
    // Following is the main application functionality,
    //   which displays something on the web UI.
    // Keep changing the name of city to mimic change in the code.
    fmt.Fprintf(w, os.Getenv("GREETING") + " Patients! We are located in: Bodo" + "\n")

}


