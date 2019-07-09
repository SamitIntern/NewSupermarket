package main

import (
	"testing"
)


func TestRepoCreateRepo(t *testing.T) {
	inventory :=   InitializeInventory()
	if len(inventory) != 4 {
		t.Errorf("RepoCreateRepo was incorrect, got: %d elements, wanted: %d.", len(inventory), 4)
	}
}

func TestCheckForExistingProducePositive(t *testing.T) {
	var product Product
	inventory = InitializeInventory()
	product = inventory[0]
	product = CheckForExistingProduce(product)
	if product.Id != 1{
		t.Errorf("CheckForExistingProduce was incorrect, got: %d as ID, wanted: %d.", product.Id, 1)
	}else{
		t.Log("CheckForExistingProduce was correct, got: 1 as ID, wanted: 1.")
	}
}


func TestCheckForExistingProduceNegative(t *testing.T) {
	var product Product

	inventory = InitializeInventory()
	product.Id = 7
	product.ProduceCode = "1111"
	product.Name = "Nothing"
	product.UnitPrice = 100
	product = CheckForExistingProduce(product)

	if product.Id != 0{
		t.Errorf("Negative tests for CheckForExistingProduce was incorrect, got: %d as ID, wanted: %d.", product.Id, 0)
	}else{
		t.Log("Negative tests for CheckForExistingProduce was correct, got: 0 as ID, wanted: 0.")
	}
}


func TestCheckForNonExistingProduceNegative(t *testing.T) {
	var product Product

	inventory = InitializeInventory()
	product.Id = CheckForNonExistingProduce(7)

	if product.Id != 0{
		t.Errorf("Negative tests for CheckForNonExistingProduce was incorrect, got: %d as ID, wanted: %d.", product.Id, 0)
	}else{
		t.Log("Negative tests for CheckForNonExistingProduce was correct, got: 0 as ID, wanted: 0.")
	}
}

func TestCheckForNonExistingProducePositive(t *testing.T) {
	var product Product

	inventory = InitializeInventory()
	product.Id = CheckForNonExistingProduce(2)

	if product.Id != 2{
		t.Errorf("Positive tests for CheckForNonExistingProduce was incorrect, got: %d as ID, wanted: %d.", product.Id, 0)
	}else{
		t.Log("Positive tests for CheckForNonExistingProduce was correct, got: 2 as ID, wanted: 2.")
	}
}

func TestRepoFindItemPositive(t *testing.T) {
	var product Product

	inventory = InitializeInventory()
	product = FindItem(2)

	if product.Id != 2{
		t.Errorf("Positive tests for RepoFindItem was incorrect, got: %d as ID, wanted: %d.", product.Id, 0)
	}else{
		t.Log("Positive tests for RepoFindItem was correct, got: 2 as ID, wanted: 2.")
	}
}

func TestRepoFindItemNegative(t *testing.T) {
	var product Product

	inventory = InitializeInventory()
	product = FindItem(7)

	if product.Id != 0{
		t.Errorf("Negative tests for RepoFindItem was incorrect, got: %d as ID, wanted: %d.", product.Id, 0)
	}else{
		t.Log("Negative tests for RepoFindItem was correct, got: 0 as ID, wanted: 0.")
	}
}

func TestRepoCreateItemPositive(t *testing.T) {
	var product Product
	var tempId = currentId
	inventory = InitializeInventory()
	product.ProduceCode = "1111"
	product.Name = "Nothing"
	product.UnitPrice = 100
	product = CreateItem(product)
	if product.Id != (tempId+1){
		t.Errorf("Positive tests for RepoCreateItem was incorrect, got: %d as ID, wanted: %d.", product.Id, tempId+1)
	}else{
		t.Log("Positive tests for RepoCreateItem was correct.")
	}
}

func TestRepoCreateItemNegative(t *testing.T) {
	var product Product
	var tempId = currentId
	inventory = InitializeInventory()
	product.ProduceCode = "1111"
	product.Name = "Nothing"
	product.UnitPrice = 100
	product = CreateItem(product)
	if product.Id == 0{
		t.Errorf("Positive tests for RepoCreateItem was incorrect, got: %d as ID, wanted: %d.",0, tempId+1)
	}else{
		t.Log("Positive tests for RepoCreateItem was correct.")
	}
}

func TestRepoDeleteItemNegative(t *testing.T) {

	var contains bool
	contains = false
	inventory = InitializeInventory()
	inventory = DeleteItem(1)

	for _, t2 := range inventory {
		t.Log(t2.Id)
		if t2.Id == 1{
			contains = true
		}
	}
	if contains{
		t.Errorf("Negative tests for RepoDeleteItem was incorrect, got produce with %d as ID, wanted it to be deleted.", 1)
	}else{
		t.Log("Negative tests for RepoDeleteItem was correct, got produce with 1 as ID deleted.")
	}
}

func TestRepoDeleteItemPositive(t *testing.T) {

	var count int

	inventory = InitializeInventory()
	count = len(inventory)
	inventory = DeleteItem(1)

	if count == len(inventory){
		t.Errorf("Positive tests for RepoDeleteItem was incorrect, the element was not deleted")
	}else{
		t.Log("Positive tests for RepoDeleteItem was correct, got produce deleted.")
	}
}

