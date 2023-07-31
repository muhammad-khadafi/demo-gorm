package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"pustaka-api/handler"
	"pustaka-api/repository"
	"pustaka-api/service"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("DB Connection Error")
	}

	fmt.Println("DB Connection Establish..")

	// * Gunakan auto migrate jika ingin create table berdasarkan entity
	//db.AutoMigrate(&entity.Book{})

	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()
	api := router.Group("/api")
	api.POST("/books", bookHandler.InsertBook)
	api.PUT("/books", bookHandler.UpdateBook)
	api.GET("/books", bookHandler.GetAllBook)
	api.GET("/books/:id", bookHandler.GetBookById)
	api.GET("/cheap-good-books", bookHandler.GetCheapGoodBook)
	api.DELETE("/books/:id", bookHandler.DeleteBook)
	router.Run(":8888")
}
