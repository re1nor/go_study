package main

import (
	"fmt"
	"strings"
)

func main() {

	// Короткое объявление переменных и присваевание значений
	name := "bulat"
	year, month, day := 2001, 12, 26
	education := true
	weight := 58.2

	// Использлвание переменных
	fmt.Println("Имя гражданина: ", strings.Title(name))
	fmt.Println("Дата рождения:", day, ".", month, ".", year)
	fmt.Println("Наличие образования:", education)
	fmt.Println("Вес:", weight)
}
