package db

type Product struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Book struct {
	ID       *string `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Category *string `json:"category,omitempty"`
}

func ref[T any](val T) *T {
	return &val
}

var products = map[string]*Product{
	"ID-001": {
		ID:   ref("ID-001"),
		Name: ref("The Legend of Zelda"),
	},
	"ID-002": {
		ID:   ref("ID-002"),
		Name: ref("Mario Kart"),
	},
	"ID-003": {
		ID:   ref("ID-003"),
		Name: ref("Animal Crossing"),
	},
}

var books = map[string]*Book{
	"BOOK-001": {
		ID:       ref("BOOK-001"),
		Name:     ref("Thinking, Fast and Slow"),
		Category: ref("Non-Fiction"),
	},
	"BOOK-002": {
		ID:       ref("BOOK-002"),
		Name:     ref("The Power of Habit"),
		Category: ref("Non-Fiction"),
	},
	"BOOK-003": {
		ID:       ref("BOOK-003"),
		Name:     ref("The Lord of the Rings"),
		Category: ref("Fantasy"),
	},
	"BOOK-004": {
		ID:       ref("BOOK-004"),
		Name:     ref("A Song of Ice and Fire"),
		Category: ref("Fantasy"),
	},
}

var booksByCategory = map[string][]*Book{
	"Non-Fiction": {
		books["BOOK-001"],
		books["BOOK-002"],
	},
	"Fantasy": {
		books["BOOK-003"],
		books["BOOK-004"],
	},
}

func GetProductByID(ID string) *Product {
	p, ok := products[ID]
	if !ok {
		return nil
	}
	return p
}

func GetBooksByCategory(category string) []*Book {
	return booksByCategory[category]
}

func GetBooksByID(ID string) *Book {
	return books[ID]
}
