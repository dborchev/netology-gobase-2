package main

import (
"fmt"
"log"
"os"
"strings"
"errors"
)

func ReadProcessWrite(
inputPath string,
outputPath string,
process func(string) (string, error),
) error {
    // Открыть файл inputPath и прочитать всё содержимое.
    data, err := os.ReadFile(inputPath)
    if err != nil {
        return fmt.Errorf("Ошибка чтения файла %v: %w", inputPath, err)
    }
    // Использовать defer для закрытия исходного файла (если используется os.Open).
    // -- не используется
    
    // Вызвать process, передав прочитанный текст.
    processed_data, err := process(string(data))
    if err != nil {
        return fmt.Errorf("Ошибка обработки: %w", err)
    }
    
    // Создать или перезаписать файл outputPath.
    file, err := os.Create(outputPath)
    if err != nil {
        return fmt.Errorf("Ошибка создания файла %v: %w", outputPath, err)
    }
    // Закрыть файл записи с помощью defer
    defer file.Close()

    // Записать в него результат process.
    _, err = file.WriteString(processed_data)
    if err != nil {
        return fmt.Errorf("Ошибка записи в файл %v: %w", outputPath, err)
    }

    return nil
}

var ErrToUpperAlreadyUp = errors.New("строка уже в верхнем регистре")

func strings_to_upper(str string) (string, error) {
    // передать strings.ToUpper третьим аргументом предварительно обернув в функцию возвращающую ошибку
    upp_data := strings.ToUpper(str)
    var err error = nil
    if upp_data == str {
        err = fmt.Errorf("strings.ToUpper: %w", ErrToUpperAlreadyUp)
    }
    return upp_data, err
}

func main() {
    err := ReadProcessWrite("input.txt", "output.txt", strings_to_upper)
    if errors.Is(err, ErrToUpperAlreadyUp) {
        log.Printf("данные уже в верхнем регистре")
    } else if err != nil {
        log.Fatalf("Ошибка: %w", err)
    }
}