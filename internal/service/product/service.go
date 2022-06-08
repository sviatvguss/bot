package product

import (
	"errors"
	"fmt"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	if idx > len(allProducts) {
		return nil, fmt.Errorf("we have have only %v products", len(allProducts))
	} else if idx <= 0 {
		return nil, errors.New("incorrect index")
	}
	return &allProducts[idx-1], nil
}
