package todo

import (
	"encoding/json"
	"os"
)

type Store struct {
	FilePath string
}

func (s *Store) LoadTasks() ([]Task, error) {
	file, err := os.Open(s.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var Tasks []Task
	err = json.NewDecoder(file).Decode(&Tasks)
	if err != nil {
		return nil, err
	}

	return Tasks, nil
}

func (s *Store) SaveTasks(tasks []Task) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return os.WriteFile(s.FilePath, data, 0644)
}
