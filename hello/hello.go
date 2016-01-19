package hello

import (
    "fmt"
    "net/http"
)

// func init() {
//     http.HandleFunc("/", Handler)
// }

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}