func TestGetAllItems(t *testing.T) {
	inventory = InitializeInventory()
	if len(inventory) != 4{
		t.Errorf("Test for GetAllItems was incorrect, did not get the correct number of elements")
	}else{
		t.Log("Test for GetAllItems was correct, got the correct number of elements.")
	}
}

func TestCheckValidAddParamNumberNamesNegative(t *testing.T) {
	var product Product
	var valid bool
	product.ProduceCode = "A12T-4GH7-QPL9-3N4M"
	product.Id = 1
	product.UnitPrice = 2
	valid = CheckValidAddParamNumber(product)
	if valid{
		t.Errorf("Test for CheckValidAddParamNumber was incorrect, expected it to fail but it passed")
	}else{
		t.Log("Test for CheckValidAddParamNumber was correct, expected it to fail and it failed")
	}
}

func TestCheckValidAddParamNumberNamesPositive(t *testing.T) {
	var product Product
	var valid bool
	product.ProduceCode = "A12T-4GH7-QPL9-3N4M"
	product.Id = 1
	product.UnitPrice = 2
	product.Name = "TestCase"
	valid = CheckValidAddParamNumber(product)
	if !valid{
		t.Errorf("Test for CheckValidAddParamNumber was incorrect, expected it to pass but it failed")
	}else{
		t.Log("Test for CheckValidAddParamNumber was correct, expected it to pass and it passed")
	}
}

func TestCheckValidAddParamNumberProduceCodePositive(t *testing.T) {
	var product Product
	var valid bool
	product.ProduceCode = "A12T-4GH7-QPL9-3N4M"
	product.Id = 1
	product.UnitPrice = 2
	product.Name = "TestCase"
	valid = CheckValidAddParamNumber(product)
	if !valid{
		t.Errorf("Test for CheckValidAddParamNumber was incorrect, expected it to pass but it failed")
	}else{
		t.Log("Test for CheckValidAddParamNumber was correct, expected it to pass and it passed")
	}
}

func TestCheckValidAddParamNumberProduceCodeNegative(t *testing.T) {
	var product Product
	var valid bool
	product.Id = 1
	product.UnitPrice = 2
	product.Name = "TestCase"
	valid = CheckValidAddParamNumber(product)
	if valid{
		t.Errorf("Test for CheckValidAddParamNumber was incorrect, expected it to fail but it passed")
	}else{
		t.Log("Test for CheckValidAddParamNumber was correct, expected it to fail and it failed")
	}
}

func TestCheckValidAddParamTypesCodeLengthPositive(t *testing.T) {
	var product Product
	var valid bool
	product.ProduceCode = "A12T-4GH7-QPL9-3N4M"
	product.Id = 1
	product.UnitPrice = 2
	product.Name = "TestCase"
	valid = CheckValidAddParamNumber(product)
	if !valid{
		t.Errorf("Test for CheckValidAddParamTypes was incorrect, expected it to pass but it failed")
	}else{
		t.Log("Test for CheckValidAddParamTypes was correct, expected it to pass and it passed")
	}
}

func TestCheckValidAddParamTypesCodeLengthNegative(t *testing.T) {
	var product Product
	var valid bool
	t.Log(valid)
	product.ProduceCode = "A12T-4GH7-QPL9"
	product.Name = "TestCase"
	product.Id = 1
	product.UnitPrice = 2
	valid = CheckValidAddParamTypes(product)
	t.Log(valid)
	if valid{
		t.Errorf("Test for CheckValidAddParamTypes was incorrect, expected it to fail but it passed")
	}else{
		t.Log("Test for CheckValidAddParamTypes was correct, expected it to fail and it failed")
	}
}

func TestCheckValidAddParamTypesCodeStructurePositive(t *testing.T) {
	var product Product
	var valid bool
	product.ProduceCode = "A12T-4GH7-QPL9-3N4M"
	product.Name = "TestCase"
	product.Id = 1
	product.UnitPrice = 2
	valid = CheckValidAddParamTypes(product)
	if !valid{
		t.Errorf("Test for CheckValidAddParamTypes was incorrect, expected it to pass but it failed")
	}else{
		t.Log("Test for CheckValidAddParamTypes was correct, expected it to pass and it passed")
	}
}

func TestCheckValidAddParamTypesCodeStructureNegative(t *testing.T) {
	var product Product
	var valid bool
	product.ProduceCode = "A12T-4GH7-Q-L9"
	product.Name = "TestCase"
	product.Id = 1
	product.UnitPrice = 2
	valid = CheckValidAddParamTypes(product)
	if valid{
		t.Errorf("Test for CheckValidAddParamTypes was incorrect, expected it to fail but it passed")
	}else{
		t.Log("Test for CheckValidAddParamTypes was correct, expected it to fail and it failed")
	}
}

/*func main(){

	var t *testing.T
	TestRepoCreateRepo(t)
	TestCheckForExistingProducePositive(t)
	TestCheckForExistingProduceNegative(t)
	TestCheckForNonExistingProduceNegative(t)
	TestCheckForNonExistingProducePositive(t)



}*/