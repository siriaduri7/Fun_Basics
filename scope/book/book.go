package book

import "fmt"

type Book struct {
	id         int
	name       string
	authorName string
	stock      int
	instock    *int
}

func NewBook(id int, name string, authorName string, stock int, instock *int) Book {
	instockvalue := stock //don't want to change the total value so making copy of stock and adding operation on it below
	book := Book{
		id:         id,
		name:       name,
		authorName: authorName,
		stock:      stock,
		instock:    &instockvalue, // pointing to stock if you change value in instock will effect stock
	} //By using a pointer, you can indirectly modify the value of stock through the instock pointer, ensuring consistency between the two.

	return book

}

func (b *Book) IsInStock() bool {
	fmt.Println("chekcing if book is in the stock", "stock:", b.stock, "instock:", *b.instock)
	//check if book in stock
	if *b.instock <= b.stock && *b.instock > 0 { // start instock is 0 before adding
		return true

	}

	return false

}

func (b *Book) GetName() string {
	return b.name

}


func (b *Book) AddToCartRemoveFromInventory() {
	*b.instock--
}

func (b *Book) GetProductDetails() *Book {
	return b

}
