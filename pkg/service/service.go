package service

import "github.com/la4zen/dogma-intwrview/pkg/store"

type Service struct {
	Store *store.Store
}

func NewService(store *store.Store) *Service {
	return &Service{
		Store: store,
	}
}
