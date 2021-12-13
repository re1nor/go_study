package main

/*
	Альтернативный формат инструкции "import"
	Импортируем пакет "math", чтобы использовать функцию math.Floor
	Импортируем пакет "strings", чтобы использовать функцию strings.Title
*/
import (
	"fmt"
	"math"
	"strings"
)

func main() {
	/*
		Вызываем функции Floor и Title из пакетов "math" и "strings", соответсвенно
		Функция math.Floor получает число с плавающей точкой, округляет его до ближайшего меньшего целого и возвращает полученное число.
		А функция strings.Title получает строку, преобразует первую букву каждого слова к верхнему регистру и возвращает полученную строку.
	*/
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("bulat zamaletdinov"))
}
