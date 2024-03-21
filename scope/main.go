package main

import (
	"fmt"
	phone_pkg "sri/phone"
)

func main() {

	//create few phone objects
	p1 := phone_pkg.Newphone("smartphone")
	//p1.AddToCart()
	p2 := phone_pkg.Newphone("Featurephone")
	p3 := phone_pkg.Newphone("landlinephone")

	phones := []phone_pkg.Phone{p1, p2, p3}

	PhoneInCart := addSmartPhoneToCart(phones)
	buyPhoneInCart(PhoneInCart)
	PhoneInCart1 := removeFromCart(phones)
	buyPhoneInCart(PhoneInCart1)
	listing1 := wishlistitems(phones, 1)
	fmt.Println(listing1)

}

func addSmartPhoneToCart(phones []phone_pkg.Phone) []phone_pkg.Phone {
	var PhoneInCart []phone_pkg.Phone

	for _, phone := range phones { // this phone var is under phon_pkg/Phone struct thats why it could able to access struct functions

		if phone == "smartphone" {
			phone.AddToCart()

			PhoneInCart = append(PhoneInCart, phone) // adding to array

		}

	}
	return PhoneInCart

}

func buyPhoneInCart(phones []phone_pkg.Phone) {
	for _, phone := range phones {

		phone.BuyNow()

	}

}

func removeFromCart(phones []phone_pkg.Phone) []phone_pkg.Phone {
	fmt.Println("**********")
	var PhoneInCart1 []phone_pkg.Phone
	for _, phone2 := range phones {
		//fmt.Println(phone, "((((((()))))))")
		if phone2 != "smartphone" {
			phone2.AddToCart()
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

func wishlistitems(phones []phone_pkg.Phone, i int) []phone_pkg.Phone {
	var listing []phone_pkg.Phone
	//deleting p2
	for index, phone3 := range phones {

		if index != i {
			listing = append(listing, phone3)
		}
		// listing = append(listing, phone3[:index])
		// listing = append(listing, phone3[index+1:])
		phone3.Wishlist()

	}
	return listing

}
