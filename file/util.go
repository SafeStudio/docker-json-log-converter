package file

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func IsExists(filePath string) bool {
	if len(filePath) == 0 {
		fmt.Println("Empty file path is not allowed.")
		return false
	}

	absolutePath := filePath
	var err error

	if filePath[0] != '/' {
		rfp := "./" + filePath

		absolutePath, err = GetAbsolutePath(rfp)

		if err != nil {
			fmt.Println("Cannot resolve file path", rfp)
			return false
		}
	}

	_, err = os.Stat(absolutePath)

	return !errors.Is(err, fs.ErrNotExist)
}

func GetAbsolutePath(relativePath string) (string, error) {
	abs, err := filepath.Abs(relativePath)
	if err != nil {
		return "", err
	}

	return abs, nil
}
