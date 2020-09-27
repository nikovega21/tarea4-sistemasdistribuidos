package main

import (
	"context"

	"github.com/go-kit/kit/log"
)

type Author struct {
	AuthorId    string   `json:"authorId,omitempty"`
	Name        string   `json:"name,omitempty"`
	Nationality string   `json:"nationality,omitempty"`
	Birth       string   `json:"birth,omitempty"`
	Genere      string   `json:"genere,omitempty"`
	Books       []string `json:"books,omitempty"`
}

type authorservice struct {
	logger log.Logger
}

// Service describes the Author service.
type AuthorService interface {
	CreateAuthor(ctx context.Context, author Author) (string, error)
	GetAuthorById(ctx context.Context, id string) (interface{}, error)
	UpdateAuthor(ctx context.Context, author Author) (string, error)
	DeleteAuthor(ctx context.Context, id string) (string, error)
	GetAuthorIdBooks(ctx context.Context, id string) ([]Book, error)
}

var authors = []Author{
	Author{AuthorId: "author1", Name: "test", Nationality: "test", Birth: "test", Genere: "test", Books: []string{"Book1"}},
	Author{AuthorId: "author2", Name: "test", Nationality: "test", Birth: "test", Genere: "test", Books: []string{"Book1", "Book2"}},
}

func findAuthor(x string) int {
	for i, author := range authors {
		if x == author.AuthorId {
			return i
		}
	}
	return -1
}

func NewServiceAuthor(logger log.Logger) AuthorService {
	return &authorservice{
		logger: logger,
	}
}

func (s authorservice) CreateAuthor(ctx context.Context, author Author) (string, error) {
	var msg = "success"
	authors = append(authors, author)
	return msg, nil
}
func (s authorservice) DeleteAuthor(ctx context.Context, id string) (string, error) {
	var err error
	msg := ""
	i := findAuthor(id)
	if i == -1 {
		return "", err
	}
	copy(authors[i:], authors[i+1:])
	authors[len(authors)-1] = Author{}
	authors = authors[:len(authors)-1]
	return msg, nil
}
func (s authorservice) UpdateAuthor(ctx context.Context, author Author) (string, error) {
	var empty = ""
	var err error
	var msg = "success"
	i := findAuthor(author.AuthorId)
	if i == -1 {
		return empty, err
	}
	authors[i] = author
	return msg, nil
}
func (s authorservice) GetAuthorById(ctx context.Context, id string) (interface{}, error) {
	var err error
	var author interface{}
	var empty interface{}
	i := findAuthor(id)
	if i == -1 {
		return empty, err
	}
	author = authors[i]
	return author, nil
}
func (s authorservice) GetAuthorIdBooks(ctx context.Context, id string) ([]Book, error) {
	var err error
	var authorBooks []Book
	i := findAuthor(id)
	if i == -1 {
		return authorBooks, err
	}
	var author = authors[findAuthor(id)]
	for _, book := range author.Books {
		authorBooks = append(authorBooks, books[find(book)])
	}
	return authorBooks, nil
}
