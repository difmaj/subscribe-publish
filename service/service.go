package service

import (
	"fmt"
)

// Service is a struct that defines the service
type Service struct {
	repository Repository
}

// New is a function that returns a new service
func New(repository Repository) UseCase {
	return &Service{
		repository: repository,
	}
}

// Publish is a method that publishes a message to a queue
func (s *Service) Publish(queue string, message any) error {
	if err := s.repository.Publish(queue, message); err != nil {
		return fmt.Errorf("Service.Publish: %w", err)
	}
	return nil
}

// Subscribe is a method that subscribes to a queue
func (s *Service) Subscribe(queue string, callback func(message any)) error {
	if err := s.repository.Subscribe(queue, callback); err != nil {
		return fmt.Errorf("Service.Subscribe: %w", err)
	}
	return nil
}
