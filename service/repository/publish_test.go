package repository

import (
	"errors"
	"os"
	"testing"
)

func TestRepository_recordMessage(t *testing.T) {
	dir := mockDir(t)
	defer os.RemoveAll(dir)

	repository, err := New(dir)
	if err != nil {
		t.Fatal(err)
	}
	repo := repository.(*Repository)

	t.Run("succes", func(t *testing.T) {
		if err := repo.recordMessage("queue", "message"); err != nil {
			t.Errorf("Repository.recordMessage() unexpected error = %v", err)
		}
	})

	t.Run("error_invalid_queue", func(t *testing.T) {
		err := repo.recordMessage("_/\\", "message")
		if err == nil {
			t.Errorf("Repository.recordMessage() expected error = %v", err)
		}
	})
}
func TestRepository_Publish(t *testing.T) {

	dir := mockDir(t)
	defer os.RemoveAll(dir)

	repository, err := New(dir)
	if err != nil {
		t.Fatal(err)
	}
	repo := repository.(*Repository)

	t.Run("succes", func(t *testing.T) {
		if err := repo.Publish("queue", "message"); err != nil {
			t.Errorf("Repository.recordMessage() unexpected error = %v", err)
		}
	})

	t.Run("error_empty_queue", func(t *testing.T) {
		err := repo.Publish("", "message")
		if err == nil {
			t.Errorf("Repository.Publish() expected error = %v", err)
		}
		if !errors.Is(err, ErrQueueCannotBeEmpty) {
			t.Errorf("Repository.Publish() expected error %v, got = %v", ErrQueueCannotBeEmpty, err)
		}
	})

	t.Run("error_empty_message", func(t *testing.T) {
		err := repo.Publish("queue", nil)
		if err == nil {
			t.Errorf("Repository.Publish() expected error = %v", err)
		}
		if !errors.Is(err, ErrMessageCannotBeNil) {
			t.Errorf("Repository.Publish() expected error %v, got = %v", ErrMessageCannotBeNil, err)
		}
	})

	t.Run("error_invalid_queue", func(t *testing.T) {
		err := repo.Publish("_/\\", "message")
		if err == nil {
			t.Errorf("Repository.Publish() expected error = %v", err)
		}
	})
}
