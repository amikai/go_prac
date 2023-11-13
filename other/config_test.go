package other

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBook(t *testing.T) {
	config := &Config{
		ISBN:  "978-7-111-59999-8",
		Name:  "Hello",
		Pages: 10,
	}
	book := NewBook(config)
	assert.Equal(t, "978-7-111-59999-8", book.ISBN)
	assert.Equal(t, "Hello", book.Name)
	assert.Equal(t, 10, book.Pages)
}
