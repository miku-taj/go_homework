package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1 PrintGreeting
	PrintGreeting()
	//2 PrintWeather
	PrintWeather()
	//3 PrintTrafficLight
	PrintTrafficLight()
	//4 GetGrade
	fmt.Println(GetGrade())
	//5 GetDiscount
	fmt.Println(GetDiscount())
	//6 GetTemperatureDescription
	fmt.Println(GetTemperatureDescription())
	//7 CheckNumber
	CheckNumber(-143)
	//8 CheckAge
	CheckAge(17)
	//9 CheckPassword
	CheckPassword("2341")
	//10 Add
	fmt.Println(Add(-5, -20))
	//11 CompareStrings
	fmt.Println(CompareStrings("hey", "Hey"))
	//12 Max
	fmt.Println(Max(40, -20))
	//13 operation
	fmt.Println(operation(-24, 15))
	//14 concat
	fmt.Println(concat("Good", "morning"))
	//16

	//17
	CheckCondition(21, number_check)
	//18
	FormatAndPrint("this is not an empty string", formatting_func)
}

func PrintGreeting() {
	var dayType string
	fmt.Println("Пожалуйста, введите dayType")
	fmt.Scan(&dayType)
	if dayType == "утро" {
		fmt.Println("Доброе утро!")
	} else if dayType == "день" {
		fmt.Println("Добрый день!")
	} else if dayType == "вечер" {
		fmt.Println("Добрый вечер!")
	}
}

func PrintWeather(){
	var weatherType  string
	fmt.Println("Пожалуйста, введите weatherType")
	fmt.Scan(&weatherType )
	if weatherType  == "солнечно" {
		fmt.Println("Солнечно")
	} else if weatherType  == "облачно" {
		fmt.Println("Облачно")
	} else if weatherType == "дождливо" {
		fmt.Println("Дождливо!")
	}
}

func PrintTrafficLight() {
	var lightColor string
	fmt.Println("Пожалуйста, введите lightColor")
	fmt.Scan(&lightColor)
	if lightColor == "lightColor" {
		fmt.Println("Стоп")
	} else if lightColor == "желтый" {
		fmt.Println("Внимание")
	} else if lightColor == "зеленый" {
		fmt.Println("Идите")
	} else {
		fmt.Println("Цвет не зарегистрирован")
	}
}

func GetGrade() string{
	var scope float64
	result := "Введен невозможный балл"
	fmt.Println("Пожалуйста, введите ваш балл")
	fmt.Scan(&scope)
	if scope >= 0 && scope < 75 {
		result = "C"
	} else if scope >= 75 && scope < 90 {
		result = "B"
	} else if scope >= 90 && scope <= 100 {
		result = "A"
	}
	return result
}

func GetDiscount() string {
	var amount int
	result := "0%"
	fmt.Println("Please enter the amount")
	fmt.Scan(&amount)
	if amount > 1000 { 
		result = "10%"
	} else if amount > 500 && amount <= 1000 {
		result = "5%"
	}
	return result
}

func GetTemperatureDescription() string {
	var temperature float64
	result := ""
	fmt.Println("Enter temperature")
	fmt.Scan(&temperature)
	if temperature < 10 {
		result = "Cold"
	} else if temperature < 25 {
		result = "Mild"
	} else {
		result = "Hot"
	}
	return result
}

func CheckNumber(n int) {
	if n < 0 {
		fmt.Println("Negative")
	} else if n == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Positive")
	}
}

func CheckAge(age int) {
	if age < 18 {
		fmt.Println("Несовершеннолетний")
	} else {
		fmt.Println("Совершеннолетний")
	}
}

func CheckPassword(password string) {
	if password == "1234" {
		fmt.Println("Пароль верный")
	} else {
		fmt.Println("Пароль неверный")
	}
}

func Add(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a + b
}

func CompareStrings(s1, s2 string) string {
	if s1 == s2 {
		return "равны"
	} else {
		return "неравны"
	}
}

func Max(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func operation(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

var concat func(string, string) string = func(s1, s2 string) string {
	if s1 == "" || s2 == "" {
		return s1 + s2
	} else {
		return s1 + " " + s2
	}
}

var multiply func(int, int) int = func(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a * b
}

func ApplyOperation(a, b int, f func(n1, n2 int) int) int{
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return f(a, b)
}

//17
func CheckCondition(n int, f func(int) bool) {
	if f(n) {
		fmt.Println("Условие выполнено")
	} else {
		fmt.Println("Условие не выполнено")
	}
}

func number_check(n int) bool {
	if n % 3 == 0 {
		return true
	}
	return false
}

//18
func FormatAndPrint(s string, f func(string) string)  {
	fmt.Println(f(s))
}

func formatting_func(s string) string {
	if s == "" {
		return "Пустая строк"
	} else {
		return strings.ToUpper(s)
	}
}

//19
func CreateMultiplier(n int) (func(int) int) {
	if n < 0 {
			n = -n
		}
	return func(i int) int {
		return i * n
	}
}

//20
func CreateGreeter(s string) (func(string) string) {
	return func(name string) string {
		if s == "" {
			return "Hello, " + name
		} else {
			return s + ", " + name
		}
	}
}

//21
func CreateValidator(password string) (func(string) bool) {
	return func(s string) bool {
		if s == password {
			return true
		} else {
			return false
		}
	}
}
