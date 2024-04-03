package users

import (
	"errors"
	"fmt"

	phone_pkg "sri/phone"
)

type User struct {
	name string
	// phonesInCart []products.Services
	phonesInCart    []phone_pkg.Phone
	PhoneInWishlist []phone_pkg.Phone
}

//phonesInCart := []phone_pkg.Phone cannot do the shorthand outside of functions *****

// TODO: users need to ahve seperate cart
//var phonesInCart []phone_pkg.Phone // drawback as it is adding siri cart with ram cart displaying both carts because in "users" package globally declared it is not specific to one user // data is getting shared as it is globally in package // this need to be user specific only to user that is were we need structs
//to avaoid this we using struct that can be so personal to user

func NewUser(userName string) *User {
	return &User{name: userName}

}

func (u *User) AddToCart(phone phone_pkg.Phone, quantity int) error { // when you got phone add to cart phone = p1 // var product
	if phone.IsInStock() {
		if quantity > phone.Stock {
			return errors.New("requested quantity exceeds from stock")
		}
		u.phonesInCart = append(u.phonesInCart, phone)
		phone.AddToCartAndRemoveFromInventory()
		fmt.Println(u.name, "tried to add:", phone, "in to cart now items in cart", len(u.phonesInCart))
		return nil

	}
	return errors.New(fmt.Sprintf("phone: %s is out of stock", phone.GetName()))

	// func (u *User) AddToCart(product products.Services) error { // when you got phone add to cart phone = p1 // var product
	// 	if product.IsInStock() {
	// 		u.phonesInCart = append(u.phonesInCart, product)
	// 		product.AddToCartAndRemoveFromInventory()
	// 		fmt.Println(u.name, "tried to add:", product, "in to cart now items in cart", len(u.phonesInCart))
	// 		return nil

	// 	}
	// 	return errors.New(fmt.Sprintf("phone: %s is out of stock", product.GetName()))
	// new will take string so want to give meaningful string// to return string New accep string
	// fmt.Println("Added to Cart", phone)
	// u.phonesInCart = append(u.phonesInCart, phone)

}

//user to see i am using getcart, get is to see

func (u *User) GetItemInCart() { // when you got phone add to cart phone = p1
	//print all items
	fmt.Println("getItems in Cart", u.name, u.phonesInCart) // if you dont give u.phonesInCart it give error as undefined as it could not find wer the var is declared so need to give u.

}

func (u *User) Wishlist() {
	fmt.Println("wishlist added ", u)

}

func (u *User) AddToWishlist() {
	fmt.Println("wishlist added ", u)

}

func (u *User) DeleteFromCart(position int) error {
	if position < 0 || position >= len(u.phonesInCart) {
		return errors.New("invalid position")

	}
	for i, phone := range u.phonesInCart {
		if i == position {
			u.phonesInCart = append(u.phonesInCart[:i], u.phonesInCart[i+1:]...)
			deletedPhone := u.phonesInCart[position]
			u.PhoneInWishlist = append(u.PhoneInWishlist, deletedPhone)
			fmt.Println("deleted phones are added to wishlist", u.PhoneInWishlist)
			fmt.Println(u.name, " delete:1", phone.GetName(), "from cart, nowitems in cart", len(u.phonesInCart))

		}

	}
	return nil

}

//todo delete all items in cart specific name instead of  position

// fmt.Println("buy through cart", u)

func (u *User) DeleteFromWishlist() {
	fmt.Println("delete from wishlist")
}
