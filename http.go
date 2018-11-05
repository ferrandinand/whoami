package main

import (
  "os"
  "fmt"
  "net/http"
  "log"

  "github.com/golang/glog"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Fprintf(os.Stdout, "Listening on :%s\n", port)
    log.SetOutput(os.Stderr)
    hostname, _ := os.Hostname()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(os.Stdout, "I'm %s\n", hostname)
 	      fmt.Fprintf(w, "I'm a star %s\n", hostname)
        log.Println("Hello World")
        log.Println("I am not a star")
        log.Println("Myapp is running %s", hostname)
    })


    log.Fatal(http.ListenAndServe(":" + port, nil))
}


