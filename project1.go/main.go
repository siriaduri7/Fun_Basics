package main

func main() {

	action := chooseBasket("vegBasket")
	println(action)
	println(getGroceryFillInBaskets(action, "veggies"))

	b := chooseBasket("fruitBasket")
	println(b)
	println(getGroceryFillInBaskets(b, "fruits"))

}

func chooseBasket(Basket string) string {
	var response = Basket + " " + "is choosen and" //+ getGroceryFillInBaskets("vegBasket", "veggies")
	return response
}

func getGroceryFillInBaskets(baske1, items string) string {
	var response = " collected " + items + " in " + baske1
	return response
}
