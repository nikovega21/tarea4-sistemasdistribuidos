package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

func makeCreateBookEndpoint(s BookService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateBookRequest)
		msg, err := s.CreateBook(ctx, req.book)
		return CreateBookResponse{Msg: msg, Err: err}, nil
	}
}
func makeGetBookByIdEndpoint(s BookService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBookByIdRequest)
		bookDetails, err := s.GetBookById(ctx, req.Id)
		if err != nil {
			return GetBookByIdResponse{Book: bookDetails, Err: "Id not found"}, nil
		}
		return GetBookByIdResponse{Book: bookDetails, Err: ""}, nil
	}
}
func makeDeleteBookEndpoint(s BookService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteBookRequest)
		msg, err := s.DeleteBook(ctx, req.Bookid)
		if err != nil {
			return DeleteBookResponse{Msg: msg, Err: err}, nil
		}
		return DeleteBookResponse{Msg: msg, Err: nil}, nil
	}
}
func makeUpdateBookendpoint(s BookService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateBookRequest)
		msg, err := s.UpdateBook(ctx, req.book)
		return msg, err
	}
}
func makeGetBookByIdPublishersEndpoint(s BookService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBookByIdPublishersRequest)
		publishers, err := s.GetBookIdPublishers(ctx, req.Id)
		if err != nil {
			return GetBookByIdPublishersResponse{Publishers: publishers, Err: "Id not found"}, nil
		}
		return GetBookByIdPublishersResponse{Publishers: publishers, Err: ""}, nil
	}
}
func makeGetBookByIdAuthorsEndpoint(s BookService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBookByIdAuthorsRequest)
		authors, err := s.GetBookIdAuthors(ctx, req.Id)
		if err != nil {
			return GetBookByIdAuthorsResponse{Authors: authors, Err: "Id not found"}, nil
		}
		return GetBookByIdAuthorsResponse{Authors: authors, Err: ""}, nil
	}
}

func decodeCreateBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateBookRequest
	fmt.Println("-------->>>>into Decoding")
	if err := json.NewDecoder(r.Body).Decode(&req.book); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetBookByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetBookByIdRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = GetBookByIdRequest{
		Id: vars["bookid"],
	}
	return req, nil
}
func decodeDeleteBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Delete Decoding")
	var req DeleteBookRequest
	vars := mux.Vars(r)
	req = DeleteBookRequest{
		Bookid: vars["bookid"],
	}
	return req, nil
}
func decodeUpdateBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Update Decoding")
	var req UpdateBookRequest
	if err := json.NewDecoder(r.Body).Decode(&req.book); err != nil {
		return nil, err
	}
	return req, nil
}
func decodeGetBookByIdPublishersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetBookByIdPublishersRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = GetBookByIdPublishersRequest{
		Id: vars["bookid"],
	}
	return req, nil
}
func decodeGetBookByIdAuthorsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetBookByIdAuthorsRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = GetBookByIdAuthorsRequest{
		Id: vars["bookid"],
	}
	return req, nil
}
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}

type (
	CreateBookRequest struct {
		book Book
	}
	CreateBookResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}
	GetBookByIdRequest struct {
		Id string `json:"bookid"`
	}
	GetBookByIdResponse struct {
		Book interface{} `json:"book,omitempty"`
		Err  string      `json:"error,omitempty"`
	}

	DeleteBookRequest struct {
		Bookid string `json:"bookid"`
	}

	DeleteBookResponse struct {
		Msg string `json:"response"`
		Err error  `json:"error,omitempty"`
	}
	UpdateBookRequest struct {
		book Book
	}
	UpdateBookResponse struct {
		Msg string `json:"status,omitempty"`
		Err error  `json:"error,omitempty"`
	}

	GetBookByIdPublishersRequest struct {
		Id string `json:"bookid"`
	}
	GetBookByIdPublishersResponse struct {
		Publishers []Publisher `json:"publishers",omitempty`
		Err        string      `json:"error,omitempty"`
	}

	GetBookByIdAuthorsRequest struct {
		Id string `json:"bookid"`
	}
	GetBookByIdAuthorsResponse struct {
		Authors []Author `json:"authors",omitempty`
		Err     string   `json:"error,omitempty"`
	}
)
