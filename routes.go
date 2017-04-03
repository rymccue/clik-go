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
	/*
		GET /users/{id}
			- returns photos as well as user data
	*/
	Route{
		"UserGet",
		"GET",
		"/v1/users/{id}",
		UserGet,
	},

	/*
		POST /users
	*/
	Route{
		"UserCreate",
		"POST",
		"/v1/users",
		UserCreate,
	},

	/*
		POST /users/{id}/edit
		params:
			any of the editable fields in user table
	*/
	Route{
		"UserEdit",
		"POST",
		"/v1/users/{id}/edit",
		UserEdit,
	},

	/*
		GET /users/{id}/queue
			- pagination?
	*/
	Route{
		"UserQueue",
		"GET",
		"/v1/users/{id}/queue",
		UserGetQueue,
	},

	/*
		GET /users/{id}/matches
			- pagination?
	*/
	Route{
		"UserMatches",
		"GET",
		"/v1/users/{id}/matches",
		UserGetMatches,
	},

	/*
		POST /decisions/{from_user_id}
		params:
			decision: true/false
			user_id: the user being decided upon
	*/
	Route{
		"DecisionCreate",
		"POST",
		"/v1/decisions/{id}",
		DecisionCreate,
	},

	/*
		DEL /matches/{id}
	*/
	Route{
		"MatchDelete",
		"DEL",
		"/v1/matches/{id}",
		MatchDelete,
	},

	Route{
		"AccessToken",
		"POST",
		"/access-token",
		AccessToken,
	},
}
