package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// True или False — логический тип данных в Go
	fmt.Println("Вы находитесь в темной пещере.")
	command := "выйти наружу"
	exit := strings.Contains(command, "наружу")
	fmt.Println("Вы покидаете пещеру:", exit)

	// ------------------------
	fmt.Println("------------------")

	// Операторы сравнения в Go
	fmt.Println("На знаке снаружи написано 'Несовершеннолетним вход запрещен'.")
	age := 41
	minor := age > 18
	fmt.Printf("В возрасте %v, я совершеннолетний? %v\n", age, minor)

	// ------------------------
	fmt.Println("------------------")

	// Условный оператор if-else в Go
	command = "идти на восток"
	if command == "идти на восток" {
		fmt.Println("Вы направляетесь к горе.")
	} else if command == "зайти внутрь" {
		fmt.Println("Вы заходите в пещеру, где будете жить до конца своей жизни.")
	} else {
		fmt.Println("Пока не совсем понятно.")
	}

	// ------------------------
	fmt.Println("------------------")

	// Логические операторы || и && в Go
	fmt.Println("На дворе 2000 год. Он високосный?")

	year := 2000
	leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)
	if leap {
		fmt.Println("Этот год високосный!")
	} else {
		fmt.Println("К сожалению, нет. Этот год не високосный.")
	}

	// ------------------------
	fmt.Println("------------------")

	// Оператор switch в Go
	fmt.Println("Здесь вход в пещеру и путь на восток.")
	command = "зайти внутрь"

	switch command {
	case "идти на восток":
		fmt.Println("Вы направляетесь к горе.")
	case "зайти в пещеру", "зайти внутрь":
		fmt.Println("Вы находитесь в тускло освещенной пещере.")
	case "прочитать знак":
		fmt.Println("На знаке написано 'Несовершеннолетним вход запрещен'.")
	default:
		fmt.Println("Пока не совсем понятно.")
	}

	// ------------------------
	fmt.Println("------------------")

	// Цикл for в Go
	count := 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(100) == 0 {
			break
		}
		count--

	}
	if count == 0 {
		fmt.Println("Запуск!")
	} else {
		fmt.Println("Запуск отменяется.")
	}

	// Бесконечный цикл
	degress := 0
	for {
		fmt.Println(degress)

		degress++
		if degress >= 360 {
			degress = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}
