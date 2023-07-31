package repository

import (
	"gorm.io/gorm"
	"log"
	"pustaka-api/entity"
)

type BookRepository interface {
	FindAll() ([]entity.Book, error)
	FindCheapAndGood() ([]entity.Book, error)
	FindById(ID int) (entity.Book, error)
	Create(book entity.Book) error
	Update(book entity.Book) error
	Delete(book entity.Book) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) FindAll() ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) FindCheapAndGood() ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.Raw("SELECT * FROM BOOKS WHERE price < ? AND rating = ?", 100000, 5).Scan(&books).Error
	return books, err
}

func (r *bookRepository) FindById(ID int) (entity.Book, error) {
	var book entity.Book
	err := r.db.Find(&book, ID).Error
	return book, err
}

func (r *bookRepository) Create(book entity.Book) error {
	err := r.db.Create(&book).Error
	if err != nil {
		log.Fatal("Error insert data!")
	}
	return err
}

func (r *bookRepository) Delete(book entity.Book) error {
	err := r.db.Delete(&book).Error
	if err != nil {
		log.Fatal("Error delete data!")
	}
	return err
}

func (r *bookRepository) Update(book entity.Book) error {
	err := r.db.Save(&book).Error
	if err != nil {
		log.Fatal("Error update data!")
	}
	return err
}
