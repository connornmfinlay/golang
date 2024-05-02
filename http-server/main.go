package main

import (
  "net/http"

  "github.com/connornmfinlay/golang/http-server"
)

func main() {
  srv := api.NewServer()
  http.ListenAndServe(":8888", srv)
}
