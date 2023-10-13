package safeos

import (
	"path/filepath"
	"testing"
)

func TestMkdirAll(t *testing.T) {
	t.Run("create directory", func(t *testing.T) {
		tmp := t.TempDir()
		dir := filepath.Join(tmp, "dir")

		if err := MkdirAll(dir, 0o755); err != nil {
			t.Fatalf("MkdirAll() = %v, expected nil", err)
		}
		if err := HasPermissions(dir, 0o755); err != nil {
			t.Fatalf("HasPermissions() = %v, expected nil", err)
		}
	})

	t.Run("directory already exists", func(t *testing.T) {
		tmp := t.TempDir()

		if err := MkdirAll(tmp, 0o755); err != nil {
			t.Fatalf("MkdirAll() = %v, expected nil", err)
		}
	})

	t.Run("directory already exists with wrong permissions", func(t *testing.T) {
		tmp := t.TempDir()

		if err := MkdirAll(tmp, 0o600); err == nil {
			t.Fatalf("MkdirAll() = nil, expected error")
		}
	})
}
