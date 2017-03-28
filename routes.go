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
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "LabIndex",
        "GET",
        "/labs",
        LabIndex,
    },
    Route{
        "LabShow",
        "GET",
        "/labs/{labId}",
        LabShow,
    },
      Route{
        "LabCreate",
        "POST",
        "/labs",
        LabCreate,
    },
}
