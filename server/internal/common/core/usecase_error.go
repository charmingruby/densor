package core

import "fmt"

func NewInternalErr(location string) error {
	return &ErrInternal{
		Message: fmt.Sprintf("Erro interno em: %s", location),
	}
}

type ErrInternal struct {
	Message string `json:"message"`
}

func (e *ErrInternal) Error() string {
	return e.Message
}

func NewNotFoundErr(entity string) error {
	return &ErrNotFound{
		Message: fmt.Sprintf("%s não encontrado", entity),
	}
}

type ErrNotFound struct {
	Message string `json:"message"`
}

func (e *ErrNotFound) Error() string {
	return e.Message
}

func NewConflictErr(field string) error {
	return &ErrConflict{
		Message: fmt.Sprintf("%s já está em uso", field),
	}
}

type ErrConflict struct {
	Message string `json:"message"`
}

func (e *ErrConflict) Error() string {
	return e.Message
}
