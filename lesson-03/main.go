// Программа потери веса на Марсе
package main

import "fmt"

func main() {
	fmt.Println("Мой вес на поверхности Марса равен", 70.0*0.3783, "килограммам, а мой возраст равен", 20*365/687, "годам.")

	// Форматирование вывода
	fmt.Printf("Мой вес на поверхности Марса равен %v килограммам", 70.0*0.3783)
	fmt.Printf(", а мой возраст равен %v годам.", 20*365/687)
}
