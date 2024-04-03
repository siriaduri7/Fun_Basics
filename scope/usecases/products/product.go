package products

type Services interface {
	IsInStock() bool //phone functions common in book nd phone
	GetName() string
	AddToCartAndRemoveFromInventory()
}