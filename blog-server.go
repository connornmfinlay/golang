package main

import "net/http"

func main() {

type Post struct {
  ID string
  PostedAt string
  Title string
  Body string
}

http.HandleFunc("POST /posts",
  func(w http.ResponseWriter, r *http.Request) {
  }
)

http.HandleFunc("GET /posts/{id}",
  func(w http.ResponseWriter, r *http.Request) {

  }
)

http.Handle("/static/",
  http.StripPrefix("/static/",
    http.FileServer(http.Dir("static/"))))


var postTemplate = 
template.Must(template.New("").Parse('
<html>
<body>
<h1>{{.Title}}</h1>
  <h3>{{.PostedAt}}</h3>
  <p>{{.Body}}</p>
</body>
</html>
'))

http.ListenAndServeTLS(":8989", "cert.pem", "key.pem", nil)

}
