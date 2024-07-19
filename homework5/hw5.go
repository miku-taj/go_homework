package main

import (
	"fmt"
	"math"
)

func main() {
	
	//1
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
	
	//2
	for i := 1; i < 5; i++ {
		fmt.Println(i*i)
	}
	
	//3
	for i := 1; i <= 10; i++ {
		fmt.Printf("3 * %d = %d\n", i, 3*i)
	}
	
	//4
	a, b := 1, 1
	c := 0
	fmt.Print(a, " ", b)
	for i := 3; i <= 10; i++ {
		c = a + b
		fmt.Print(" ", c)
		a = b
		b = c 
	}
	
	//5
	fmt.Scan(&a, &b)
	var temp int
	if a > b {
		temp = a
		a = b
		b = temp
	}
	for a > 0 {
		b = b % a
		temp = a
		a = b
		b = temp
	}
	fmt.Println("NOD is ", b) 
	
	//6
	for i := 1; i <= 100; i++ {
		if i % 15 == 0 {
			fmt.Print("FizzBuzz ")	
		} else if i % 3 == 0 {
			fmt.Print("Fizz ")
		} else if i % 5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(i, " ")
		}
	}
	
	//7
	fmt.Println("Enter for prime check")

	var num int
	fmt.Scan(&num)
	is_prime := true
	if num == 1 {
		is_prime = false
	}
	for i := 2; i <= int(math.Round(math.Sqrt(float64(num)))); i++ {
		if num % i == 0 {
			is_prime = false
		}
	}
	
	if is_prime {
		fmt.Println("The number is prime")
	} else {
		fmt.Println("the number is not prime")
	}
	
	//8
	fmt.Println("Enter to find all divisors")
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		if num % i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	
	//9
	fmt.Println("Number to calculate digit sum")
	fmt.Scan(&num)
	sum := 0
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	fmt.Printf("The sum is %d\n", sum)
	
	//10
	num = -10
	for num <= 0 {
		fmt.Println("Enter a positive number")
		fmt.Scan(&num)
	}
	
	//11
	fmt.Println("Enter a number to calculate factorial")
	fmt.Scan(&num)
	mult := 1
	for i := 1; i <= num; i++ {
		mult *= i
		if mult > 1000 {
			fmt.Print("The result is more than 1000\n")
			mult = 0
		}
	}
	if mult != 0 {
		fmt.Println("The result is ", mult)
	}

	//12
	fmt.Println("Enter to check if palindrom")
	fmt.Scan(&num)
	var first_digit, last_digit, power int 
	is_palindrom := true
	for num > 0 {
		last_digit = num % 10
		first_digit = num
		power = 1
		for first_digit > 9 {
			first_digit /= 10
			power *= 10
		}
		if last_digit != first_digit {
			is_palindrom = false
			break
		} else {
			num -=last_digit + first_digit * power
			num /= 10
		}
	}
	fmt.Println("Number is palindrom ", is_palindrom)
	
	//13
	fmt.Println("Enter pyramid height")
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
	
	//14
	fmt.Println("Enter for a multiplication table")
	fmt.Scan(&num)
	for i := 1; i <= num; i++ {
		for j:= 1; j < 11; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i * j)
		}
	}
	
	//15
	fmt.Println("Enter Pascal pyramid height! ")
	fmt.Scan(&num)
	
	current_line := []int {1}
	var temp_slice []int
	for i:=1; i < num; i ++ {
		for _, value := range current_line {
			fmt.Print(value, " ")
		}
		fmt.Print("\n")
		temp_slice = current_line
		current_line = nil
		current_line = append(current_line, temp_slice[0])
		for j:=1; j < len(temp_slice); j ++ {
			current_line = append(current_line, temp_slice[j-1]+temp_slice[j])
		}
		current_line = append(current_line, temp_slice[len(temp_slice) - 1])
	}


	//16
	num = 1
	fmt.Println("Enter a number")
	fmt.Scan(&num)
	for num >= 0 {
		mult = 1
		for i:= 1; i <= num; i++ {
			mult *= i
		}
		fmt.Println("The factorial is ", mult)
		fmt.Println("Enter a number")
		fmt.Scan(&num)
		}

	fmt.Println("Finally a negative number!")
	
	//17
	//for true {
	//	fmt.Println("Enter 2 numbers")
	//	fmt.Scan(&a, &b)
	//	fmt.Println(a + b)
	//}
	
	//18
	fmt.Println("1 -100 without perfect squares")
	for i := 0; i < 101; i ++ {
		if int(math.Sqrt(float64(i)))*int(math.Sqrt(float64(i))) == i {
			continue
		} else {
			fmt.Print(i, ", ")
		}
	}
	fmt.Print("\nPrime numbers from 1 to 50\n")
	
	//19
	for i := 1; i < 51; i++ {
		is_prime = true
		if i == 1 {
			is_prime = false
		}
		for j := 2; j <= int(math.Round(math.Sqrt(float64(i)))); j++ {
			if i % j == 0 {
				is_prime = false
			}
		}
		if is_prime {
			fmt.Print(i, ", ")
		}
	}
	fmt.Print("\nNumbers from 1 to 30 that aren't divisable by 2 or 3")
	
	//20
	for i := 1; i < 31; i++ {
		if i % 2 == 0 || i % 3 == 0 {
			continue
		} else {
			fmt.Print(i, ", ")
		}
	}
	fmt.Print("\n")
	
	//21
	for i := 1; i < 51; i++ {
		fmt.Print(i, " ")
		if int(math.Cbrt(float64(i)))*int(math.Cbrt(float64(i)))*int(math.Cbrt(float64(i))) == i {
			fmt.Println("Came across a perfect cube!")
			break
		}
	}
	
	//22
	sum = 0
	for i := 1; i < 101; i++ {
		fmt.Print(i, ", ")
		sum += i
		if sum > 200 {
			fmt.Print("sum is bigger than 200")
			break
		}
	}
	fmt.Print("\n")
	
	//23
	num = 1
	for num % 7 != 0 {
		fmt.Println("Enter a number")
		fmt.Scan(&num)
	}
	
}
