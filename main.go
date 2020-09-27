package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	r := mux.NewRouter()

	var svc BookService
	svc = NewService(logger)

	// svc = loggingMiddleware{logger, svc}
	// svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	CreateBookHandler := httptransport.NewServer(
		makeCreateBookEndpoint(svc),
		decodeCreateBookRequest,
		encodeResponse,
	)
	GetByBookIdHandler := httptransport.NewServer(
		makeGetBookByIdEndpoint(svc),
		decodeGetBookByIdRequest,
		encodeResponse,
	)
	DeleteBookHandler := httptransport.NewServer(
		makeDeleteBookEndpoint(svc),
		decodeDeleteBookRequest,
		encodeResponse,
	)
	UpdateBookHandler := httptransport.NewServer(
		makeUpdateBookendpoint(svc),
		decodeUpdateBookRequest,
		encodeResponse,
	)
	GetByBookIdPusblishersHandler := httptransport.NewServer(
		makeGetBookByIdPublishersEndpoint(svc),
		decodeGetBookByIdPublishersRequest,
		encodeResponse,
	)
	GetByBookIdAuthorsHandler := httptransport.NewServer(
		makeGetBookByIdAuthorsEndpoint(svc),
		decodeGetBookByIdAuthorsRequest,
		encodeResponse,
	)

	http.Handle("/", r)
	http.Handle("/book", CreateBookHandler)
	http.Handle("/book/update", UpdateBookHandler)
	r.Handle("/book/{bookid}", GetByBookIdHandler).Methods("GET")
	r.Handle("/book/{bookid}/publishers", GetByBookIdPusblishersHandler).Methods("GET")
	r.Handle("/book/{bookid}/authors", GetByBookIdAuthorsHandler).Methods("GET")
	r.Handle("/book/{bookid}", DeleteBookHandler).Methods("DELETE")

	var svcPublisher PublisherService
	svcPublisher = NewServicePublisher(logger)

	// svc = loggingMiddleware{logger, svc}
	// svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	CreatePublisherHandler := httptransport.NewServer(
		makeCreatePublisherEndpoint(svcPublisher),
		decodeCreatePublisherRequest,
		encodeResponse,
	)
	GetByPublisherIdHandler := httptransport.NewServer(
		makeGetPublisherByIdEndpoint(svcPublisher),
		decodeGetPublisherByIdRequest,
		encodeResponse,
	)
	DeletePublisherHandler := httptransport.NewServer(
		makeDeletePublisherEndpoint(svcPublisher),
		decodeDeletePublisherRequest,
		encodeResponse,
	)
	UpdatePublisherHandler := httptransport.NewServer(
		makeUpdatePublisherendpoint(svcPublisher),
		decodeUpdatePublisherRequest,
		encodeResponse,
	)
	GetByPublisherIdBooksHandler := httptransport.NewServer(
		makeGetPublisherByIdBooksEndpoint(svcPublisher),
		decodeGetPublisherByIdBooksRequest,
		encodeResponse,
	)

	http.Handle("/publisher", CreatePublisherHandler)
	http.Handle("/publisher/update", UpdatePublisherHandler)
	r.Handle("/publisher/{publisherid}", GetByPublisherIdHandler).Methods("GET")
	r.Handle("/publisher/{publisherid}/books", GetByPublisherIdBooksHandler).Methods("GET")
	r.Handle("/publisher/{publisherid}", DeletePublisherHandler).Methods("DELETE")

	var svcAuthor AuthorService
	svcAuthor = NewServiceAuthor(logger)

	// svc = loggingMiddleware{logger, svc}
	// svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	CreateAuthorHandler := httptransport.NewServer(
		makeCreateAuthorEndpoint(svcAuthor),
		decodeCreateAuthorRequest,
		encodeResponse,
	)
	GetByAuthorIdHandler := httptransport.NewServer(
		makeGetAuthorByIdEndpoint(svcAuthor),
		decodeGetAuthorByIdRequest,
		encodeResponse,
	)
	DeleteAuthorHandler := httptransport.NewServer(
		makeDeleteAuthorEndpoint(svcAuthor),
		decodeDeleteAuthorRequest,
		encodeResponse,
	)
	UpdateAuthorHandler := httptransport.NewServer(
		makeUpdateAuthorendpoint(svcAuthor),
		decodeUpdateAuthorRequest,
		encodeResponse,
	)
	GetByAuthorIdBooksHandler := httptransport.NewServer(
		makeGetAuthorByIdBooksEndpoint(svcAuthor),
		decodeGetAuthorByIdBooksRequest,
		encodeResponse,
	)

	http.Handle("/author", CreateAuthorHandler)
	http.Handle("/author/update", UpdateAuthorHandler)
	r.Handle("/author/{authorid}", GetByAuthorIdHandler).Methods("GET")
	r.Handle("/author/{authorid}/books", GetByAuthorIdBooksHandler).Methods("GET")
	r.Handle("/author/{authorid}", DeleteAuthorHandler).Methods("DELETE")

	// http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
	logger.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
