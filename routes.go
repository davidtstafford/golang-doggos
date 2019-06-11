package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"DoggoIndex",
		"GET",
		"/Doggos",
		DoggoIndex,
	},
	Route{
		"DoggoCreate",
		"POST",
		"/Doggos",
		AddDoggo,
	},	
	Route{
		"DoggoDelete",
		"DELETE",
		"/Doggos",
		DeleteDoggo,
	},/*
	Route{
		"DoggoShow",
		"GET",
		"/Doggos/{DoggoId}",
		DoggoShow,
	},*/
}
