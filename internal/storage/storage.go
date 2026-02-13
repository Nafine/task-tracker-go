package storage

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Nafine/task-tracker/internal/model"
)

func Save(tasks model.Tasks) error {
	file, err := os.OpenFile("tasks.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tasks); err != nil {
		return err
	}
	return nil
}

func Load() (model.Tasks, error) {
	var tasks model.Tasks
	file, err := os.OpenFile("tasks.json", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil && err != io.EOF {
		return nil, err
	}
	return tasks, nil
}
