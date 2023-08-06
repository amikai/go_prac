package db

type Product struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Book struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Category string `json:"category,omitempty"`
}

var GetProductByID = getProductByID
var GetBooksByCategory = getBooksByCategory
var GetBooksByID = getBooksByID

var products = map[string]*Product{
	"ID-001": {
		ID:   "ID-001",
		Name: "The Legend of Zelda",
	},
	"ID-002": {
		ID:   "ID-002",
		Name: "Mario Kart",
	},
	"ID-003": {
		ID:   "ID-003",
		Name: "Animal Crossing",
	},
}

var books = map[string]*Book{
	"BOOK-001": {
		ID:       "BOOK-001",
		Name:     "Thinking, Fast and Slow",
		Category: "Non-Fiction",
	},
	"BOOK-002": {
		ID:       "BOOK-002",
		Name:     "The Power of Habit",
		Category: "Non-Fiction",
	},
	"BOOK-003": {
		ID:       "BOOK-003",
		Name:     "The Lord of the Rings",
		Category: "Fantasy",
	},
	"BOOK-004": {
		ID:       "BOOK-004",
		Name:     "A Song of Ice and Fire",
		Category: "Fantasy",
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

func getProductByID(ID string) *Product {
	p, ok := products[ID]
	if !ok {
		return nil
	}
	return p
}

func getBooksByCategory(category string) []*Book {
	return booksByCategory[category]
}

func getBooksByID(ID string) *Book {
	return books[ID]
}
