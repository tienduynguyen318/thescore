package httprest

import (
	"fmt"
	"net/http"
	"thescore/internal/domain"
)

type errorsCollection struct {
	Errors []errorDetail `json:"errors"`
}

type errorDetail struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func newInternalServerError(err error) *httpError {
	return &httpError{
		Errors: []error{fmt.Errorf("Something went wrong on the server")},
		Status: http.StatusInternalServerError,
		LogMsg: err.Error(),
	}
}

func newNotFoundError(err error) *httpError {
	return &httpError{
		Errors: []error{err},
		Status: http.StatusNotFound,
		LogMsg: err.Error(),
	}
}

func newHTTPErrorFromDomain(err error) *httpError {
	var httpError httpError
	switch e := err.(type) {
	case *domain.NotFoundError:
		httpError = *newNotFoundError(e)
	default:
		httpError = *newInternalServerError(e)
	}
	return &httpError
}

type httpError struct {
	Errors []error
	Status int
	LogMsg string
}

func (e *httpError) Payload() errorsCollection {
	var details []errorDetail
	for _, error := range e.Errors {
		detail := errorDetail{
			Title:  http.StatusText(e.Status),
			Status: e.Status,
			Detail: error.Error(),
		}
		details = append(details, detail)
	}
	return errorsCollection{
		Errors: details,
	}
}
