package main

import (
  "net/http"

  "http-server/api"
)

func main() {
  srv := api.NewServer()
  http.ListenAndServe(":8888", srv)
}
