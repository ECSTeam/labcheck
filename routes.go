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
	Route{"Index", "GET", "/", Index},
	Route{"LabIndex", "GET", "/labindex", LabIndex},
	Route{"LabShow", "GET", "/lab/{labName}", LabShow},
	Route{"LabCreate", "POST", "/labcreate", LabCreate},
	Route{"LabCheck", "POST", "/labs", LabCheck},
}
