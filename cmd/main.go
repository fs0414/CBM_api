package main

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello cbmApi to golang air")
  })

  http.ListenAndServe(":8080", nil)
}
