package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"
)

func makeCreateAuthorEndpoint(s AuthorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAuthorRequest)
		msg, err := s.CreateAuthor(ctx, req.author)
		return CreateAuthorResponse{Msg: msg, Err: err}, nil
	}
}
func makeGetAuthorByIdEndpoint(s AuthorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAuthorByIdRequest)
		authorDetails, err := s.GetAuthorById(ctx, req.Id)
		if err != nil {
			return GetAuthorByIdResponse{Author: authorDetails, Err: "Id not found"}, nil
		}
		return GetAuthorByIdResponse{Author: authorDetails, Err: ""}, nil
	}
}
func makeGetAuthorByIdBooksEndpoint(s AuthorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAuthorByIdBooksRequest)
		books, err := s.GetAuthorIdBooks(ctx, req.Id)
		if err != nil {
			return GetAuthorByIdBooksResponse{Books: books, Err: "Id not found"}, nil
		}
		return GetAuthorByIdBooksResponse{Books: books, Err: ""}, nil
	}
}
func makeDeleteAuthorEndpoint(s AuthorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteAuthorRequest)
		msg, err := s.DeleteAuthor(ctx, req.Authorid)
		if err != nil {
			return DeleteAuthorResponse{Msg: msg, Err: err}, nil
		}
		return DeleteAuthorResponse{Msg: msg, Err: nil}, nil
	}
}
func makeUpdateAuthorendpoint(s AuthorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateAuthorRequest)
		msg, err := s.UpdateAuthor(ctx, req.author)
		return msg, err
	}
}

func decodeCreateAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateAuthorRequest
	fmt.Println("-------->>>>into Decoding")
	if err := json.NewDecoder(r.Body).Decode(&req.author); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetAuthorByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetAuthorByIdRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = GetAuthorByIdRequest{
		Id: vars["authorid"],
	}
	return req, nil
}
func decodeGetAuthorByIdBooksRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetAuthorByIdBooksRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = GetAuthorByIdBooksRequest{
		Id: vars["authorid"],
	}
	return req, nil
}
func decodeDeleteAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Delete Decoding")
	var req DeleteAuthorRequest
	vars := mux.Vars(r)
	req = DeleteAuthorRequest{
		Authorid: vars["authorid"],
	}
	return req, nil
}
func decodeUpdateAuthorRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Update Decoding")
	var req UpdateAuthorRequest
	if err := json.NewDecoder(r.Body).Decode(&req.author); err != nil {
		return nil, err
	}
	return req, nil
}

type (
	CreateAuthorRequest struct {
		author Author
	}
	CreateAuthorResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}
	GetAuthorByIdRequest struct {
		Id string `json:"authorid"`
	}
	GetAuthorByIdResponse struct {
		Author interface{} `json:"author,omitempty"`
		Err    string      `json:"error,omitempty"`
	}

	DeleteAuthorRequest struct {
		Authorid string `json:"authorid"`
	}

	DeleteAuthorResponse struct {
		Msg string `json:"response"`
		Err error  `json:"error,omitempty"`
	}
	UpdateAuthorRequest struct {
		author Author
	}
	UpdateAuthorResponse struct {
		Msg string `json:"status,omitempty"`
		Err error  `json:"error,omitempty"`
	}

	GetAuthorByIdBooksRequest struct {
		Id string `json:"authorid"`
	}
	GetAuthorByIdBooksResponse struct {
		Books []Book `json:"books",omitempty`
		Err   string `json:"error,omitempty"`
	}
)
