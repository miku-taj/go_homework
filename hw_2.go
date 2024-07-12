package main
import "fmt"
import "math"

var Pi float64 = 3.14

func main() {

		//1
	var a, P float64
	fmt.Println("Введите длину стороны: ")
	fmt.Scanln(&a)
	P = 4 * a
	fmt.Printf("Периметр равен: %g\n", P)

	//2
	var S float64
	S = a * a
	fmt.Printf("Площадь равна: %g\n", S)

	//3
	var side1, side2 float64
	fmt.Println("Введите стороны прямоугольника:")
	fmt.Scanln(&side1, &side2)
	S = side1 * side2
	P = 2 * (side1 + side2)
	fmt.Printf("Площадь: %g Периметр: %g\n", S, P)

	//4
	var d, L float64
	fmt.Println("Введите диаметр окружности:")
	fmt.Scanln(&d)
	L = Pi * d
	fmt.Printf("Длина окружности : %g\n", L)

	//5
	var cube_side, V float64
	fmt.Println("Введите сторону куба: ")
	fmt.Scanln(&cube_side)
	V = math.Pow(cube_side, 3)
	S = 6 * math.Pow(cube_side, 2)
	fmt.Printf("Объем равен %g Площадь поверхности равна %g\n", V, S)

	//6
	var xside, yside, zside float64
	fmt.Println("Введите стороны прямоугольного параллелепипеда")
	fmt.Scanln(&xside, &yside, &zside)
	V = xside * yside * zside
	S = 2 * (xside * yside + yside * zside + xside * zside)
	fmt.Printf("Объем равен %g Площадь поверхности равна %g\n", V, S)

	//7
	var R float64
	fmt.Println("Введите радиус окружности:")
	fmt.Scanln(&R)
	L = 2 * Pi * R
	S = Pi * R * R
	fmt.Printf("Длина окружности : %g, Площадь: %g\n", L, S)

	//8
	var num1, num2, arith_mean float64
	fmt.Println("Введите два числа:")
	fmt.Scanln(&num1, &num2)
	arith_mean = (num1 + num2)/2
	fmt.Printf("Их среднее арифметическое: %g\n", arith_mean)

	//9
	var geom_mean float64
	geom_mean = math.Sqrt(num1 * num2)
	fmt.Printf("Их среднее геометрическое: %g\n", geom_mean)

	//10
	var number1, number2, sq_sum, sq_dif, sq_prod, sq_quo float64
	fmt.Println("Введите два ненулевых числа:")
	fmt.Scanln(&number1, &number2)
	sq_sum = math.Pow(number1, 2) + math.Pow(number2, 2)
	sq_dif = math.Pow(number1, 2) - math.Pow(number2, 2)
	sq_prod = math.Pow(number1, 2) * math.Pow(number2, 2)
	sq_quo = math.Pow(number1, 2) / math.Pow(number2, 2)
	fmt.Printf("сумма их квадратов: %g\nразность их квадратов: %g\nпроизведение их квадратов: %g\nчастное их квадратов: %g\n", sq_sum, sq_dif, sq_prod, sq_quo)

	//11
	var length float64
	var full_meters int
	fmt.Println("Введите длину в см: ")
	fmt.Scan(&length)
	full_meters = int(length) / 100
	fmt.Printf("Полных метров: %d\n", full_meters)

	//12
	var mass float64
	var full_tons int
	fmt.Println("Введите массу в кг: ")
	fmt.Scan(&mass)
	full_tons = int(mass) / 1000
	fmt.Printf("Полных тонн: %d\n", full_tons)

	//13
	var bytes float64
	var full_kb int
	fmt.Println("Введите размер в байтах: ")
	fmt.Scan(&bytes)
	full_kb = int(bytes) / 1024
	fmt.Printf("Полных килобайтов: %d\n", full_kb)

	//14, 15
	var seg1, seg2, fits, remainder int
	fmt.Println("Введите длины двух отрезков (два целых, положительных числа): ")
	fmt.Scanln(&seg1, &seg2)
	fits = seg1 / seg2
	remainder = seg1 % seg2
	fmt.Printf("Число наложений отрезка%d на отрезок%d: %d\nНезанятая часть первого отрезка: %d\n", seg1,seg2 ,fits, remainder)
	





}
