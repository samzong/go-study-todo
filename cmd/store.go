package cmd

import (
	"encoding/json"
	"os"
)

type Task struct {
	Description string `json:"description"`
}

type Store struct {
	FilePath string
}

func (s *Store) Add(task Task) error {
	return nil
}

func (s *Store) Delete(task Task) error {
	return nil
}

func (s *Store) List() ([]Task, error) {
	file, err := os.Open(s.FilePath)

	if err != nil {
		if os.IsExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}

	defer file.Close()

	var tasks []Task
	err = json.NewDecoder(file).Decode(&tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil

}
