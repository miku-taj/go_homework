package main
import (
	"fmt"
	"strings"
)

func main(){
	//concatenator := func(s1, s2 string) string { return s1+s2 }
	//println(concatenator("yes", "sir"))
	//1
	PrintGreeting()
	DisplaySeparator()
	NotifyStart()

	//2
	message := GetWelcomeMessage()
	fmt.Println(message)

	pi := GetPiValue()
	fmt.Println(pi)

	on_off := IsServerActive()
	fmt.Println(on_off)

	//3
	PrintNumber(49)
	GreetUser("Mark")
	ToggleLight(true)
	fmt.Println(Calculate(5, 20, adder))

	//4
	a := Add(5, 9)
	fmt.Println(a)

	s := Concat("hel", "lo")
	fmt.Println(s)

	is_odd := !IsEven(a)
	fmt.Println(is_odd)

	//5
	b := adder(3, 2)
	fmt.Println(b)

	s2 := concatenator("mes", "sage")
	fmt.Println(s2)

	info := isPositive(-32)
	fmt.Println(info)
	
	//6
	c := Calculate(10, -5, adder)
	fmt.Println(c)

	Execute(true, ToggleLight)

	d:= Apply(32, Multiplier(5))
	fmt.Println(d)

	//7
	multiply_to_five := Multiplier(5)
	fmt.Println(multiply_to_five(3))

	repeat_3times := StringRepeater(3)
	fmt.Println(repeat_3times("repeated! "))

	printer := ConditionalPrinter(true)
	printer("Conditional printer worked!")
}

//1

func PrintGreeting() {
	fmt.Println("Hello, World!" )
}

func DisplaySeparator() {
	fmt.Println(strings.Repeat("-", 20))
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