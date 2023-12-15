package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Println("Сегодня понедельник.")
	case time.Tuesday:
		fmt.Println("Сегодня вторник.")
	case time.Wednesday:
		fmt.Println("Сегодня среда.")
	case time.Thursday:
		fmt.Println("Сегодня четверг.")
	case time.Friday:
		fmt.Println("Сегодня пятница.")
	case time.Saturday:
		fmt.Println("Сегодня суббота.")
	case time.Sunday:
		fmt.Println("Сегодня воскресенье.")
	}

	// -----------------------
	// Несколько выражений в операторе switch в Go
	switch time.Now().Weekday() {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Будний день.")
	case time.Saturday, time.Sunday:
		fmt.Println("Выходной день.")
	}

	// -----------------------
	fmt.Println("-----------")
	// Оператор switch пример с default в Go
	size := "XXXL"

	switch size {
	case "XXS":
		fmt.Println("очень очень маленький")
	case "XS":
		fmt.Println("очень маленький")
	case "S":
		fmt.Println("маленький")
	case "M":
		fmt.Println("средний")
	case "L":
		fmt.Println("большой")
	case "XL":
		fmt.Println("очень большой")
	case "XXL":
		fmt.Println("очень очень большой")
	default:
		fmt.Println("неизвестно")
	}

	// -----------------------
	fmt.Println("-----------")
	// Необязательный оператор при работе со switch в Go
	switch num := 6; num%2 == 0 {

	case true:
		fmt.Println("even value")

	case false:
		fmt.Println("odd value")
	}

	// -----------------------
	fmt.Println("-----------")
	// Оператор break в Go
	w := "a b c\td\nefg hi"

	for _, e := range w {
		switch e {
		case ' ', '\t', '\n':
			break
		default:
			fmt.Printf("%c\n", e)
		}
	}

	// -----------------------
	fmt.Println("-----------")
	// Оператор switch без выражения в Go
	now := time.Now()

	switch {
	case now.Hour() < 12:
		fmt.Println("AM")

	default:
		fmt.Println("PM")
	}

	// -----------------------
	fmt.Println("-----------")
	// Ключевое слово fallthrough в switch
	nextStop := "B"
	fmt.Println("Stops ahead of us:")

	switch nextStop {
	case "A":
		fmt.Println("A")
		fallthrough

	case "B":
		fmt.Println("B")
		fallthrough

	case "C":
		fmt.Println("C")
		fallthrough

	case "D":
		fmt.Println("D")
		fallthrough

	case "E":
		fmt.Println("E")
	}

	// -----------------------
	fmt.Println("-----------")
	// Использование type в switch
	var data interface{}
	data = false

	switch myType := data.(type) {
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("boolean")
	case float64:
		fmt.Println("float64 type")
	case float32:
		fmt.Println("float32 type")
	case int:
		fmt.Println("int")
	default:
		fmt.Printf("%T", myType)
	}
}
