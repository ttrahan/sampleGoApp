package hello

import (
    "fmt"
    "routes"
)

// func init() {
//     http.HandleFunc("/", Handler)
// }

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}
