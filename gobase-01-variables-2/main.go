package main

import "fmt"

const ProgramName string = "GoFit"

const Output string = "Добро пожаловать в %s!\nПрофиль пользователя:\nИмя: %s\nВозраст: %d лет\nРост: %.2f м\nПодписан на рассылку: %t\n"

type User struct {
	name         string
	age          int
	height       float64
	isSubscriber bool
}

func main() {

	var user User
	user.name = "Иван"
	user.age = 30
	user.height = 1.75
	user.isSubscriber = true

	fmt.Printf(
		Output,
		ProgramName,
		user.name,
		user.age,
		user.height,
		user.isSubscriber,
	)

}
