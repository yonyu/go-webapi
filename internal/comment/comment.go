package comment

import (
	"context"
	"fmt"
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// All the methods the service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// The struct which all our logic will be built uon top of
type Service struct {
	Store Store
}

func NewService(store Store) *Service {
	return &Service{Store: store}
}

func (s *Service) GetComment(ctx context.Context, ID string) (Comment, error) {
	fmt.Println("Retrieve a comment")
	cmt, err := s.Store.GetComment(ctx, ID)
	if err != nil {
		return Comment{}, err
	}

	return cmt, nil
}
