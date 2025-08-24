package storage

import (
  "errors"
  "os"
  "encoding/json"
  "todo"
)

func LoadJSON(path string) ([]todo.Task, error) {
    return []todo.Task{}, nil
}

func SaveJSON(path string, tasks []todo.Task) error {
    return nil
}