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
		InitializeRepository,
	},
	Route{
		"ShowRepository",
		"GET",
		"/supermarket/{produceId}/",
		ShowRepository,
	},
	Route{
		"TodoNewProduce",
		"POST",
		"/supermarket/new/",
		AddNewProduct,
	},
	Route{
		"DeleteProduce",
		"DELETE",
		"/supermarket/delete/{produceId}/",
		DeleteProduct,
	},
	Route{
		"ShowAll",
		"GET",
		"/supermarket/all/showall/",
		GetRepository,
	},


}
