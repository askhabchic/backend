package client

import (
	"context"
)

type clientDAO interface {
	Create(ctx context.Context, client *Client) error
	FindAll(ctx context.Context, limit, offset int) (clients []Client, err error)
	FindOne(ctx context.Context, id string) (Client, error)
	Update(ctx context.Context, client Client, args ...interface{}) error
	Delete(ctx context.Context, id string) error
}

type Service struct {
	dao clientDAO
}

// interface repository from db/postgresql
func NewClientService(dao clientDAO) *Service {
	return &Service{
		dao: dao,
	}
}

func (s *Service) Create(ctx context.Context, client *Client) error {
	err := s.dao.Create(ctx, client)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteClient(ctx context.Context, id string) error {
	err := s.dao.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetByID(ctx context.Context, id string) (Client, error) {
	client, err := s.dao.FindOne(ctx, id)
	if err != nil {
		return Client{}, err
	}
	return client, nil
}

func (s *Service) GetAll(ctx context.Context, limit, offset int) (clients []Client, err error) {
	all, err := s.dao.FindAll(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (s *Service) UpdateClient(ctx context.Context, client Client, args ...interface{}) error {
	err := s.dao.Update(ctx, client, args)
	if err != nil {
		return err
	}
	return nil
}
