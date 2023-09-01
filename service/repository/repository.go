package repository

import (
	"fmt"
	"os"
	"sync"

	"github.com/difmaj/studiosol-backend-challenge/service"
)

// Repository is a struct that defines the repository
type Repository struct {
	mu       sync.RWMutex
	channels map[string][]chan any
	dirPath  string
}

// New is a function that returns a new repository
func New(dirPath string) (service.Repository, error) {

	if dirPath == "" {
		return nil, fmt.Errorf("NewRepository: %w", ErrDirectoryCannotBeEmpty)
	}

	if err := createDirectory(dirPath); err != nil {
		return nil, fmt.Errorf("NewRepository: %w", err)
	}

	channels, err := loadChannels(dirPath)
	if err != nil {
		return nil, fmt.Errorf("NewRepository: %w", err)
	}

	return &Repository{
		dirPath:  dirPath,
		channels: channels,
	}, nil
}

// createDirectory creates a directory if it does not exist.
func createDirectory(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			return fmt.Errorf("createDirectory: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("createDirectory: %w", err)
	}
	return nil
}

func loadChannels(dirPath string) (map[string][]chan any, error) {

	channels := make(map[string][]chan any)
	files, err := os.ReadDir(dirPath)
	if os.IsNotExist(err) {
		return channels, nil
	} else if err != nil {
		return nil, fmt.Errorf("loadQueues: %w", err)
	}

	for _, file := range files {
		queue := file.Name()
		channels[queue] = []chan any{}
	}
	return channels, nil
}
