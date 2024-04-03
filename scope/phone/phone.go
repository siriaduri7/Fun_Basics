package phone_pkg

import (
	"errors"
	"fmt"
)

type Phone struct { // to add properties of phone
	id               int
	name             string
	Stock            int  //fixed num total //2
	instock          *int // current based on adding phonincart in user it will change //firs2 //after -- // become1
	TechnicalDetails TechnicalDetails
}
type TechnicalDetails struct {
	productDescription string
	specialFeature     string
	itemWeight         string
	color              string
}

// create newphone object to access the methods
var usedIDS []int

func Newphone(id int, category string, stock int) (*Phone, error) {

	if IdUsed(id) {
		return nil, errors.New("id is already ")
	}
	usedIDS = append(usedIDS, id)

	inStockValue := stock

	phone := Phone{
		name:    category,
		id:      id, // assigning these values to phone properties phone sambandinchina struct ki adding
		Stock:   stock,
		instock: &inStockValue,
	}

	return &phone, nil

}

// stock := phone.stock

func IdUsed(id int) bool {
	for _, userid := range usedIDS {
		if userid == id {
			return true
		}

	}
	return false
}

func (p *Phone) IsInStock() bool { // true antey add // false avtey print outof stock
	fmt.Println("checking if phones are in stock", "stock:", p.stock, "instock:", *p.instock)

	if *p.instock >= p.stock && *p.instock > 0 { //instock 2 = stock 2

		return true
	}
	return false

}

func (p *Phone) Stock1() int {
	return p.stock

}

func (p *Phone) GetName() string {

	return p.name //name:  category, id:    id, stock: stock, to avoid exploding this to data to someone else we are restrictinga nd using via functions // if i am returning name not other data means I am securing my data

}

func (p *Phone) AddToCartAndRemoveFromInventory() {

	*p.instock-- //so with this added phone in user phonetocart and removing from instock first it was 2 now 1
	//2 became 1
	//values not updating in stock so we kept pointer remember this
}

// func (p *Phone) BuyNow() string {

// 	return p.name
// }
