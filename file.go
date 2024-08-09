package fsutils

import (
	"bufio"
	"fmt"
	"os"
)

// File struct. Should be initialized with CreateFile() function.
type File struct {
	Path    string
	Content []string
}

// Outputs file to console. Format:
//
// | -- 'Path' -- |
//
// 'content'
//
// Example:
//
//	file, err := fsutils.CreateFile("hi.txt", "HI!", "BYE!", "YAY")
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	file.Output()
func (f File) Output() {
	fmt.Printf("| -- '%v' -- |\n", f.Path)

	for i := 0; i < len(f.Content); {
		for _, v := range f.Content {
			i++
			fmt.Printf("%v. %s\n", i, v)
		}
	}
}

// Creates file with 'os.Create()'.
// Uses 'bufio.Writer()' to paste content into file.
// Each element of array 'content' is line.
// Example:
//
//	file, err := fsutils.CreateFile("hi.txt", "HI!", "BYE!", "YAY")
//	if err != nil {
//		fmt.Println(err)
//	}
func CreateFile(path string, content ...string) (File, error) {
	file, nerr := os.Create(path)
	if nerr != nil {
		return File{}, nerr
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	for _, line := range content {
		_, serr := writer.WriteString(line + "\n")
		if serr != nil {
			return File{}, serr
		}
	}

	return File{Path: path, Content: content}, nil
}

// Removes file with 'os.Remove()'.
// Example:
//
//	file, err := fsutils.RemoveFile("hi.txt")
//	if err != nil {
//		fmt.Println(err)
//	}
func RemoveFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

// Writes content to file. Each element of 'content' is line.
// Example:
//
//	err := fsutils.WriteContent("hi.txt", "HI!", "BYE!", "YAY")
//	if err != nil {
//		fmt.Println(err)
//	}
func WriteContent(path string, content ...string) error {
	file, err := os.OpenFile(path, 666, os.FileMode(os.O_WRONLY))
	if err != nil {
		return err
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	for _, line := range content {
		_, serr := writer.WriteString(line + "\n")
		if serr != nil {
			return serr
		}
	}

	return nil
}
