package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//These are all the defined routes for every functionality of this API

type Routes []Route

var routes Routes = Routes{

	Route{
		"Index",
		"GET",
		"/supermarket/",
		Index,
	},
	Route{
		"CreateRepo",
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
		"NewProduce",
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
