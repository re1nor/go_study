package main

import (
	"fmt"
	"strings"
)

func main() {
	//	Объявление переменных:
	//	Ключевое слово "var" имя_переменной тип
	var name string
	var year, month, day int
	var education bool
	// Объявление переменной и присвоение ей значения в одной строке
	var weight = 58.2

	// Присвоение значений переменным
	name = "bulat"
	year, month, day = 2001, 12, 26
	education = true

	// Использование переменных
	fmt.Println("Имя гражданина: ", strings.Title(name))
	fmt.Println("Дата рождения:", day, ".", month, ".", year)
	fmt.Println("Наличие образования:", education)
	fmt.Println("Вес:", weight)
}
