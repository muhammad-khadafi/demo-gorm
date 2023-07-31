package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pustaka-api/request"
	"pustaka-api/service"
	"pustaka-api/util"
	"strconv"
)

type bookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *bookHandler {
	return &bookHandler{bookService}
}

func (bookHandler *bookHandler) InsertBook(c *gin.Context) {
	var book request.BookRequest
	if err := c.ShouldBindJSON(&book); err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	err := bookHandler.bookService.Create(book)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Insert data success!", http.StatusOK, "ok", nil)
}

func (bookHandler *bookHandler) GetCheapGoodBook(c *gin.Context) {
	books, err := bookHandler.bookService.FindCheapAndGood()
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Retrieve data success!", http.StatusOK, "ok", books)
}

func (bookHandler *bookHandler) GetBookById(c *gin.Context) {
	idBuku, _ := strconv.Atoi(c.Param("id"))
	book, err := bookHandler.bookService.FindById(idBuku)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Retrieve data success!", http.StatusOK, "ok", book)
}

func (bookHandler *bookHandler) GetAllBook(c *gin.Context) {
	books, err := bookHandler.bookService.FindAll()
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Retrieve data success!", http.StatusOK, "ok", books)
}

func (bookHandler *bookHandler) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := bookHandler.bookService.Delete(id)
	if err != nil {
		util.APIResponse(c, "Failed to delete data", http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Delete data success!", http.StatusOK, "ok", nil)
}

func (bookHandler *bookHandler) UpdateBook(c *gin.Context) {
	var book request.BookRequest
	if err := c.ShouldBindJSON(&book); err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	err := bookHandler.bookService.Update(book)
	if err != nil {
		util.APIResponse(c, err.Error(), http.StatusBadRequest, "error", nil)
		return
	}
	util.APIResponse(c, "Update data success!", http.StatusOK, "ok", nil)
}
