package safeos

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestExists(t *testing.T) {
	t.Run("file exists", func(t *testing.T) {
		dir := t.TempDir()
		name := filepath.Join(dir, "file")
		_, _ = os.Create(name)

		if !Exists(name) {
			t.Errorf("Exists() = false, expected true")
		}
	})

	t.Run("file does not exist", func(t *testing.T) {
		dir := t.TempDir()
		name := filepath.Join(dir, "file")

		if Exists(name) {
			t.Errorf("Exists() = true, expected false")
		}
	})

	t.Run("directory exists", func(t *testing.T) {
		dir := t.TempDir()

		if !Exists(dir) {
			t.Errorf("Exists() = false, expected true")
		}
	})

	t.Run("directory does not exist", func(t *testing.T) {
		dir := t.TempDir()
		_ = os.Remove(dir)

		if Exists(dir) {
			t.Errorf("Exists() = true, expected false")
		}
	})
}

func TestHasPermissions(t *testing.T) {
	t.Run("file has permissions", func(t *testing.T) {
		dir := t.TempDir()
		name := filepath.Join(dir, "file")
		_, _ = os.Create(name)

		if err := HasPermissions(name, 0o644); err != nil {
			t.Fatalf("HasPermissions() = %v, expected nil", err)
		}
	})

	t.Run("file does not have permissions", func(t *testing.T) {
		dir := t.TempDir()
		name := filepath.Join(dir, "file")
		_, _ = os.Create(name)

		err := HasPermissions(name, 0o600)
		if err == nil {
			t.Fatalf("HasPermissions() = nil, expected error")
		}

		expectErr := fmt.Errorf("%q has permissions 0644, expected 0600", name)
		if err.Error() != expectErr.Error() {
			t.Errorf("HasPermissions() = %v != %v", err, expectErr)
		}
	})

	t.Run("directory has permissions", func(t *testing.T) {
		dir := t.TempDir()

		if err := HasPermissions(dir, 0o755); err != nil {
			t.Fatalf("HasPermissions() = %v, expected nil", err)
		}
	})

	t.Run("directory does not have permissions", func(t *testing.T) {
		dir := t.TempDir()

		err := HasPermissions(dir, 0o600)
		if err == nil {
			t.Fatalf("HasPermissions() = nil, expected error")
		}

		expectErr := fmt.Errorf("%q has permissions 0755, expected 0600", dir)
		if err.Error() != expectErr.Error() {
			t.Errorf("HasPermissions() = %v != %v", err, expectErr)
		}
	})

	t.Run("file does not exist", func(t *testing.T) {
		dir := t.TempDir()
		name := filepath.Join(dir, "file")

		err := HasPermissions(name, 0o600)
		if err == nil {
			t.Fatalf("HasPermissions() = nil, expected error")
		}
	})
}
