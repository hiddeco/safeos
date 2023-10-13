package safeos

import (
	"fmt"
	"os"
)

// Exists returns true if the given path exists. If there is an error
// accessing the path, it returns false.
func Exists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

// HasPermissions returns true if the given path has the given permissions.
// If the path does not exist, it returns false. If there is an error
// accessing the path, it returns the error.
func HasPermissions(name string, perm os.FileMode) error {
	info, err := os.Stat(name)
	if err != nil {
		return err
	}

	if modePerm := info.Mode().Perm(); modePerm != perm {
		return fmt.Errorf("%q has permissions 0%o, expected 0%o", name, modePerm, perm)
	}
	return nil
}
