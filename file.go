package php

import (
	"io/ioutil"
	"os"
)

// php file_get_contents()
func FileGetContents(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	return string(content), err
}

// php file_put_contents()
func FilePutContents(filename string, data string, mode os.FileMode) error {
	return ioutil.WriteFile(filename, []byte(data), mode)
}

// php unlink()
func Unlink(filename string) error {
	return os.Remove(filename)
}
