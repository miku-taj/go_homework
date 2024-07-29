package main

import "fmt"

func main() {
	var a [9]int
	fmt.Println("Enter ", len(a) ," array elements")
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}

	maxim := a[0]
	minim := a[0]
	positive := 0
	sum := 0

	for i := range a {
		//1
		if a[i] > maxim {
			maxim = a[i]
		}
		//2
		if a[i] < minim {
			minim = a[i]
		}
		//3
		if a[i] > 0 {
			positive++
		}
		//4
		sum += a[i]
	}
	//5
	average := float64(sum) / float64(len(a))
	fmt.Printf("Max el: %d; Min el: %d; Summ of elements: %d; Arithmetic mean is: %g", maxim, minim, sum, average)

	//6
	var delete_el int
	fmt.Println("Enter an element to delete ")
	fmt.Scan(&delete_el) 
	a1 := []int{}
	for i := range a {
		if a[i] != delete_el {
			a1 = append(a1, a[i])
		}
	}
	fmt.Println("Array with deleted element(s): ", a1)

	//7
	var multiplier int
	fmt.Println("Enter a number to multiply the array to")
	fmt.Scan(&multiplier) 
	for i := range a {
		a[i] = a[i] * multiplier
	}
	fmt.Println("Updated array is ", a)

	//8
	var search_number int
	fmt.Println("Enter a number to search by index")
	fmt.Scan(&search_number)
	var indexes []int
	for i := range a {
		if a[i] == search_number {
			indexes = append(indexes, i)
		}
	} 
	fmt.Println("The indexes of searched number are ", indexes)

	//9
	a_copy := make([]int, len(a))
	copy(a_copy, a[:])
	fmt.Println("Copy of a ", a_copy)

	//10
	var a_united []int
	a_united = append(a_united, a[:]...)
	a_united = append(a_united, a1...)
	fmt.Println("2 Arrays: ", a, a1)
	fmt.Println("United arrays: ", a_united)

	//11 если в массиве неск мин и макс элементов, то они все поменяются местами
	//еще можно в 1 цикле найти макс и мин, а в след цикле просто заменить,
	//тогда не придется сохранять номера элементов макс и мин
	maxim = a[0]
	max_id := []int{0}
	minim = a[0]
	min_id := []int{0}
	fmt.Println("array before min/max switch: ", a)
	for i := 1; i < len(a); i++ {
		if a[i] > maxim {
			maxim = a[i]
			max_id = []int{i}
		} else if a[i] == maxim {
			max_id = append(max_id, i)
		}
		if a[i] < minim {
			minim = a[i]
			min_id = []int{i}
		} else if a[i] == minim {
			min_id = append(min_id, i)
		}
	}
	for _, id := range max_id {
		a[id] = minim
	}
	for _, id := range min_id {
		a[id] = maxim
	}
	fmt.Println("array after min/max switch: ", a)

	//12
	is_palindrom := true
	for i := 0; i <= len(a)/2; i++ {
		if a[i] != a[len(a)-1-i] {
			is_palindrom = false
			break
		}
	}
	if is_palindrom {
		fmt.Println("Array is A PALINDROM")
	} else {
		fmt.Println("Array is NOT A PALINDROM")
	}

	//13
	if len(a) < 2 {
		fmt.Println("В массиве меньше 2 чисел, нет второго наибольшего числа")
	} else {
		maxim = a[0]
		sec_maxim := a[1]
		for i := range a {
			if a[i] > maxim {
				sec_maxim = maxim
				maxim = a[i]
			} else if a[i] > sec_maxim && a[i] != maxim {
				sec_maxim = a[i]
			}
		}
	fmt.Println("Second largest element is ", sec_maxim)
	}

	
	//14
	fmt.Println("Array before flipping: ", a)
	for i := 0; i < len(a)/2; i ++ {
		temp := a[i]
		a[i] = a[len(a)-1-i]
		a[len(a)-1-i] = temp
	}
	fmt.Println("Array after flipping: ", a)

	//15
	a_without_duplicates := []int{}
	var is_included bool
	for _, val := range a {
		is_included = false
		//проверим есть ли элемент в масиве без дубликатов
		for _, val2 := range a_without_duplicates {
			if val == val2 {
				is_included = true
				break
			}
		}
		//если элемента в массиве без дубликатов нет, то добавим его
		if !is_included {
			a_without_duplicates = append(a_without_duplicates, val)
		}
	}
	fmt.Println("Array without_duplicates: ", a_without_duplicates)
	
	//16
	zero_num := 0
	a_zero_to_end := []int{}
	for _, val := range a {
		if val != 0 {
			a_zero_to_end = append(a_zero_to_end, val)
		} else {
			zero_num ++
		}
	}
	for i := 0; i < zero_num; i ++ {
		a_zero_to_end = append(a_zero_to_end, 0)
	}
	fmt.Println("Array with zeros moved to the tail: ", a_zero_to_end)

	//17 это без проверки на дубликаты, т.е. если в 1 массиве 2 пятерки
	//а в другом 1, то в итоге мы зачтем 2
	a_intersection := []int{}
	for _, val := range a {
		is_included = false
		//проверим есть ли элемент в масиве без дубликатов
		for _, val2 := range a1 {
			if val == val2 {
				is_included = true
				break
			}
		}
		if is_included {
			a_intersection = append(a_intersection, val)
		}
	}

	//18
	is_subset := true
	for _, val := range a_intersection {
		is_included = false
		for _, val2 := range a {
			if val == val2 {
				is_included = true
				break
			}
		}
		//если хоть один элемент из a_intersection не входит в а, то все, это не подмнож
		//+ выход из цикла
		if !is_included {
			is_subset = false
			break
		}
	}
	fmt.Println("Is subset: ", is_subset)

	//19
	var sorted1, sorted2 []int
	sorted1 = []int{0, 3, 6, 9}
	sorted2 = []int{2, 4, 5, 10}
	lim1 := 0
	lim2 := 0
	sorted_united := make([]int, 0, len(sorted1) + len(sorted2))
	for len(sorted_united) < len(sorted1) + len(sorted2) {
		if lim1 < len(sorted1) {
			if sorted1[lim1] <= sorted2[lim2] {
				sorted_united = append(sorted_united, sorted1[lim1])
				lim1 ++
			}
		}
		if lim1 >= len(sorted1) {
			sorted_united = append(sorted_united, sorted2[lim2:]...)
			break
		}
		if lim2 < len(sorted2) {
			if sorted2[lim2] <= sorted1[lim1] {
				sorted_united = append(sorted_united, sorted2[lim2])
				lim2 ++
			}
		}
		if lim2 >= len(sorted2) {
			sorted_united = append(sorted_united, sorted1[lim1:]...)
			break
		}
	}

	fmt.Println(sorted_united)
	
	//20
	
	subset := []int{}
	longest_subset := []int{}
	for _, val := range a {
		is_included = false
		for _, val2 := range subset {
			if val == val2 {
				is_included = true
				break
			}
		}
		if !is_included {
			subset = append(subset, val)
		} else {
			if len(subset) > len(longest_subset) {
				longest_subset = subset
				subset = []int{val}
			}
		}
	}
	fmt.Println("Longest subset is ", longest_subset)

	//21
	var n int
	fmt.Println("Task 21, enter a number: ")
	fmt.Scan(&n)
	subsets := [][]int{}
	for i := range a {
		sum = 0
		for j := i; j < len(a); j ++ {
			sum += a[j]
			if sum == n {
				subsets = append(subsets, a[i : j+1])
			}
		}
	}
	fmt.Println(subsets)

	//22
	fmt.Scan(&n)
	for i := 0; i < len(a) -1 ; i++ {
		for j := i + 1; j < len(a); j ++ {
			if a[i] + a[j] == n {
				fmt.Println("Элементы номер ", i, " и ", j, "дают сумму ", n)
			}
		}
	}

	//23
	m := 1
	cond := true
	for cond {
		for i, val := range a {
			if val == m {
				m ++
				break
			}
			if i == len(a)-1 {
				cond = false
			}
		}
	}
	fmt.Println("наименьший положительный элемент, отсутствующий в массиве: ", m)

	//25
	fmt.Println("Task 25, enter a number: ")
	fmt.Scan(&n)
	max_len := 0
	for i := range a {
		sum = 0
		for j := i; j < len(a); j ++ {
			sum += a[j]
			if sum == n && len(a[i : j+1]) > max_len {
				max_len = len(a[i : j+1])
			}
		}
	}	
	fmt.Println("максимальная длина подмассива, сумма элементов которого равна заданному числу ", n, ", равна ", max_len)

	//26
	product := a[0] * a[1] * a[2]
	for i := 0; i < len(a) - 2; i++ {
		for j := i + 2; j < len(a); j++ {
			if a[i] * a[i+1] * a[j] > product {
				product = a[i] * a[i+1] * a[j]
			}
		}
	}
	fmt.Println("максимальное произведение трех чисел в массиве: ", product)

	//27
	max_sum := a[0]
	max_sum_slice := []int{a[0]}
	for i := range a {
		sum = 0
		for j := i; j < len(a); j ++ {
			sum += a[j]
			if sum > max_sum {
				max_sum = sum
				max_sum_slice = a[i:j+1]
			}
		}
	}
	fmt.Println("подмассив с максимальной суммой: ", max_sum_slice, " его сумма: ", max_sum)

	//28
	pos_slice := []int{}
	neg_slice := []int{}
	for _, val := range a {
		if val < 0 {
			neg_slice = append(neg_slice, val)
		} else {
			pos_slice = append(pos_slice, val)
		}
	}
	a_neg_to_beg := append(neg_slice, pos_slice...)
	fmt.Println("Массив, где отрицательные числа перемещены в начало массива: ", a_neg_to_beg)

	//29
	longest_zero_subset := []int{}
	for i := range a {
		sum = 0
		for j := i; j < len(a); j ++ {
			sum += a[j]
			if sum == 0 && len(a[i : j+1]) > len(longest_zero_subset){
				longest_zero_subset = a[i : j+1]
			}
		}
	}
	fmt.Println(longest_zero_subset)
	
	//30
	


}