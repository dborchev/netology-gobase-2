---

# Задание

>Чек-лист готовности к домашнему заданию
>
>* Установлена версия Go не ниже 1.18✅
>* Инициализирован Go-модуль (go mod init)✅
>* Подготовлен файл input.txt с любым коротким текстом✅

>Написать функцию ReadProcessWrite, которая читает весь текст из файла
>inputPath, передаёт его в функцию process и записывает результат в файл
>outputPath. В случае ошибки на любом этапе возвращает её вызывающему.

>Напишите функцию:
```go
func ReadProcessWrite(
inputPath string,
outputPath string,
process func(string) (string, error),
) error
```


# Выполнение

1. Инициализируйте модуль проекта: go mod init module_name(имя вашего модуля) ✅
1. Создайте файл main.go✅
  * [main.go](main.go) 

1. Импортируйте пакеты: ✅
1. Опишите функцию ReadProcessWrite. Реализуйте в ней логику чтения, обработки и записи с проверкой ошибок после каждой операции.✅
1. В функции main продемонстрируйте работу ✅
1. При вызове программы попробуйте вызвать ошибки. К примеру укажите неверный путь к файлу inputPath. И проверьте как происходит обработка ошибок при выводе в консоль.

### нормальное исполнение
```go
func main() {
    err := ReadProcessWrite("input.txt", "output.txt", strings_to_upper)
    if errors.Is(err, ErrToUpperAlreadyUp) {
        log.Printf("данные уже в верхнем регистре")
    } else if err != nil {
        log.Fatalf("Ошибка: %w", err)
    }
}
```
```
> go run .\main.go
> cat .\output.txt
OCTOPUSES HAVE THREE HEARTS AND BLUE BLOOD TO HELP TRANSPORT OXYGEN EFFICIENTLY UNDERWATER.
A GROUP OF FLAMINGOS IS CALLED A "FLAMBOYANCE"—FITTING FOR SUCH VIVIDLY COLORED BIRDS.
ELEPHANTS CAN RECOGNIZE THEMSELVES IN MIRRORS, SHOWING SIGNS OF SELF-AWARENESS.
SEA OTTERS HOLD HANDS WHILE SLEEPING TO AVOID DRIFTING APART IN THE WATER.
COWS HAVE BEST FRIENDS AND GET STRESSED WHEN THEY ARE SEPARATED FROM THEM.
```


### файл не найден
```go
func main() {
    err := ReadProcessWrite("DoesNotExist.txt", "output.txt", strings_to_upper)
    if errors.Is(err, ErrToUpperAlreadyUp) {
        log.Printf("данные уже в верхнем регистре")
    } else if err != nil {
        log.Fatalf("Ошибка: %w", err)
    }
}
```
```
> go run .\main.go
2025/07/20 16:16:59 Ошибка: %!w(*fmt.wrapError=&{Ошибка чтения файла DoesNotExist.txt: open DoesNotExist.txt: The system cannot find the file specified. 0xc00001e210})
exit status 1
```

### файл уже в верхнем регистре
```go
func main() {
    err := ReadProcessWrite("output.txt", "output.txt", strings_to_upper)
    if errors.Is(err, ErrToUpperAlreadyUp) {
        log.Printf("данные уже в верхнем регистре")
    } else if err != nil {
        log.Fatalf("Ошибка: %w", err)
    }
}
```
```
> go run .\main.go
2025/07/20 16:20:39 данные уже в верхнем регистре
```

### файл защищён от записи
```go
func main() {
    err := ReadProcessWrite("input.txt", "output.txt", strings_to_upper)
    if errors.Is(err, ErrToUpperAlreadyUp) {
        log.Printf("данные уже в верхнем регистре")
    } else if err != nil {
        log.Fatalf("Ошибка: %w", err)
    }
}
```
```
> go run .\main.go
2025/07/20 16:22:18 Ошибка: %!w(*fmt.wrapError=&{Ошибка создания файла output.txt: open output.txt: Access is denied. 0xc00001e210})
exit status 1
```
