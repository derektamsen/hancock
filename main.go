package main

import (
  "log"
  "net/http"
  "github.com/derektamsen/awss3urlsigner/server"
)

func main()  {
  http.HandleFunc("/", server.Handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
