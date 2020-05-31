package main

import (
	"net/http"

	"github.com/rapael3020/Golang-Restful-API/WEB5/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
