package main

import (
	"fmt"
	"strings"
	"math"
	"math/rand/v2"
	"time"
)

var alphabet = [...] string {"a", "b", "c", "d",
 "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o",
"p", "q", "r", "s", "t", "w", "x", "y", "z"}

func main() {
	
	//t := time.Now() 
	//for _, i in range
	//time = time.String()
	
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	//fmt.Printf("Go launched at %s\n", t.Local())

	fmt.Printf("%T\n", t)
	//fmt.Println(GenerateInvoiceNumber(t))

	//PrintFormattedDate()

	//users := [...]int{return_three_values()}
	var users []string = make([]string, 10)
	fmt.Println(users)

	str_slice := []string{"hey", "this", "would", "be", "for", "bbq"}
	PrintFilteredStrings(str_slice, has_b)

	//CalculateStatistics
	int_nums := []int{4, 5, 21, 2, 10, 5, 14, 7}
	CalculateStatistics(int_nums)

	//FindMax
	float_nums := []float64{2, 21, 10.2, 56.7, 20}
	fmt.Println(FindMax(float_nums))

	//FilterEvenNumbers
	fmt.Println(FilterEvenNumbers(int_nums))

	//calculator 
	fmt.Println(calculator(3, 10))

	//stringManipulator
	s1 := "Hey, my name is george h"
	fmt.Println(stringManipulator(s1))
}

func GenerateRandomPassword() string {
	
	password := ""
	var is_digit, to_add_digit int
	var to_add_letter string

	 for i := 0; i < 10; i++{
        is_digit = rand.IntN(2)
		if is_digit == 1 {
			to_add_digit = rand.IntN(10)
			password = fmt.Sprintf(password+"%d", to_add_digit) 
		} else {
			to_add_letter = alphabet[rand.IntN(len(alphabet) - 1)]
			password = fmt.Sprintf(password+"%s", to_add_letter) 
		}
    }
	return password
}

func ReverseString(str string) string{
	final_str := ""
	for i := 0; i < len(str); i ++ {
		final_str = final_str + str[len(str) - i - 1:len(str) - i]
	}
	return final_str
}

func CountWords(str string) int {
	words_num := 0
	if str == "" {
		return words_num
	} else {
		words_num = 1
		for i := 0; i < len(str); i ++ {
			if str[i:i+1] == " " {
				words_num += 1
			}
		}
		return words_num
	}
}

func PrintMultiplicationTable(num int) {
	for i:=1; i<11; i++ {
		fmt.Printf("%d × %d = %d\n", num, i, num*i)
	}
}

func GenerateInvoiceNumber (t time.Time) string {
	inv_name := "INV-"
	number := rand.IntN(1000)
	inv_name = fmt.Sprintf(inv_name + "%d-%02d-%03d",
	t.Year(), t.Month(), number)
	return inv_name
}

func PrintFormattedDate() {
	
	fmt.Printf("%d %s %d", time.Now().Day(), time.Now().Month().String(), time.Now().Year())
		
}

func return_three_values() (int, int, int) {
	return 1, 2, 3
}

func PrintFilteredStrings(strings []string, f func(s string)bool) {
	for _, value := range strings {
		if f(value) {
			fmt.Println(value)	
		}
	}
	
} 

func has_b(s string) bool {
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "b" {
			return true
			}
	}
	return false
}

func CalculateStatistics(numbers []int) {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	mean := float64(sum) / float64(len(numbers))
	fmt.Printf("sum is %d, arith mean is %g\n", sum, math.Round(mean*100)/100)
}

func FindMax(numbers []float64) float64 {
	max:= numbers[0]
	for _, value := range numbers {
		if value > max {
			max = value
		}
	}
	return max
}

func FormatName(name string, surname string) string {
	formatted_name := fmt.Sprintf("%s, %s", surname, name)
	return formatted_name
}

func FilterEvenNumbers(numbers[]int) []int {
	result := []int{}
	for _, value := range numbers {
		if value % 2 == 0 {
			result = append(result, value)
		}
	}
	return result
}

var calculator func(float64, float64) string = func(a, b float64) string {
	result := fmt.Sprintf("%g + %g = %g\n%g - %g = %g\n%g × %g = %g\n",
	a, b, a + b, a, b, a - b, a, b, a*b)
	if b!= 0 {
		result = fmt.Sprintf(result + "%g / %g = %g\n", a, b, a/b)
	} else {
		result = fmt.Sprintf(result + "Деление невозможно\n")
	}
	return result
}

func stringManipulator(s string) string {
	s = strings.ToUpper(string(s[0])) + s[1:]
	for i := range s {
		if string(s[i]) == " " {
			s = s[:i+1]+strings.ToUpper(string(s[i+1])) + s[i+2:]
		}
	}
	return s
	//	return strings.Title(s)
}

func filter(numbers []int, f func(int) bool) []int {
	result := []int {}
	for _, value := range numbers {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func ProcessNumbers(numbers []int, f func(int) int) {
	for _, value := range numbers {
		fmt.Println(f(value))
	}	
}

func ApplyToStrings(ss[]string, f func(string) string) []string{
	result := []string{}
	for _, value := range ss {
		result = append(result, f(value))
	}
	return result
}

func FilterAndPrint(numbers []float64, f func(float64) bool) {
	for _, value := range numbers {
		if f(value) {
			fmt.Println(value)
		}
	}
}

func CreateMultiplier(factor int) (func(int) int) {
	return func(i int) int {
		return i * factor
	}
}

func CreateFormatter(s string) (func(string) string) {

	if s == strings.ToTitle(s) {
		return func(ss string) string {
			return strings.ToTitle(ss)
		}
	} else if s == strings.ToLower(s) {
		return func(ss string) string {
			return strings.ToLower(ss)
		}
	} else if s == strings.ToUpper(s) {
		return func(ss string) string {
			return strings.ToUpper(ss)
		}
	} else if strings.ToUpper(string(s[0]))+strings.ToLower(s[1:]) == s {
		return func(ss string) string {
			return strings.ToUpper(string(ss[0]))+strings.ToLower(ss[1:])
		}
	} else {
		return func(ss string) string {
			return "Format unidentified"
		}
	}

}

