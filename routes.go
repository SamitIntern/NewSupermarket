package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes Routes = Routes{

	Route{
		"TodoIndex",
		"GET",
		"/supermarket/",
		Index,
	},
	Route{
		"TodoCreateRepo",
		"POST",
		"/supermarket/createrepo/",
		TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/supermarket/{produceId}/",
		TodoShow,
	},
	Route{
		"TodoNewProduce",
		"POST",
		"/supermarket/new/",
		TodoCreateFromJSON,
	},
	Route{
		"DeleteProduce",
		"DELETE",
		"/supermarket/delete/{produceId}/",
		TodoDelete,
	},
	Route{
		"ShowAll",
		"GET",
		"/supermarket/all/showall/",
		TodoShowAll,
	},


}
