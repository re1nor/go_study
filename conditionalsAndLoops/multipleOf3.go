// multipleOf3 - программа, которая выводит числа от 1 до 100, делящееся на 3.
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Число кратное 3:  ", i)
		}
	}
}
