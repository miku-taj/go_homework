package main
import (
	"fmt"
	"strings"
)

func main(){
	//concatenator := func(s1, s2 string) string { return s1+s2 }
	//println(concatenator("yes", "sir"))
	fmt.Println(Calculate(5, 20, adder))

	multiply_to_five := Multiplier(5)
	fmt.Println(multiply_to_five(3))

}

//1

func PrintGreeting() {
	fmt.Println("Hello, World!" )
}

func DisplaySeparator() {
	fmt.Println("--------------------")
}

func NotifyStart() {
	fmt.Println("Program started")
}

//2

func GetWelcomeMessage() string {
	return "Welcome!"
}

func GetPiValue() float64 {
	return 3.14
}

func IsServerActive() bool {
	return true
}

//3
func PrintNumber(n int) {
	fmt.Println(n)
}

func GreetUser(name string) {
	fmt.Printf("Hello, %s\n", name)
}

func ToggleLight(state bool) {
	if state {
		fmt.Println("Light on")
	} else {
		fmt.Println("Light off")
	}
}

//4
func Add(a, b int) int {
	return a + b
}

func Concat(s1, s2 string) string {
	return s1 + s2
}

func IsEven(n int) bool {
	return n % 2 == 0
}

//5

var adder func(int, int) int = func( a, b int) int { return a + b }
var subtractor func(int, int) int = func( a, b int) int { return a - b }
var concatenator func(string, string) string = func(s1, s2 string) string { return s1 + s2}
var isPositive func(int) bool = func(i int) bool { return i > 0}

//6
func Calculate(a, b int, f func(int, int) int) int{
	return f(a, b)
}
func Execute(value bool, f func(bool)) {
	f(value)
}
func Apply(n int, f func(int) int) int {
	return f(n)
}

//7
func Multiplier(factor int) (func(int) int) {
	return func(n int) int {
		return n * factor
	}
}

func StringRepeater(n int) (func(string) string) {
	return func(s string) string {
		return strings.Repeat(s, n)
	}
}

func ConditionalPrinter(condition bool) (func(string)) {
	return func(s string) {
		if condition {
			fmt.Println(s)
		}
	}
}