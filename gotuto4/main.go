package main

import (
	"log"
	"net/http"

	"github.com/rapael3020/WEB17/app"
	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
