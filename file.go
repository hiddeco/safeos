package safeos

import "os"

// WriteFile writes data to a file named by filename. If the file does already
// exist, it confirms it has the given permissions before truncating and
// writing to it.
//
// Usage of this function is preferred over os.WriteFile in certain applications
// where the permissions of the file are important. For example, if you are
// writing a file that will contain sensitive information, you may want to
// ensure that the file is only readable by the owner.
func WriteFile(name string, data []byte, perm os.FileMode) error {
	if Exists(name) {
		if err := HasPermissions(name, perm); err != nil {
			return err
		}
	}
	return os.WriteFile(name, data, perm)
}
