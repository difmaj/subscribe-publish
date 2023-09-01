package service

// UserCase
type UseCase interface {
	// Publish is a method that publishes a message to a queue
	Publish(queue string, message any) error
	// Subscribe is a method that subscribes to a queue
	Subscribe(queue string, callback func(message any)) error
}
