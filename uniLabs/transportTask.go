/*
	transportTask -  решение транспортной задачи линейного программирования

	TODO:
	- Возможность работы с различными размерностями 		/ ENG: Ability to work with different dimensions
	-
	DONE:
	- Инициализация поставщиков, потребителей и стоимости перевозок / ENG: Initialization of suppliers, consumers and transportation costs
	- Проверка на замкнутость 					/ ENG: Checking for closeness
	- Заполнение 2D-массива элементами 				/ ENG: Filling a 2D array with elements
	-
*/

package main

import (
	"fmt"
)

func sumArray(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func outputArray(array [3][4]int) {
	for _, value := range array {
		fmt.Println(value)
	}
}

// - Заполнение 2D-массива элементами / ENG: Filling a 2D array with elements
func filling2DArray(array *[3][4]int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[0]); j++ {
			fmt.Println("Введите значение", i+1, j+1, "элемента:")
			fmt.Scan(&array[i][j])
			fmt.Println("Результат:")
			outputArray(*array)
		}
	}
}

func main() {
	// Инициализация поставщиков, потребителей и стоимости перевозок
	// ENG: Initialization of suppliers, consumers and transportation costs
	provider := [3]int{30, 40, 20}
	client := [4]int{20, 25, 30, 15}
	transportCost := [len(provider)][len(client)]int{}

	// Вывод информации:
	fmt.Println("Provider:", provider)
	fmt.Println("Client:", client)

	// Проверка на замкнутость / ENG: Checking for closeness
	if sumArray(provider[:]) == sumArray(client[:]) {
		fmt.Println("Условие замкнутости соблюдено! \nКоличество поставщиков = Количество потребителей")
		filling2DArray(&transportCost)
	} else {
		fmt.Println("Условие замкнутости не соблюдено! \nКоличество поставщиков != Количество потребителей")
	}

}
