package main

import (
	"context"

	"github.com/go-kit/kit/log"
)

type Book struct {
	BookId     string   `json:"bookId,omitempty"`
	Title      string   `json:"title,omitempty"`
	Edition    string   `json:"edition,omitempty"`
	Copyright  string   `json:"copyright,omitempty"`
	Language   string   `json:"language,omitempty"`
	Pages      string   `json:"pages,omitempty"`
	Authors    []string `json:"authors,omitempty"`
	Publishers []string `json:"publishers,omitempty"`
}

type bookservice struct {
	logger log.Logger
}

// Service describes the Book service.
type BookService interface {
	CreateBook(ctx context.Context, book Book) (string, error)
	GetBookById(ctx context.Context, id string) (interface{}, error)
	UpdateBook(ctx context.Context, book Book) (string, error)
	DeleteBook(ctx context.Context, id string) (string, error)
	GetBookIdPublishers(ctx context.Context, id string) ([]Publisher, error)
	GetBookIdAuthors(ctx context.Context, id string) ([]Author, error)
}

var books = []Book{
	Book{BookId: "Book1", Title: "Operating System Concepts", Edition: "9th",
		Copyright: "2012", Language: "ENGLISH", Pages: "976",
		Authors: []string{"author1"}, Publishers: []string{"publisher1"}},
	Book{BookId: "Book2", Title: "Computer Networks", Edition: "5th",
		Copyright: "2010", Language: "ENGLISH", Pages: "960",
		Authors: []string{"author2"}, Publishers: []string{"publisher2"}},
}

func find(x string) int {
	for i, book := range books {
		if x == book.BookId {
			return i
		}
	}
	return -1
}

func NewService(logger log.Logger) BookService {
	return &bookservice{
		logger: logger,
	}
}

func (s bookservice) CreateBook(ctx context.Context, book Book) (string, error) {
	var msg = "success"
	books = append(books, book)
	return msg, nil
}

func (s bookservice) GetBookById(ctx context.Context, id string) (interface{}, error) {
	var err error
	var book interface{}
	var empty interface{}
	i := find(id)
	if i == -1 {
		return empty, err
	}
	book = books[i]
	return book, nil
}
func (s bookservice) DeleteBook(ctx context.Context, id string) (string, error) {
	var err error
	msg := ""
	i := find(id)
	if i == -1 {
		return "", err
	}
	copy(books[i:], books[i+1:])
	books[len(books)-1] = Book{}
	books = books[:len(books)-1]
	return msg, nil
}
func (s bookservice) UpdateBook(ctx context.Context, book Book) (string, error) {
	var empty = ""
	var err error
	var msg = "success"
	i := find(book.BookId)
	if i == -1 {
		return empty, err
	}
	books[i] = book
	return msg, nil
}
func (s bookservice) GetBookIdPublishers(ctx context.Context, id string) ([]Publisher, error) {
	var err error
	var bookPublishers []Publisher
	i := find(id)
	if i == -1 {
		return bookPublishers, err
	}
	var book = books[find(id)]
	for _, publisher := range book.Publishers {
		bookPublishers = append(bookPublishers, publishers[findPublisher(publisher)])
	}
	return bookPublishers, nil
}

func (s bookservice) GetBookIdAuthors(ctx context.Context, id string) ([]Author, error) {
	var err error
	var bookAuthors []Author
	i := find(id)
	if i == -1 {
		return bookAuthors, err
	}
	var book = books[find(id)]
	for _, author := range book.Authors {
		bookAuthors = append(bookAuthors, authors[findAuthor(author)])
	}
	return bookAuthors, nil
}
