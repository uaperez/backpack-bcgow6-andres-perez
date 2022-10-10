package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

const (
	FileType Type = "file"
	MonoType Type = "mongo"
)

// Factory de Stores
func New(store Type, filename string) Store {
	switch store {
	case FileType:
		return &fileStore{filename}
	}
	return nil
}

type fileStore struct {
	FilePath string
}

func (fs *fileStore) Read(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}

func (fs *fileStore) Write(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}
