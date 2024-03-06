package main

//import "fmt"

func main() {
	//==================================HW===============================================================
	// basket := vegBasket
	println(chooseBasket("vegBasket"))
	println(chooseBasket("meatBasket"))
	println(GroceryStore("siri", "veggies", "vegBasket"))
	println(GroceryStore("ram", "meat", "meatBasket"))

}

func chooseBasket(basket string) string {
	var response = basket + " " + "of vegetables are choosen"
	return response
}

func GroceryStore(name, items, basket string) string {
	var response = name + "collected " + items + " in  " + basket
	return response
}

//=================================================================================================
