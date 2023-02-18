package main

import (
	"Crud/controller"
	"Crud/service"
	"Crud/store"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type environmentVariables struct {
	Url string `required:"true"`
}

func loadEnv() environmentVariables {
	var env environmentVariables
	if err := envconfig.Process("", &env); err != nil {
		log.Fatal(err)
	}
	return env
}
func main() {
	env := loadEnv()
	db := store.Init(env.Url)
	bookStore := store.NewBookStore(db)
	bookService := service.NewBookService(bookStore)
	bookCntrl := controller.NewBookController(bookService)
	router := gin.Default()

	grpV1 := router.Group("/v1")

	grpV1.GET("/all/books", bookCntrl.GetAllBook)
	grpV1.GET("/books/author/:author", bookCntrl.GetAuthorsBooks)
	grpV1.GET("/book/title/:title", bookCntrl.GetBookWithTitle)
	grpV1.GET("/book/id/:id", bookCntrl.GetBookWithID)
	grpV1.DELETE("/delete/:bookID", bookCntrl.DeleteBook)
	grpV1.POST("/new/book", bookCntrl.CreateBook)
	grpV1.PATCH("/update/book", bookCntrl.UpdateBook)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
