package comment

import (
	"context"
	"fmt"
)

// Comment - a representation of a Comment structure
// for our service
type Comment struct {
	ID     int
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Serive - a struct on which all of our
// logic will be build on top of
type Service struct {
	Store Store
}

// Constructor function for the Service
// NewService returns a pointer to the new Service
func NewService(store Store) *Service {
	return &Service{Store: store}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Get the comment id : ", id)
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil

}
