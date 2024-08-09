package main

import (
	"fmt"
	"strings"
)

func main() {

	//1
	str_length_map := make(map[string]int)
	str_length_map["No"] = 2
	fmt.Println(str_length_map)

	//2 Создайте map с несколькими элементами
	// напишите функцию, которая проверяет наличие заданного ключа.
	HarryStylesMap := map[int]string{
		2014: "Night Changes",
		2017: "Sign of the Times",
		2019: "Watermelon High",
		2022: "As it was",
		2023: "Watermelon High",
	}
	DoesExist(HarryStylesMap, 2025)
	//3
	UpdateElement(HarryStylesMap, 2022, "Late Night Talking")
	s := "i met him near a car station . A cop pulled me over and i had problems with documents when he stepped in and helped"
	fmt.Println(WordFrequncy(s))
	//6
	fmt.Println(SortKeys(HarryStylesMap))
	//7
	fmt.Println("Is HarryStylesMap empty? ", isEmpty(HarryStylesMap))
	//8
	fmt.Printf("%#v\n", InvertKeysValues(HarryStylesMap))
	//9
	fmt.Println(CheckDublicates(HarryStylesMap))
	//10
	fmt.Printf("HarryStylesMap Values: %#v\n", GetValues(HarryStylesMap))
	//11
	s_slice := []string{"a", "a", "a", "b", "b", "c"}
	fmt.Println("Number of unique el:", CountUnique(s_slice))

	//20
	new_map := map[string]string{
		"running": "fast",
		"walking": "slow",
		"standing": "zero",
	}
	f := func(s string)string {
		return strings.ReplaceAll(s, "", " ")
	}
	UpdateMapFunc(new_map, f)
	fmt.Println(new_map)

	//23
	maps := []map[string]int{
		{"a": 1,},
		{"c": 3,},
		{"a": 2, "b": 2,},
	}
	fmt.Println(UniteSeveralMaps(maps))

	//25
	User1 := User{Name: "Bloom", Address: "Alphea Dorms, room 1"}
	pointer_map := make(map[string]*User)
	pointer_map["user1"] = &User1


	//30
	FilmmakersMap := map[string][]string{
		"Steven Spielberg": {"The Color Purple", "Schindler's List", "Jurrasic Park"},
		"Christopher Nolan": {"Oppenheimer", "Dunkierk", "Interstellar", "The Dark Knight"},
	}
	AddSlice(FilmmakersMap, "James Cameron", []string{"Avatar", "Titanic"})
	DeleteSlice(FilmmakersMap, "James Cameron")
	fmt.Println(GetSlice(FilmmakersMap, "Steven Spielberg"))
}

//2
func DoesExist(m map[int]string, key int) bool{
	_, exists := m[key]
	if !exists {
		fmt.Println("Element with this key is not found")
	}
	fmt.Println("Element exists")
	return exists
}

//3
func UpdateElement(m map[int]string, key int, new_value string) bool{
	_, exists := m[key]
	if !exists{
		fmt.Println("Element with this key is not found")
		return false
	}
	m[key] = new_value
	return true
}

//4
func DeleteElement(m map[int]string, key int) bool{
	_, exists := m[key]
	if !exists{
		fmt.Println("Element with this key is not found")
		return false
	}
	delete(m, key)
	return true
}

//5
func WordFrequncy(s string) map[string]int{
	frequency_map := map[string]int{}
	words := strings.Split(s, " ")

	for _, word := range words{
		frequency_map[word] ++
	}
	return frequency_map
}

//6
func SortKeys(m map[int]string) []int{
	sorted_keys := []int{}
	k:= 0
	temp := 0
	for key := range m{
		k = len(sorted_keys) - 1
		sorted_keys = append(sorted_keys, key)
		for k >= 0 {
			if sorted_keys[k] > sorted_keys[k+1]{
				temp = sorted_keys[k]
				sorted_keys[k] = sorted_keys[k+1]
				sorted_keys[k+1] = temp
				k --
			} else {
				break
			}
		}
	}
	return sorted_keys
}

