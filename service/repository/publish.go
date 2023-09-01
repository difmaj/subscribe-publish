package repository

import (
	"fmt"
	"os"
	"path/filepath"
)

// Publish sends a message to a queue.
// If the queue does not exist, it returns an error.
func (r *Repository) Publish(queue string, message any) error {

	switch {
	case queue == "":
		return ErrQueueCannotBeEmpty
	case message == nil:
		return ErrMessageCannotBeNil
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	if err := r.recordMessage(queue, message); err != nil {
		return err
	}

	chs, exists := r.channels[queue]
	if !exists {
		return nil
	}

	for _, ch := range chs {
		go func(ch chan<- any) {
			ch <- message
		}(ch)
	}

	return nil
}

// recordMessage records the message in a file.
func (r *Repository) recordMessage(queue, message any) error {

	filePath := filepath.Join(r.dirPath, fmt.Sprintf("%s.txt", queue))
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("recordMessage: %w", err)
	}

	_, err = file.Write([]byte(fmt.Sprint(message) + "\n"))
	file.Close()
	if err != nil {
		return fmt.Errorf("recordMessage: %w", err)
	}
	return nil
}
