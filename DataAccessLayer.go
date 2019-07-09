package main

import (
	"fmt"
	"strings"
	"reflect"
	"regexp"
)

var currentId int

var inventory Inventory

//This function contributes to the functionality of addition of a new produce to the repository
//This function checks the repository if a produce with the same produce code is present in the repository
//This function is called by the function "AddItem" of this class

func CheckForExistingProduce(product Product) Product {
	for _, products := range inventory {
		if strings.EqualFold(product.ProduceCode, products.ProduceCode) {
			return products
		}
	}
	return Product{}
}

//This function contributes to the functionality of deletion of a produce from the repository
//This function checks the repository if a produce with the same produce ID is present in the repository
//This function is called by the function "DeleteItem" of this class

func CheckForNonExistingProduce(produceId int) int {
	var index = 0
	for _, currentProduct := range inventory {
		index++
		if currentProduct.Id == produceId {
			return index
		}
	}
	return 0
}

//This function contributes to the functionality of searching for a produce in the repository
//This function checks the repository if a produce with the same produce ID is present in the repository
//This function is called by the function "ShowRepository" of the "handlers.go" class

func FindItem(id int) Product {
	for _, currentProduct := range inventory {
		if currentProduct.Id == id {
			return currentProduct
		}
	}
	return Product{}
}

//This function contributes to the functionality of creating the repository
//This function is called by the function "InitializeRepository" of the "handlers.go" class

func InitializeInventory() Inventory {

	var newInventory Inventory

	var newProduct Product
	currentId = 0
	newProduct = CreateItem(Product{Name: "Lettuce", ProduceCode:"A12T-4GH7-QPL9-3N4M", UnitPrice:3.46 })
	newInventory = append(newInventory, newProduct)
	newProduct = CreateItem(Product{Name: "peach", ProduceCode:"E5T6-9UI3-TH15-QR88", UnitPrice:2.99})
	newInventory = append(newInventory, newProduct)
	newProduct = CreateItem(Product{Name: "Green Pepper", ProduceCode:"YRT6-72AS-K736-L4AR", UnitPrice:0.79})
	newInventory = append(newInventory, newProduct)
	newProduct = CreateItem(Product{Name: "Gala Apple", ProduceCode:"TQ4C-VV6T-75ZX-1RMR", UnitPrice:3.59})
	newInventory = append(newInventory, newProduct)
	inventory = newInventory

	return inventory
}

//This function contributes to the functionality of addition of a new produce to the repository and creation of repository
//This function is called by the functions "AddItem" and "InitializeInventory" of this class

func CreateItem(currentProduct Product) Product {

	currentId += 1
	currentProduct.Id = currentId
	inventory = append(inventory, currentProduct)
	return currentProduct
}

//This function contributes to the functionality of addition of a new produce to the repository
//This function is called by the function "AddNewProduct" of "handlers.go" class

func AddItem(currentProduct Product) Product {

	var oldProduct = CheckForExistingProduce(currentProduct)
	if oldProduct.Id > 0 {
		return Product{}
	}

	if !CheckValidAddParamTypes(currentProduct){
		return Product{}
	}

	if !CheckValidAddParamNumber(currentProduct){
		return Product{}
	}

	/*currentId += 1
	currentProduct.Id = currentId
	inventory = append(inventory, currentProduct)*/
	currentProduct = CreateItem(currentProduct)
	return currentProduct
}

//This function contributes to the functionality of removal of a produce from the repository
//This function is called by the function "DeleteProduct" of "handlers.go" class

func DeleteItem( produceId int) Inventory {

	var index = CheckForNonExistingProduce(produceId)
	if index == 0 {
		return Inventory{}
	}

	/*var newInventory Inventory
	for _, currentProduct := range inventory {
		if currentProduct.Id == produceId{
			continue
		}else {
			newInventory = append(newInventory, currentProduct)
		}
	}
	inventory = newInventory
	return newInventory*/

	inventory[index] = inventory[len(inventory)-1]
	inventory = inventory[:len(inventory)-1]

	return inventory
}

//This function contributes to the functionality of removal of a produce from the repository
//This function is called by the function "GetRepository" of "handlers.go" class

func GetAllItems() Inventory {
	fmt.Print(inventory)
	return inventory
}

//This function contributes to the functionality of addition of a new produce to the repository
//This function validates the params before a new produce is added to the repository
//It checks if the "Produce Code" and "Produce Name" are strings
//It also validates if the "Produce Code" is 19 character long alphanumeric string with each 4 characters separated by
// a hyphen
//This function is called by the function "AddItem" of this class

func CheckValidAddParamTypes(currentProduct Product) bool{
	var valid bool
	valid = true
	if reflect.TypeOf(currentProduct.Name).String() != "string"{
		valid = false
		return valid
	}
	if reflect.TypeOf(currentProduct.ProduceCode).String() != "string"{
		valid = false
		return valid
	}
	if len(currentProduct.ProduceCode) != 19 {
		valid = false
		return valid
	}

	x := regexp.MustCompile(`-`)
	substrings := x.Split(currentProduct.ProduceCode, -1)
	if len(substrings) == 4{
		return valid
	}else{
		valid = false
		return valid
	}

	return valid
}

//This function contributes to the functionality of addition of a new produce to the repository
//This function validates the params before a new produce is added to the repository by making sure that neither the
// "Produce Name" or the "Produce Code" fields are empty

func CheckValidAddParamNumber(currentProduct Product) bool{
	var valid bool
	valid = true
	if currentProduct.Name == ""{
		valid = false
		return valid
	}
	if currentProduct.ProduceCode == ""{
		valid = false
		return valid
	}
	return valid
}
