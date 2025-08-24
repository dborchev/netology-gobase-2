package main

import (
  "storage"
  "todo"
  "fmt"
  "errors"
  "flag"
  "os"
)

dataFile := "tasks.json"

func main() {
    // 1. Считайте команду из os.Args[1] и аргументы в args := os.Args[2:].
    cmd := os.Args[1]
    args := os.Args[2:]
    
    // 2. Подгрузите текущие задачи:
    
    tasks, err := storage.LoadJSON(dataFile)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    
    // 3. Для каждой команды создайте flag.FlagSet, задайте нужные флаги и вызовите
    switch cmd {
    
        case "add":
            // add: проверка *desc, вызов todo.Add и storage.SaveJSON
        case "list":
            // list: флаг --filter, вызов todo.List и печать результатов
        case "complete":
            // complete: флаг --id, вызов соответствующей функции и сохранение
        case "delete":
            // delete: флаг --id, вызов соответствующей функции и сохранение
        case "load":
            // load: флаг --file, определение формата по расширению (filepath.Ext), загрузка через LoadJSON/LoadCSV, сохранение в tasks.json с помощью SaveJSON
        case "export":
            // export: флаги --format, --out, вызов SaveJSON/SaveCSV
    }
}