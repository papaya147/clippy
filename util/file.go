package util

import (
	"os"
	"path/filepath"
)

func CheckFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func CreateFile(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return err
	}
	if _, err := os.Create(path); err != nil {
		return err
	}
	return nil
}

func WriteFile(path string, data []byte) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	if _, err := file.Write(data); err != nil {
		return err
	}
	return nil
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	data := make([]byte, 1024)
	n, err := file.Read(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}
