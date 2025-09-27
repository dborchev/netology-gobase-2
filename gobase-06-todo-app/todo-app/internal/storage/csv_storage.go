package storage

import (
  "errors"
  "os"
  "encoding/json"
  "todo"
)

func LoadCSV(path string) ([]todo.Task, error) {
    return []todo.Task{}, nil
}

func SaveCSV(path string, tasks []todo.Task) error {
    return nil
}