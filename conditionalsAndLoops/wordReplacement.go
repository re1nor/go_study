package main

import (
	// Импортируем пакет "time", чтобы использовать тип time.Time
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Метод time.Now возвращает значение time.Time, представляющее текущую дату и время
	var now time.Time = time.Now()

	// У значений time.Time имеется метод Year, который возвращает текущий год
	var year int = now.Year()

	/* Пакет strings содержит тип Reepalcer, который ищет подстроку в строке
	и заменяет каждое вхождение этой подстроки в другой строке */
	broken := "Поздравляю Вас с новым 2019 годом!"
	replacer := strings.NewReplacer("2019", strconv.Itoa(year))
	// Возвращает значение strings.Replacer, настроенное для замены всех "2019" в строке на нынешний год
	fixed := replacer.Replace(broken)

	fmt.Println("На данный момент", year, "год.")
	fmt.Println(fixed)

}
