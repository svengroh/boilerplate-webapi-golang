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
		"ModelIndex",
		"GET",
		"/public/models",
		ModelIndex,
	},
	Route{
		"ModelDetail",
		"GET",
		"/public/models/{modelId}",
		ModelDetail,
	},
	Route{
    	"ModelCreate",
    	"POST",
    	"/public/models",
    	ModelCreate,
	},
}
