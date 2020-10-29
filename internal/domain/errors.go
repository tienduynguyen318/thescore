package domain

import (
	"fmt"
)

type NotFoundError struct {
	Err      error
	Name     string
	Resource string
}

func NewNotFoundError(resource, name string) *NotFoundError {
	return &NotFoundError{
		Name:     name,
		Resource: resource,
	}
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("Resource %s with name: %s does not exist.", e.Resource, e.Name)
}