---

# Задание

>Чек-лист готовности к домашнему заданию
>
>* Установлена версия Go не ниже 1.18
>* Инициализирован Go-модуль (go mod init)
>* Подготовлен файл input.txt с любым коротким текстом

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
