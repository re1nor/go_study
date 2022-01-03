// fizzBuzz - программа, которая выводит числа от 1 до 100. Но для кратных трём нужно вывести «Fizz» вместо числа,
// для кратных пяти - «Buzz», а для кратных как трём, так и пяти — «FizzBuzz».
package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)

		}
	}
}
