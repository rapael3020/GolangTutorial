package main

import (
	"net/http"

	"github.com/rapael3020/WEB5/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
