package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/amikai/go_prac/httpex/db"
)

type Resp[T any] struct {
	Data  T      `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func ProductHandler(c *gin.Context) {
	ID := c.Param("id")
	resp := Resp[*db.Product]{
		Data: db.GetProductByID(ID),
	}
	c.JSON(http.StatusOK, resp)
}

func BooksHandler(c *gin.Context) {
	ID := c.Param("id")
	resp := Resp[*db.Book]{
		Data: db.GetBooksByID(ID),
	}
	c.JSON(http.StatusTeapot, resp)
}

func BooksCategoryHandler(c *gin.Context) {
	category := c.Param("category")
	resp := Resp[[]*db.Book]{
		Data: db.GetBooksByCategory(category),
	}
	c.JSON(http.StatusOK, resp)
}
