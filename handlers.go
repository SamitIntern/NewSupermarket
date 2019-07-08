package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

//This is like a welcome page for the supermarket
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Supermarket ! \n")
}


//This function is used when the user wants to check details of a particular produce
//This needs the Produce ID as a parameter

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todoId int
	var err error
	fmt.Println(vars)

	if todoId, err = strconv.Atoi(vars["produceId"]); err != nil {
		panic(err)
	}
	fmt.Println(err)
	todo := FindItem(todoId)
	if todo.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(todo); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Produce not found as per the Produce ID entered, please enter a valid Produce ID"}); err != nil {
		panic(err)
	}
}

// This function is called to initialize the repository
//It does not need any params
//Regardless of the number of adds or deletes called, this method would re-initialize the repository to have just the
// four objects with their respective produce codes and prices as mentioned in the ticket

func TodoCreate(w http.ResponseWriter, r *http.Request) {

	t := InitializeInventory()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

//This function is called when the user wants to add a new produce to the repository
//This function consumes a JSON which has 3 key-value pairs "Product Name", "Product Code" and "Price"
//The product name should be an alphanumeric string
//The product code should also be a 19 characters long alphanumeric string where every four characters are separated by
// a hyphen
//The price should be a float value

func TodoCreateFromJSON(w http.ResponseWriter, r *http.Request) {

	var todo Product
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	//if todo.Id == 0{
		if err := json.Unmarshal(body, &todo); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422)
			if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusUnprocessableEntity, Text: "Please provide a valid JSON"}); err != nil {
				panic(err)
			}
			return
		}
	//}

	t := AddItem(todo)
	if t.Id == 0{
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusBadRequest, Text: "Please provide valid entry of the product to be entered into the repository"}); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}


//This function is called when the user wants to delete a produce from the repository
//This function just takes the Produce ID as a param which has to be an integer

func TodoDelete(w http.ResponseWriter, r *http.Request) {

	var newtodos Inventory

	vars := mux.Vars(r)
	var produceId int
	var err error
	var todoslength int

	if produceId, err = strconv.Atoi(vars["produceId"]); err != nil {
		panic(err)
	}

	todoslength = len(inventory)

	newtodos = DeleteItem(produceId)

	if len(newtodos) == 0{
		if todoslength == 1{
			inventory = newtodos
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(inventory); err != nil {
				panic(err)
			}
			return
		}
		if todoslength > 1 {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusNotFound)
			if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "No such produce in the repository"}); err != nil {
				panic(err)
			}
			return
		}
	}
		inventory = newtodos
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(inventory); err != nil {
			panic(err)
		}
		return
}

//This function shows the current repository, after all the adds and deletes that may have performed.
//This function does not need any params

func TodoShowAll(w http.ResponseWriter, r *http.Request) {

	todos := GetAllItems()
	fmt.Print("Printing from handlers.go  ::   ")
	fmt.Print(todos)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
		return
}