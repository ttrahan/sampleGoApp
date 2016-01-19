package routes

import (
	"net/http"
	"github.com/ttrahan/sampleGoApp/hello"
)

func init() {
	http.HandleFunc("/", hello.Handler)
}
