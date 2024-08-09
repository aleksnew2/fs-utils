package fsutils

import (
	"os"
	"strings"
)

// Creates directory with 'os.Mkdir()'.
// Example:
//
//	err := fsutils.CreateDir("directory")
//	if err != nil {
//		fmt.Println(err)
//	}
func CreateDir(path string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// Removes directory with 'os.RemoveAll()'.
// If 'path' hasn't had suffix '\',
// then function puts this on the end of 'path'.
// Example:
//
//	err := fsutils.RemoveDir("directory")
//	if err != nil {
//		fmt.Println(err)
//	}
func RemoveDir(path string) error {
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	err := os.RemoveAll(path)
	if err != nil {
		return err
	}

	return nil
}
