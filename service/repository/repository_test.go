package repository

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func mockDir(t *testing.T) string {
	path := strings.NewReplacer("/", "-").Replace(t.Name())
	dir, err := os.MkdirTemp("", path)
	if err != nil {
		t.Fatal(err)
	}
	return filepath.Join(dir, "test")
}

func TestNew(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		dir := mockDir(t)
		defer os.RemoveAll(dir)

		_, err := New(dir)
		if err != nil {
			t.Errorf("NewRepository() unexpected error = %v", err)
		}
	})

	t.Run("empty_directory", func(t *testing.T) {
		_, err := New("")
		if err == nil {
			t.Errorf("NewRepository() expected error = %v", err)
		}
		if !errors.Is(err, ErrDirectoryCannotBeEmpty) {
			t.Errorf("NewRepository() expected error = %v", ErrDirectoryCannotBeEmpty)
		}
	})

	t.Run("error_invalid_directory", func(t *testing.T) {
		_, err := New("_/~\\")
		if err == nil {
			t.Errorf("NewRepository() expected error = %v", err)
		}
	})
}

func Test_createDirectory(t *testing.T) {
	t.Run("succes", func(t *testing.T) {
		dir := mockDir(t)
		defer os.RemoveAll(dir)

		_, err := New(dir)
		if err != nil {
			t.Errorf("NewRepository() unexpected error = %v", err)
		}
	})
}

func Test_loadChannels(t *testing.T) {
	dir := mockDir(t)
	defer os.RemoveAll(dir)

	repository, err := New(dir)
	if err != nil {
		t.Fatalf("NewRepository() unexpected error = %v", err)
	}
	repo := repository.(*Repository)

	t.Run("succes", func(t *testing.T) {

		mockQueue := "queue"
		err = repo.Subscribe(mockQueue, func(message any) {})
		if err != nil {
			t.Fatalf("Subscribe() unexpected error = %v", err)
		}

		err = repo.Publish(mockQueue, "message")
		if err != nil {
			t.Fatalf("Publish() unexpected error = %v", err)
		}

		channels, err := loadChannels(dir)
		if err != nil {
			t.Fatalf("loadChannels() unexpected error = %v", err)
		}

		if len(channels) != 1 {
			t.Fatalf("loadChannels() expected len = %v", 1)
		}

		if reflect.DeepEqual(channels, repo.channels) {
			t.Fatalf("loadChannels() expected channels = %v", repo.channels)
		}
	})

	t.Run("non-existent_directory", func(t *testing.T) {
		channels, err := loadChannels("_/\\")
		if err != nil {
			t.Fatalf("loadChannels() unexpected error = %v", err)
		}

		if len(channels) != 0 {
			t.Fatalf("loadChannels() unexpected values in channels = %v", channels)
		}
	})
}