//7
func isEmpty(m map[int]string) bool{
	if len(m) == 0 {
		return true
	}
	return false
}

//8
func InvertKeysValues(m map[int]string) map[string]int{
	new_m := make(map[string]int)
	for key, element := range m{
		new_m[element] = key
	}
	return new_m
}

//9
func CheckDublicates(m map[int]string) bool{
	for key, el := range m {
		for key1, el1 := range m{
			if key1 != key && el1 == el{
				return true
			}
		}
	}
	return false
}

//10
func GetValues(m map[int]string) []string{
	values := []string{}
	for _, el := range m{
		values = append(values, el)
	}
	return values
}

//11
func CountUnique(s []string) int{
	unique := 0
	m := map[string]int{}
	for _, str := range s{
		if m[str] == 0{
			unique ++
		}
		m[str] = 1
	}
	return unique
}

//12
func FindKeysFunc(m map[int]string, f func(string) bool) []int{
	keys := []int{}
	for key, el := range m{
		if f(el){
			keys = append(keys, key)
		}
	}
	return keys
}

//13
func UniteMaps(inferior_map, superior_map map[int]string) map[int]string{
	united_map := map[int]string{}
	for key, el := range inferior_map{
		united_map[key] = el
	}
	for key, el := range superior_map{
		united_map[key] = el
	}
	return united_map
}

//14
func FindKeyByValue(m map[string]string, value string) string{
	for key, el := range m{
		if el == value {
			return key
		}
	}
	return ""
}

//15
func FindIntersection(m1, m2 map[string]string) map[string]string{
	m := map[string]string{}
	for key, el := range m1{
		if el1, exists := m2[key]; exists && el1 == el{
			m[key] = el
		}
	}
	return m
}

//16
func GetKeys(m map[int]string) []int{
	keys := []int{}
	for key := range m{
		keys = append(keys, key)
	}
	return keys
}

//17
func ValuesToString(m map[string]string) string{
	s := ""
	for _, el := range m{
		s += string(el) + ", "
	}
	return strings.TrimSuffix(s, ", ")
}

//18
func CopyMap(m map[string]string) map[string]string{
	new_m := map[string]string{}
	for key, el := range m{
		new_m[key] = el
	}
	return new_m
}

//19
func GetKeyValuePairs(m map[string]string) [][2]string{
	result := [][2]string{}
	for key, el := range m{
		result = append(result, [2]string{key, el})
	}
	return result
}

//20
func UpdateMapFunc(m map[string]string, f func(string)string){
	for key, el := range m{
		m[key] = f(el)
	}
}

//21

//22
func FilterMap(m map[string]string, f_key func(string)bool, f_val func(string)bool){
	for key, el := range m{
		if !f_key(key) || !f_val(el){
			delete(m, key)
		}
	}
}

//23
func UniteSeveralMaps(maps []map[string]int) map[string]int{
	united_map := map[string]int{}
	for _, m := range maps{
		for key, el := range m{
			united_map[key]+= el
		}
	}
	return united_map
} 

//24
func ModifyMapFunc(m map[string]string, f func(string)string){
	for key, val := range m {
		m[f(key)] = f(val)
		delete(m, key)
	}
}

//25
type User struct{
	Name string
	Address string
	Email string
}

//27 Реализуйте метод для структуры, который работает с map в этой структуре.
type MovieCollection struct{
	Movies map[int]string
	Domain string
}
func (m MovieCollection) GetMovies() []string{
	movie_names := []string{}
	for _, name := range m.Movies{
		movie_names = append(movie_names, name)
	}
	return movie_names
}

//30
func AddSlice(m map[string][]string, key string, val []string) bool{
	if _, exists := m[key]; exists{
		return false
	}
	m[key] = val
	return true
}
func DeleteSlice(m map[string][]string, key string) bool{
	if _, exists := m[key]; !exists{
		return false
	}
	delete(m, key)
	return true
}
func GetSlice(m map[string][]string, key string) []string{
	return m[key]
}