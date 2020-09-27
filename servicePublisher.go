package main

import (
	"context"

	"github.com/go-kit/kit/log"
)

type Publisher struct {
	PublisherId string   `json:"publisherId,omitempty"`
	Name        string   `json:"name,omitempty"`
	Country     string   `json:"country,omitempty"`
	Founded     string   `json:"founded,omitempty"`
	Genere      string   `json:"genere,omitempty"`
	Books       []string `json:"books,omitempty"`
}

type publisherservice struct {
	logger log.Logger
}

// Service describes the Publisher service.
type PublisherService interface {
	CreatePublisher(ctx context.Context, publisher Publisher) (string, error)
	GetPublisherById(ctx context.Context, id string) (interface{}, error)
	UpdatePublisher(ctx context.Context, publisher Publisher) (string, error)
	DeletePublisher(ctx context.Context, id string) (string, error)
	GetPublisherIdBooks(ctx context.Context, id string) ([]Book, error)
}

var publishers = []Publisher{
	Publisher{PublisherId: "publisher1", Name: "test", Country: "test", Founded: "test", Genere: "test", Books: []string{"Book1", "Book2"}},
	Publisher{PublisherId: "publisher2", Name: "test", Country: "test", Founded: "test", Genere: "test", Books: []string{"Book1", "Book2"}},
}

func findPublisher(x string) int {
	for i, publisher := range publishers {
		if x == publisher.PublisherId {
			return i
		}
	}
	return -1
}

func NewServicePublisher(logger log.Logger) PublisherService {
	return &publisherservice{
		logger: logger,
	}
}

func (s publisherservice) CreatePublisher(ctx context.Context, publisher Publisher) (string, error) {
	var msg = "success"
	publishers = append(publishers, publisher)
	return msg, nil
}

func (s publisherservice) GetPublisherById(ctx context.Context, id string) (interface{}, error) {
	var err error
	var publisher interface{}
	var empty interface{}
	i := findPublisher(id)
	if i == -1 {
		return empty, err
	}
	publisher = publishers[i]
	return publisher, nil
}
func (s publisherservice) DeletePublisher(ctx context.Context, id string) (string, error) {
	var err error
	msg := ""
	i := findPublisher(id)
	if i == -1 {
		return "", err
	}
	copy(publishers[i:], publishers[i+1:])
	publishers[len(publishers)-1] = Publisher{}
	publishers = publishers[:len(publishers)-1]
	return msg, nil
}
func (s publisherservice) UpdatePublisher(ctx context.Context, publisher Publisher) (string, error) {
	var empty = ""
	var err error
	var msg = "success"
	i := findPublisher(publisher.PublisherId)
	if i == -1 {
		return empty, err
	}
	publishers[i] = publisher
	return msg, nil
}
func (s publisherservice) GetPublisherIdBooks(ctx context.Context, id string) ([]Book, error) {
	var err error
	var publisherBooks []Book
	i := findPublisher(id)
	if i == -1 {
		return publisherBooks, err
	}
	var publisher = publishers[findPublisher(id)]
	for _, book := range publisher.Books {
		publisherBooks = append(publisherBooks, books[find(book)])
	}
	return publisherBooks, nil
}
