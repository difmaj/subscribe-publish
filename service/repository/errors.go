package repository

import (
	"errors"
)

var (
	// ErrQueueDoesNotExist is an error that occurs when a queue does not exist.
	ErrQueueDoesNotExist = errors.New("queue does not exist")
	// ErrQueueCannotBeEmpty is an error that occurs when a queue is empty.
	ErrQueueCannotBeEmpty = errors.New("queue cannot be empty")
	// ErrMessageCannotBeNil is an error that occurs when a message is nil.
	ErrMessageCannotBeNil = errors.New("message cannot be nil")
	// ErrDirectoryCannotBeEmpty is an error that occurs when a directory is empty.
	ErrDirectoryCannotBeEmpty = errors.New("directory cannot be empty")
)
