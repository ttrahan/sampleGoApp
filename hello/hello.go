package hello

import (
    "fmt"
    "github.com/ttrahan/sampleGoApp"
)

// func init() {
//     http.HandleFunc("/", Handler)
// }

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}
