package service

import (
	"pustaka-api/entity"
	"pustaka-api/repository"
	"pustaka-api/request"
)

type BookService interface {
	FindAll() ([]entity.Book, error)
	FindById(ID int) (entity.Book, error)
	FindCheapAndGood() ([]entity.Book, error)
	Create(bookRequest request.BookRequest) error
	Update(bookRequest request.BookRequest) error
	Delete(ID int) error
}

type bookService struct {
	repository repository.BookRepository
}

func NewBookService(repository repository.BookRepository) *bookService {
	return &bookService{repository}
}

func (s *bookService) FindAll() ([]entity.Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	//return s.repository.FindAll()
}

func (s *bookService) FindCheapAndGood() ([]entity.Book, error) {
	books, err := s.repository.FindCheapAndGood()
	return books, err
	//return s.repository.FindAll()
}

func (s *bookService) FindById(ID int) (entity.Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *bookService) Create(bookRequest request.BookRequest) error {
	price, _ := bookRequest.Price.Int64()
	book := entity.Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       int(price),
	}
	err := s.repository.Create(book)
	return err
}

func (s *bookService) Update(bookRequest request.BookRequest) error {
	id, _ := bookRequest.ID.Int64()
	price, _ := bookRequest.Price.Int64()
	book := entity.Book{
		ID:          uint(id),
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		Price:       int(price),
	}
	err := s.repository.Update(book)
	return err
}

func (s *bookService) Delete(ID int) error {
	book := entity.Book{
		ID: uint(ID),
	}
	err := s.repository.Delete(book)
	return err
}
