package repository

import (
	"os"
	"testing"
)

func TestRepository_Subscribe(t *testing.T) {
	dir := mockDir(t)
	defer os.RemoveAll(dir)

	repository, err := New(dir)
	if err != nil {
		t.Fatalf("NewRepository() unexpected error = %v", err)
	}
	repo := repository.(*Repository)

	t.Run("succes", func(t *testing.T) {
		err = repo.Subscribe("queue", func(message any) {})
		if err != nil {
			t.Fatalf("Subscribe() unexpected error = %v", err)
		}
	})
}
