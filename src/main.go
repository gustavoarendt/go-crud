package main

import (
	"go-crud/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe("127.0.0.1:8080", nil)
}
