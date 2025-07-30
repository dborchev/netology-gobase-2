package main

import "fmt"

// Структура для хранения данных о сотруднике
type Employee struct {
    Surname string //Фамилия
    Name string    //Имя
    Age int        //Возраст
    Position string //Должность
    Salary int //Зарплата
}

// Интерфейс для вывода информации
type Displayable interface {
    Display() // выводит информацию о сотруднике
}

// Реализация метода Display для Employee
func (e Employee) Display() {
   fmt.Println("Информация о сотруднике:")
   fmt.Printf("\t%s %s, %d лет\n", e.Name, e.Surname, e.Age)
   fmt.Printf("\tДолжность: %s, ЗП: %d\n", e.Position, e.Salary)
}

// Функция фильтрации сотрудников по возрасту и зарплате
func FilterEmployees(employees []Employee, minAge int, minSalary int) []Employee {
    var filtered []Employee
    for _, e := range employees {
        if e.Age >= minAge && e.Salary >= minSalary {
            filtered = append(filtered, e)
        }
    }
    return filtered
}

func main() {
    // Инициализация списка сотрудников
    employees := []Employee{
        {
            Surname: "Цепешов",
            Name: "Владислав",
            Age: 576,
            Position: "Заведующий",
            Salary: 123489,
        },
        {
            Surname: "Иванова",
            Name: "Мария",
            Age: 28,
            Position: "Продавец-консультант",
            Salary: 42000,
        },
        {
            Surname: "Петров",
            Name: "Алексей",
            Age: 34,
            Position: "Старший продавец",
            Salary: 52000,
        },
        {
            Surname: "Смирнова",
            Name: "Елена",
            Age: 22,
            Position: "Кассир",
            Salary: 38000,
        },
        {
            Surname: "Кузнецов",
            Name: "Дмитрий",
            Age: 41,
            Position: "Менеджер по закупкам",
            Salary: 60000,
        },
        {
            Surname: "Волкова",
            Name: "Анастасия",
            Age: 26,
            Position: "Продавец мороженого",
            Salary: 40000,
        },
        {
            Surname: "Сидоров",
            Name: "Иван",
            Age: 30,
            Position: "Грузчик",
            Salary: 35000,
        },
        {
            Surname: "Морозова",
            Name: "Ольга",
            Age: 45,
            Position: "Бухгалтер",
            Salary: 58000,
        },
        {
            Surname: "Никитин",
            Name: "Сергей",
            Age: 39,
            Position: "Водитель-доставщик",
            Salary: 47000,
        },
        {
            Surname: "Федорова",
            Name: "Татьяна",
            Age: 31,
            Position: "Специалист по работе с клиентами",
            Salary: 50000,
        },
        {
            Surname: "Орлов",
            Name: "Михаил",
            Age: 24,
            Position: "Помощник продавца",
            Salary: 36000,
        },
    }

// Параметры фильтрации
    minAge := 30
    minSalary := 50000

// Фильтрация и вывод
    filteredEmployees := FilterEmployees(employees, minAge, minSalary)
    for _, s := range filteredEmployees {
		s.Display()
	}
}