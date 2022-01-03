// guess - игра, в которой игрок должен угадать задуманное число от 1 до 100.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Получаем текущую дату и время в формате целого числа
	seconds := time.Now().Unix()
	// Инициализируем генератор случайных чисел
	rand.Seed(seconds)
	// Генерируем случайное число от 1 до 100
	target := rand.Intn(100) + 1
	fmt.Println("Я загадал случайное число от 1 до 100. Попробуешь угадать?")

	// Создаем bufio.Reader для чтения ввода с клавиатуры
	reader := bufio.NewReader(os.Stdin)

	// По умолчанию будет выводиться сообщение о проигрыше
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("Осталось попыток: ", 10-guesses)
		fmt.Print("Ваше предположение: ")
		// Читаем данные, введеные пользователем до нажатия Enter
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// Удаление символа новой строки
		input = strings.TrimSpace(input)
		// Входная строка преобразуется в целое число
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Упс. Ваше предположение МЕНЬШЕ загаданного числа")
		} else if guess > target {
			fmt.Println("Упс. Ваше предположение БОЛЬШЕ загаданного числа")
		} else {
			success = true
			fmt.Println("Ура! Вы угадали загаданное число!")
			break
		}
	}
	// Если переменная "success" равна false, сообщить игроку загаданное число
	if !success {
		fmt.Println("Ты не угадал загаданное число :( Это было число:", target)
	}
}
