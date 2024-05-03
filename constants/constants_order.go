package constants

type OrderStatus int
type OrderCategory int

type OrderCategories []OrderCategory

const (
	Available OrderStatus = iota
	Buyed
)

const (
	Electronics OrderCategory = iota
	Fashion
	Home
	Sports
	Books
	Automotive
	Other
)
