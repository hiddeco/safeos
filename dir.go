package safeos

import "os"

// MkdirAll creates a directory at the given path with the given permissions.
// If the directory already exists, it confirms that it has the given
// permissions.
//
// Usage of this function is preferred over os.MkdirAll in certain applications
// where the permissions of the directory are important. For example, if you
// are creating a directory that will contain sensitive information, you may
// want to ensure that the directory is only readable by the owner.
func MkdirAll(path string, mode os.FileMode) error {
	if Exists(path) {
		return HasPermissions(path, mode)
	}
	return os.MkdirAll(path, mode)
}
