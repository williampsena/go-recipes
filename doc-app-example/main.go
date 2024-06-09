// This package represents the application command for starting a web server.
package main

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/williampsena/go-recipes/doc-app-example/web"
)

// This function is responsible for setting up the program before it runs
func init() {
	gofakeit.Seed(0)
}

// Application entrypoint
func main() {
	svr := web.BuildServer()
	web.ListenAndServe(svr)
}
