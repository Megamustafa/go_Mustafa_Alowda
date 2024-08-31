package entities

type Wishlist struct {
	ID          string
	Name        string
	Description string
	Items       []Item
}

type Item struct {
	ID          string
	Name        string
	Description string
}
