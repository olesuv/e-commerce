package constants

type OrderStatus int
type OrderCategory int

type OrderCategories []OrderCategory

type OrderCurrency int

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

const (
	UAH OrderCurrency = iota
	USD
	EUR
)
