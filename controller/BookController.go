package controller

import (
	"Crud/models"
	"Crud/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookController struct {
	bookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return BookController{bookService: bookService}
}

func (b BookController) GetAllBook(ctx *gin.Context) {
	books, err := b.bookService.GetAllBooks()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (b BookController) GetAuthorsBooks(ctx *gin.Context) {
	var author models.GetAuthorsBooks
	if ctx.Param("author") == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err:": "empty path var"})
		return
	}
	author.Author = ctx.Param("author")
	books, err := b.bookService.GetAuthorsBooks(author)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (b BookController) GetBookWithTitle(ctx *gin.Context) {
	var title models.GetBookWithTitle
	if ctx.Param("title") == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err:": "empty path var"})
		return
	}
	title.Title = ctx.Param("title")
	books, err := b.bookService.GetBooksWithTitle(title)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (b BookController) GetBookWithID(ctx *gin.Context) {
	var bookID models.GetBook
	if ctx.Param("id") == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err:": "empty path var"})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err:": err.Error()})
		return
	}
	bookID.ID = id
	books, err := b.bookService.GetBookWithID(bookID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, books)
}

func (b BookController) DeleteBook(ctx *gin.Context) {
	var del models.DeleteRequest
	if ctx.Param("bookID") == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err:": "empty path var"})
		return
	}
	id, err := strconv.Atoi(ctx.Param("bookID"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err:": err.Error()})
		return
	}
	del.ID = id
	err = b.bookService.DeleteBook(del)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (b BookController) UpdateBook(ctx *gin.Context) {
	var req models.UpdateBook
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err:": err.Error()})
		return
	}
	err := b.bookService.UpdateBook(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}

func (b BookController) CreateBook(ctx *gin.Context) {
	var req models.SaveBook
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err:": err.Error()})
		return
	}
	err := b.bookService.CreateBook(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Status(http.StatusOK)
}
