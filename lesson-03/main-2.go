// Сколько времени нужно для путешествия на Марс?
package main

import "fmt"

func main() {
	const hoursPerDay = 24
	var speed = 100800
	var distance = 96300000

	fmt.Println(distance/speed/hoursPerDay, "дней")

	distance = 401000000
	fmt.Println(distance/speed, "секунд")
}
