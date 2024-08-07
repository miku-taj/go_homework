package main

import (
	"fmt"
	"strings"
	"strconv"
	"math/rand/v2"
)

func main() {
	s := "golagonggo"
	s1 := "go"
	s2 := "hey"
	//fmt.Println(substringIndex(s, s1))
	fmt.Println(replaceSubstring(s, s1, s2))
	s = " hey ! "
	fmt.Println(trimWhitespace(s)+"lol")
	s = "abca&mdash;adfc"
	fmt.Println(DeleteChar(s, "a"))
	fmt.Println(DeleteDublicates(s))
	fmt.Println(replaceHTML(s))

	v := "3 31 22 "
	fmt.Println(SumNumbers(v))

	s = "moley is name my"
	fmt.Println(ReverseWords(s))
	s = "hey my hey hey car is new, i just just bought it"
	fmt.Println(CountUnique(s))

	s = "moley is name my avada harry lover stone dad unfortunately died"
	fmt.Println(MixWords(s))


}

// 1
func ConcateStrings(s1, s2 string) string {
	return s1 + s2
}

// 2
func StringLength(s string) int {
	/*	n := 0
		for i := range s{
			n ++
		}
		return n
	*/
	return len(s)
}

//function returns a slice of indexes where substring starts
func detectSubstring(main_s, sub_s string) []int{
	if len(sub_s) > len(main_s){
		return []int{}
	}
	indexes := []int{}
	present := false
	for i := 0; i < len(main_s); i++{
		present = true
		for k := 0; k < len(sub_s); k++{
			if i+k >= len(main_s){
				present = false
				break
			}
			if main_s[i+k:i+k+1] != sub_s[k:k+1]{
				present = false
				break
			}
		} 
		if present {
			indexes = append(indexes, i)
		}
	}
	return indexes
}


// 3
func isSubstring(main_s, sub_s string) bool {
/*or:
	return strings.Contains(main_s, sub_s)

*/	
	k := 0
	for i := 0; i < len(main_s); i ++ {
		if main_s[i] == sub_s[k]{	
			k++
		} else {
			k = 0
		}
		if k == len(sub_s) {
			return true
		}
	}
	return false
}

//4
func substringIndex(main_s, sub_s string) int{
	//using strings:
	//return strings.Index(main_s, sub_s)
	//or:
	indexes := detectSubstring(main_s, sub_s)
	if indexes == nil {
		return -1
	} else {
		return indexes[0]
	}
}

//5
func replaceSubstring(main_s, sub_delete, sub_insert string) string{
	//usings strings:
	//return strings.Replace(main_s, sub_delete, sub_insert, -1)
	//or
	//return strings.ReplaceAll(main_s, sub_delete, sub_insert)
	//or my method:
	indexes := detectSubstring(main_s, sub_delete)
	if indexes == nil {
		return main_s
	}
	result_string := main_s[0:indexes[0]]
	for i := 0; i < len(indexes); i++{
		if i != len(indexes) - 1{
			result_string += sub_insert + main_s[indexes[i] + len(sub_delete): indexes[i+1]]
		} else {
			result_string += sub_insert + main_s[indexes[i] + len(sub_delete):]
		}
		
	}
	return result_string
}

//6
func trimWhitespace(s string) string{
	s = strings.TrimPrefix(s, " ")
	s = strings.TrimSuffix(s, " ")
	return s
}

//7
func changeCase(s string) []string{
	return []string{strings.ToLower(s), strings.ToUpper(s)}
}

//8
func repeatString(s string, n int) string{
	return strings.Repeat(s, n)
}

//9
func splitString(s string, sep string) []string{
	return strings.Split(s, sep)
}

//10
func CompareNoCase(a, b string) int{
	return strings.Compare(strings.ToLower(a), strings.ToLower(b))
}

//11
func ReplaceFirst(s, old, new string) string{
	return strings.Replace(s, old, new, 1)
}

//12
func InvertString(s string) string{
	result := ""
	for _, char := range s {
		result = string(char) + result
	}
	return result
}

//13
func CountChar(s string) map[string]int{
	char_map := map[string]int{}
	for _, char := range s{
		char_map[string(char)] += 1
	}
	return char_map
}

//14
func DeleteChar(s, char string) string{
	return strings.ReplaceAll(s, char, "")
}

//15
func WordCount(s string) int{
	words := strings.Split(s, " ")
	return len(words)
}

//16
func startsPrefix(s, prefix string) bool{
	return strings.HasPrefix(s, prefix)
}
func endsSuffix(s, suffix string) bool{
	return strings.HasSuffix(s, suffix)
}

//17
func DeleteDublicates(s string) string{
	for i, char := range s{
		if i < len(s) - 1 {
			s = s[0:i+1] + strings.ReplaceAll(s[i+1:], string(char), "")
		}
	}
	return s
}

//18
var specialHTML map[string]string = map[string]string{
	"&nbsp;": " ",
	"&ndash;": "–",
	"&mdash;": "—",
	"&bull;": "•",
	"&copy;": "©",
	"&lt;": "<",
	"&gt;": ">",
	"&amp;": "&",
	"&para;": "¶",
}
func replaceHTML(s string) string{
	for i, char := range specialHTML{
		s = strings.ReplaceAll(s, i, char)
	}
	return s
}

//19

//20
func SumNumbers(s string) int {
	sum := 0
	number := 0
	for i := range s {
		num, err := strconv.Atoi(s[i:i+1])
		if err != nil {
			sum +=number
			number = 0
		} else {
			number = number*10 + num
			fmt.Println("digit is ", num, " number: ", number)
		}
	}
	return sum + number
}

//21
func ReverseWords(s string) string{
	words:= strings.Split(s, " ")
	final := ""
	for i := len(words) - 1; i >=0; i--{
		final += words[i] + " "
	}
	return strings.TrimSuffix(final, " ")
}

//22
func CountUnique(s string) int{
	n := 0
	words:= strings.Split(s, " ")
	for _, word := range words{
		if strings.Count(s, word) > 0 {
			n += 1
			s = strings.ReplaceAll(s, word, "")
		}
	}
	return n
}

//23


//24
func MixWords(s string) string{
	words := strings.Split(s, " ")
	final_str := ""
	id := -1
	for len(words) > 1 {
		id = rand.IntN(len(words))
		final_str += words[id] + " "
		words = append(words[:id], words[id+1:]...)
	}
	final_str += words[0]
	return final_str
}

//25
func SortByLength(s string) string{
	final := ""
	words := strings.Split(s, " ")
	temp := ""
	k:= 0
	for i := 1; i < len(words); i ++{
		k = i - 1
		for k >= 0 {
			if len(words[k]) > len(words[k+1]){
				temp = words[k]
				words[k] = words[k+1]
				words[k+1] = temp
				k --
			} else {
				break
			}
		}
	}
	for _, el := range words{
		final += el + " "
	}
	return strings.TrimSuffix(final, " ")
}

