package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"index", "GET", "/", index},
	Route{"labsHandler", "GET", "/labs", labsHandler},
	Route{"loadData", "POST", "/load", loadData},
	Route{"labCheck", "POST", "/labs", labCheck},
}
