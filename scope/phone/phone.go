package phone_pkg

import "fmt"

type Phone string

//create newphone object to access the methods

func Newphone(category string) Phone {
	return Phone(category)

}

func (p Phone) AddToCart() {
	fmt.Println("Added to Cart", p)

}

func (p Phone) Wishlist() {
	fmt.Println("wishlist added ", p)

}

func (p Phone) BuyNow() {
	fmt.Println("buy through cart", p)

}
