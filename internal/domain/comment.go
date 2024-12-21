package domain

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch domain by ID")
	ErrNotImplemented  = errors.New("not implemented")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - All the methods the service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - The struct which all our logic will be built uon top of
type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{Store: store}
}

func (s *Service) GetComment(ctx context.Context, ID string) (Comment, error) {
	fmt.Println("Retrieve a domain")

	//ctx = context.WithValue(ctx, "request_id", "unique-string")
	//fmt.Println(ctx.Value("request-id"))

	comment, err := s.Store.GetComment(ctx, ID)
	if err != nil {
		fmt.Println(err) // Original error
		return Comment{}, ErrFetchingComment
	}

	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, comment Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, comment Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
