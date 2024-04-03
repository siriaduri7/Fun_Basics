package main

import (
	"fmt"
	phone_pkg "sri/phone"
	"sri/users"
)

// phones := []phone_pkg.Phone{p1, p2, p3}

func main() {

	// book1 := book.NewBook(1, "lupin_mystery", "lupin", 4, 3)
	// err := ram.AddToCart(&book1)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//create few phone objects

	sphone, err := phone_pkg.Newphone(1, "smartphone", 2)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("phone is", sphone)
	}

	//p1.AddToCart()
	fphone, err := phone_pkg.Newphone(2, "Featurephone", 3)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("phone is ", fphone)
	}
	lphone, err := phone_pkg.Newphone(3, "landlinephone", 4)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("phone is", lphone)
	}

	// phones := []phone_pkg.Phone{p1, p2, p3}
	// fmt.Println("all phone in cart", phones)

	//allow users to add phones to their own carts,Make sure user have there own carts
	//siri and ram and as many users as possible

	// siri := users.NewUser("siri") // siri shud add phone to cart
	// //fmt.Println(&siri) debug to check address
	// siri.AddToCart(fphone )
	// siri.GetItemInCart()
	ram := users.NewUser("Ram")

	err = ram.AddToCart(*sphone, 4)
	err = ram.AddToCart(*lphone, 3)
	// err := ram.AddToCart(fphone)
	// err = ram.AddToCart(lphone)
	if err != nil {
		fmt.Println(err)
	}
	// ram.GetItemInCart()
	// if err != nil {
	// 	fmt.Prinltn(err)
	// }
	err = ram.DeleteFromCart(5) // keep 2 valid data, 5 to catch error

	if err != nil {
		fmt.Println(err)
	}

	// siri := user_products.NewUser("siri")
	// // err = siri.AddToCart(sphone)
	// err = siri.AddToCart(book1)

	//ToDO delete from cart and it should not be in cart
	//and then deleted item from cart has to be added in wishlist & i print it shud show all wishlist

	// PhoneInCartByRam, err := addSmartPhoneToCart(phones, "smartphone")
	// if err != nil {

	// }

	// PhoneInCartBySiri, err := addSmartPhoneToCart(phones, "smartphone")
	// if err != nil {

	// }

	// 	buyPhoneInCart(PhoneInCartByRam)
	// 	buyPhoneInCart(PhoneInCartBySiri)
	// 	PhoneInCart1 := removeFromCart(phones)
	// 	buyPhoneInCart(PhoneInCart1)
	// 	listing1, removeditem := wishlistitems(phones, 1)
	// 	fmt.Println("added new listing of phones", listing1)
	// 	fmt.Println("added deleted item to wishlist", removeditem)

}

func addSmartPhoneToCart(phones []phone_pkg.Phone, phone phone_pkg.Phone) ([]phone_pkg.Phone, error) {
	// phone.AddToCart()
	return nil, nil

	//var PhoneInCart []phone_pkg.Phone

	// for _, phone := range phones { // this phone var is under phon_pkg/Phone struct thats why it could able to access struct functions

	// 	if phone == phoneCategory {
	// 		phone.AddToCart()

	// 		PhoneInCart = append(PhoneInCart, phone) // adding to array

	// 	}

	// }
	// return PhoneInCart, nil

}

// func buyPhoneInCart(phones []phone_pkg.Phone) {
// 	for _, phone := range phones {

// 		phone.BuyNow()

// 	}

// }

// HOMEWORK//////////////// removeFromCart
func removeFromCart(phones []phone_pkg.Phone) []phone_pkg.Phone {
	fmt.Println("**********")
	var PhoneInCart1 []phone_pkg.Phone
	for _, phone2 := range phones {
		//fmt.Println(phone, "((((((()))))))")
		if phone2 != "smartphone" {
			phone2.AddToCartAndRemoveFromInventory()
			PhoneInCart1 = append(PhoneInCart1, phone2)
		} else {
			fmt.Println("smartphone is sold out")

		}

	}

	return PhoneInCart1

	// PhoneInCart := addSmartPhoneToCart(phones)
	// fmt.Println(PhoneInCart, "(((((((*****)))))))")

	// if phone != "smartphone" {
	// 	fmt.Println("smart phone is soldout")
	// }

}

//phoneInCart := addSmartPhoneToCart(phones)
//fmt.Println(phoneInCart)

func wishlistitems(phones []phone_pkg.Phone, i int) ([]phone_pkg.Phone, phone_pkg.Phone) {
	var listing []phone_pkg.Phone
	var removed phone_pkg.Phone
	//deleting p2
	for index, phone3 := range phones {

		if index != i {
			listing = append(listing, phone3)

		} else {
			removed = phone3

		}
		// listing = append(listing, phone3[:index])
		// listing = append(listing, phone3[index+1:])

	}
	return listing, removed

}
