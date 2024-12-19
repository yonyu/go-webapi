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

// The struct which all our logic will be built upon
type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetComment(ctx context.Context, ID string) (Comment, error) {
	fmt.Println("Retrieve a comment")
	return Comment{}, nil
}
