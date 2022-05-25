package file

import (
	"crypto/sha256"
	"io"
	"os"
)

func GetFileHash(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

func MakeFilePath(fileName string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	} else {
		return home + string(os.PathSeparator) + fileName, nil
	}
}
