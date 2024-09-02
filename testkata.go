package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Symbol - хранит оператор, line - хранит 2 операнда
var symbol byte
var line []string

// Создание функции, определяющей оператор в строке
func spl(str string) ([]string, byte, error) {

	if line = strings.Split(str, "+"); len(line) == 2 {
		symbol = '+'
		return line, symbol, nil
	}

	if line = strings.Split(str, "-"); len(line) == 2 {
		symbol = '-'
		return line, symbol, nil
	}

	if line = strings.Split(str, "*"); len(line) == 2 {
		symbol = '*'
		return line, symbol, nil
	}

	if line = strings.Split(str, "/"); len(line) == 2 {
		symbol = '/'
		return line, symbol, nil
	}

	return nil, 0, errors.New("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
}

// Проверка, является ли 1 операнд арабской цифрой
func firstNumArab([]string) bool {
	strToInt := map[string]int{
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
	}

	str1 := line[0]

	num1 := strToInt[str1]

	if num1 >= 1 && num1 <= 10 {
		return true
	} else {
		return false
	}
}

//Проверка, является ли 2 операнд арабской цифрой

func secondNumArab([]string) bool {
	strToInt := map[string]int{
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
	}
	str2 := line[1]

	num2 := strToInt[str2]

	if num2 >= 1 && num2 <= 10 {
		return true
	} else {
		return false
	}
}

//Проверка, является ли 1 операнд римской цифрой

func firstNumRome([]string) bool {
	str1 := line[0]
	romeToArab := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	num1 := romeToArab[str1]

	if num1 >= 1 && num1 <= 10 {
		return true
	} else {
		return false
	}

}

//Проверка, является ли 2 операнд римской цифрой

func secondNumRome([]string) bool {
	str2 := line[1]
	romeToArab := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	num2 := romeToArab[str2]

	if num2 >= 1 && num2 <= 10 {
		return true
	} else {
		return false
	}
}

func main() {
	arabToRome := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}

	romeToArab := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	var ex string
	var err error

	for {
		fmt.Println("Введите пример: ")
		fmt.Scanln(&ex)

		if len(ex) < 3 {
			panic("Выдача паники, так как строка не является математической операцией.")
		}

		if line, symbol, err = spl(ex); err != nil {
			panic("Выдача паники, так как строка не является математической операцией.")
		}

		if firstNumArab(line) == true && secondNumRome(line) == true {
			panic("Выдача паники, так как используются одновременно разные системы счисления.")
		}

		if firstNumRome(line) == true && secondNumArab(line) == true {
			panic("Выдача паники, так как используются одновременно разные системы счисления.")
		}

		if firstNumArab(line) == true && secondNumArab(line) == true {
			var res int
			num1, _ := strconv.Atoi(line[0])
			num2, _ := strconv.Atoi(line[1])

			switch symbol {
			case '+':
				res = num1 + num2
			case '-':
				res = num1 - num2
			case '*':
				res = num1 * num2
			case '/':
				res = num1 / num2
			}
			fmt.Println(res)
		}

		if firstNumRome(line) == true && secondNumRome(line) == true {
			var res int
			num1 := romeToArab[line[0]]
			num2 := romeToArab[line[1]]

			switch symbol {
			case '+':
				res = num1 + num2
			case '-':
				res = num1 - num2
			case '*':
				res = num1 * num2
			case '/':
				res = num1 / num2
			}

			if res <= 0 {
				panic("Выдача паники, так как в римской системе нет отрицательных чисел и нуля")
			}
			fmt.Println(arabToRome[res])
		}
	}
}