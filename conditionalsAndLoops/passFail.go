/* passFail сообщает, сдал ли пользователь экзамен. Пользователь вводит процент правильных ответов и узнает, прошел он экзамен или нет.
Результат: при значении 60% и выше экзамен успешно сдан, а при значении ниже 60% - провален. */
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Введите значение:")
	// Создаем "буферизованное средство чтения" для получения текста с клавиатуры
	reader := bufio.NewReader(os.Stdin)
	// Возвращает текст, введенный пользователем до нажатия клавиши Enter и сохраняет возвращаемое значение ошибки в переменной err
	input, err := reader.ReadString('\n')
	if err != nil {
		// Сообщает об ошибке и прерывает программу
		log.Fatal(err)
	}

	// Удаляет символ новой строки из введенных данных
	input = strings.TrimSpace(input)
	// Преобразовывает введенную строку в значение float64
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "Экзамен сдан"
	} else {
		status = "Экзамен провален"
	}

	fmt.Println("Ваш результат = ", grade, ".", status)
}
