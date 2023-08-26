// Специально для Ката-академии, код написан 26.08.2023
package main

import (
	"fmt"
	"strconv"
)

func ConvertArabtoInt(numberOne, operatorS, numberTwo string) string {

	var resultPlusOperator int
	var result1 string
	var result11 int
	var result22 int

	operator := operatorS
	number1 := numberOne
	number2 := numberTwo

	mapByLiteralArabic := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10, "XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20, "XXI": 21, "XXIV": 24, "XXV": 25, "XXVII": 27, "XXX": 30, "XXXII": 32, "XXXV": 35, "XXXVI": 36, "XL": 40, "XLII": 42, "XLIV": 45, "XLVIII": 48, "XLIX": 49, "L": 50, "LVI": 56, "LX": 60, "LXIII": 63, "LXIV": 64, "LXX": 70, "LXXII": 72, "LXXX": 80, "LXXXI": 81, "XC": 90, "C": 100}
	mapByIntToArabic := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X", 11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX", 21: "XXI", 24: "XXIV", 25: "XXV", 27: "XXVII", 30: "XXX", 32: "XXXII", 35: "XXXVI", 36: "XXXVI", 40: "XL", 42: "XLII", 45: "XLIV", 48: "XLVIII", 49: "XLIX", 50: "L", 56: "LVI", 60: "LX", 63: "LXIII", 64: "LXIV", 70: "LXX", 72: "LXXII", 80: "LXXX", 81: "LXXXI", 90: "XC", 100: "C"}

	for i := range mapByLiteralArabic {
		if number1 == i {
			resultNumberOne := mapByLiteralArabic[i]
			result11 = resultNumberOne

		}
	}

	for j := range mapByLiteralArabic {
		if number2 == j {
			resultNumberSecond := mapByLiteralArabic[j]
			result22 = resultNumberSecond

		}
	}

	switch operator {
	case "+":
		resultPlusOperator = result11 + result22
	case "-":
		resultPlusOperator = result11 - result22
	case "*":
		resultPlusOperator = result11 * result22
	case "/":
		resultPlusOperator = result11 / result22
	default:
		panic("Используйте только +, -, *, /")
	}
	if resultPlusOperator < 1 {
		panic("В римской системе ответ не может быть меньше 1")
	}

	for i := range mapByIntToArabic {
		if resultPlusOperator == i {
			resultNumberOne := mapByIntToArabic[i]
			result1 = resultNumberOne

		}
	}

	return result1

}

func main() {
	var operator string
	var number1 string
	var number2 string
	var result int
	var ok bool
	var ok2 bool
	var okArabic bool
	var result1 string
	var result2 string

	mapByIntToArabic := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X", 11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX", 21: "XXI", 24: "XXIV", 25: "XXV", 27: "XXVII", 30: "XXX", 32: "XXXII", 35: "XXXVI", 36: "XXXVI", 40: "XL", 42: "XLII", 45: "XLIV", 48: "XLVIII", 49: "XLIX", 50: "L", 56: "LVI", 60: "LX", 63: "LXIII", 64: "LXIV", 70: "LXX", 72: "LXXII", 80: "LXXX", 81: "LXXXI", 90: "XC", 100: "C"}

	fmt.Println("Введите операцию:")

	_, err := fmt.Scanln(&number1, &operator, &number2)
	if err != nil {
		panic("Использование 1-ого или более 2-ух чисел не поддерживается / Вы ввели строку, которая не является арифметической операцией")
	}

	for i := range mapByIntToArabic {
		if number1 == mapByIntToArabic[i] {
			resultNumberOne := mapByIntToArabic[i]
			result1 = resultNumberOne
			ok = true
		} else {
			ok = false
		}

		for j := range mapByIntToArabic {
			if number2 == mapByIntToArabic[j] {
				resultNumberOne1 := mapByIntToArabic[j]
				result2 = resultNumberOne1
				ok2 = true
			} else {
				ok2 = false
			}

		}
	}

	if number1 != result1 && number2 == result2 {
		panic("Используются одновременно разные системы счисления")

	} else if number1 == result1 && number2 != result2 {
		panic("Используются одновременно разные системы счисления")

	} else if number1 != result1 && number2 != result2 {
		okArabic = true
	}

	if ok == false && ok2 == false && (number1 == "I" || number1 == "II" || number1 == "III" || number1 == "IV" || number1 == "V" || number1 == "VI" || number1 == "VII" || number1 == "VIII" || number1 == "IX" || number1 == "X") && (number2 == "I" || number2 == "II" || number2 == "III" || number2 == "IV" || number2 == "V" || number2 == "VI" || number2 == "VII" || number2 == "VIII" || number2 == "IX" || number2 == "X") {
		fmt.Println(ConvertArabtoInt(number1, operator, number2))
	} else if okArabic == true && (number1 == "1" || number1 == "2" || number1 == "3" || number1 == "4" || number1 == "5" || number1 == "6" || number1 == "7" || number1 == "8" || number1 == "9" || number1 == "10") && (number2 == "1" || number2 == "2" || number2 == "3" || number2 == "4" || number2 == "5" || number2 == "6" || number2 == "7" || number2 == "8" || number2 == "9" || number2 == "10") {
		number1ToInt, _ := strconv.Atoi(number1)
		number2ToInt, _ := strconv.Atoi(number2)
		switch operator {
		case "+":
			result = number1ToInt + number2ToInt
		case "-":
			result = number1ToInt - number2ToInt
		case "*":
			result = number1ToInt * number2ToInt
		case "/":
			result = number1ToInt / number2ToInt
		default:
			panic(fmt.Sprintf("Используемый оператор '%s' не существует", operator))

		}
		fmt.Print(result)
	} else {
		panic("Используйте числа от 1 до 10")
	}
}
