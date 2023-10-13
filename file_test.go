package safeos

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFile(t *testing.T) {
	t.Run("write to existing file", func(t *testing.T) {
		dir := t.TempDir()
		name := filepath.Join(dir, "file")

		f, err := os.Create(name)
		if err != nil {
			t.Fatalf("os.Create() = %v, expected nil", err)
		}
		if err := f.Close(); err != nil {
			t.Fatalf("f.Close() = %v, expected nil", err)
		}

		if err := WriteFile(name, []byte("hello"), 0o644); err != nil {
			t.Fatalf("WriteFile() = %v, expected nil", err)
		}
	})

	t.Run("write to new file", func(t *testing.T) {
		dir := t.TempDir()
		name := filepath.Join(dir, "file")

		if err := WriteFile(name, []byte("hello"), 0o600); err != nil {
			t.Fatalf("WriteFile() = %v, expected nil", err)
		}
		if err := HasPermissions(name, 0o600); err != nil {
			t.Fatalf("HasPermissions() = %v, expected nil", err)
		}
	})

	t.Run("write to file with wrong permissions", func(t *testing.T) {
		dir := t.TempDir()
		name := filepath.Join(dir, "file")

		f, err := os.Create(name)
		if err != nil {
			t.Fatalf("os.Create() = %v, expected nil", err)
		}
		if err := f.Close(); err != nil {
			t.Fatalf("f.Close() = %v, expected nil", err)
		}

		if err = WriteFile(name, []byte("hello"), 0o600); err == nil {
			t.Fatalf("WriteFile() = nil, expected error")
		}
	})
}
