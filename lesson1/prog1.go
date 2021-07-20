// Package les2 Это главный пакет
package les2

import (
	"fmt"
	"os"
	"time"
)

//ErrorWithTime Это обявленная структура
type ErrorWithTime struct {
	text string
	time time.Time
}

//NewError Это создание ошибки
func NewError(text string) error {
	return &ErrorWithTime{
		text: text,
		time: time.Now(),
	}
}

//Это функция вывода ошибки
func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("error: %s, time: %s\n", e.text, e.time)
}

//Главная функция
func main() {
	f, err := os.Create("new_file")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, _ = fmt.Fprintln(f, "data")
	input()
}

//Функция ввода данных
func input() {
	defer func() {
		if t := recover(); t != nil {
			fmt.Println(NewError("Recovered"), t)
			input()
		}
	}()

	var a, b, r int
	var op string

	fmt.Print("Input a: ")
	if _, err := fmt.Scanln(&a); err != nil {
		fmt.Println(NewError("a must be numeric"))
		return
	}

	fmt.Print("Input b: ")
	if _, err := fmt.Scanln(&b); err != nil {
		fmt.Println(NewError("b must be numeric"))
		return
	}

	fmt.Print("Input op: ")
	if _, err := fmt.Scanln(&op); err != nil {

		return
	}
	r = calc(a, b, op)
	fmt.Println(a, op, b, "=", r)
}

//Функция выполнения операций
func calc(a int, b int, op string) (r int) {
	switch op {
	case "+":
		r = a + b
	case "-":
		r = a - b
	case "*":
		r = a * b
	case "/":
		r = a / b

	default:
		fmt.Println(NewError("Operation must be +, -, /, *, ^"))
		return
	}
	return r
}
