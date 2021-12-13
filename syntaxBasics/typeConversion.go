package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int = 2
	var myFloat float64 = 4.5

	// Без преобразования
	fmt.Println("After: value:", myInt, "type:", reflect.TypeOf(myInt))

	// Преобразование int -> float64
	fmt.Println("int -> float64: value:", myInt, "type:", reflect.TypeOf(float64(myInt)))
	fmt.Println()

	//Преобразование float64 -> int отбрасывает дробную часть
	fmt.Println("After: value:", myFloat, "type:", reflect.TypeOf(myFloat))
	myInt = int(myFloat)
	fmt.Println("float64 -> int: value:", myInt, "type:", reflect.TypeOf(myInt))

}
