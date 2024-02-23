package store

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/papaya147/clippy/util"
)

type Store struct {
	file string
	data map[string]string
}

func NewStore(file string) *Store {
	if !util.CheckFileExists(file) {
		if err := util.CreateFile(file); err != nil {
			fmt.Println("unable to create store file: ", err)
		}
	}

	return &Store{file: file}
}

func (store *Store) Get(key string) (string, error) {
	if store.data == nil {
		fileData, err := util.ReadFile(store.file)
		if err != nil {
			return "", err
		}
		if err := json.Unmarshal(fileData, &store.data); err != nil {
			return "", err
		}
	}

	if value, ok := store.data[key]; ok {
		return value, nil
	}

	return "", errors.New("key does not exist")
}

func (store *Store) Set(key, value string) error {
	if store.data == nil {
		store.data = make(map[string]string)
	}

	store.data[key] = value

	data, _ := json.Marshal(store.data)
	if err := util.WriteFile(store.file, data); err != nil {
		fmt.Println("error saving to store file: ", err)
	}

	return nil
}
