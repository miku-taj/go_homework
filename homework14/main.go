package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)
//1
func countCharacters(fileName string) (int, error) {
	n := 0
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("Error ", err)
		return 0, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan(){
		n++
	}
	return n, nil
}

//2
func countLines(fileName string) (int, error){
	lines_number := 0
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("error encountered", err)
		return 0, err
	}
	scanner:= bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		lines_number ++
	}
	return lines_number, err
}

//3
func countWords(s string) int{
	words_num := 0
	word_started := false
	reader := strings.NewReader(s)
	for {
		r, _, err := reader.ReadRune()
		if unicode.IsLetter(r){
			if !word_started{
				words_num ++
				word_started = true
			}
		} else {
			if word_started{
				word_started = false
			}
		}
		if err == io.EOF{
			break
		}
	}
	return words_num
}

//4
func writeStringToFile(fileName, content string) error{
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	defer file.Close()
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(content))
	return err
}

//5
func readFirstLine(fileName string) (string, error){
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		return scanner.Text(), nil
	}
	return "", scanner.Err()

}

//6
func copyFile(src, dst string) error{
	src_file, err1 := os.Open(src)
	dst_file, err2 := os.OpenFile(dst, os.O_WRONLY|os.O_APPEND, 0666)
	defer src_file.Close()
	defer dst_file.Close()
	if err1!= nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	buf := make([]byte, 2)
	for {
		n, err1 := src_file.Read(buf)
		if err1 == io.EOF {
			break
		}
		_, err2 = dst_file.Write(buf[:n])
		if err1 != nil {
			return err1
		}
		if err2 != nil {
			return err2
		}
	}
	return nil
}

//7
func readAndWriteToFile(fileName string) error{
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	if err := scanner.Err(); err != nil {
		return err
	} 
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		file.Write([]byte(scanner.Text()))
		break
	}
	return nil
}
//8
func reverseReadFile(fileName string) (string, error){
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	buf := make([]byte, 1)

	position := 0
	s := ""
	for {
		position --
		cur, err := file.Seek(int64(position), io.SeekEnd)
		fmt.Println("reading from ", cur)
		_, err1 := file.Read(buf)
		s += string(buf)
		if cur == 0{
			break
		}
		if err != nil {
			return "", err
		}
		if err1 != nil {
			return "", err
		}
		fmt.Println(s)
	}
	return s, nil
}

//9
func concatenateFiles(file1, file2, outputFile string) error{
	f1, err := os.Open(file1)
	if err != nil {
		return err
	}
	f2, err := os.Open(file2)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f1.Close()
	defer f2.Close()
	defer f.Close()
	scanner := bufio.NewScanner(f1)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		_, err := io.WriteString(f, scanner.Text()+"\n")
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	scanner = bufio.NewScanner(f2)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan(){
		_, err := io.WriteString(f, scanner.Text()+"\n")
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

//10
func fileExists(fileName string) bool{
	_, err := os.Open(fileName)
	return !errors.Is(err, os.ErrNotExist)
}

//11
func countUniqueWords(fileName string) (int, error){
	words_map := map[string]int{}
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan(){
		words_map[strings.ToLower(scanner.Text())] ++
		if scanner.Err() != nil {
			return 0, nil
		}
	}
	return len(words_map), nil
}

//12
// func replace



func main(){
	// fmt.Println(countCharacters("text.txt"))
	// fmt.Println(countLines("text.txt"))
	//fmt.Println(countWords("!this is a text,yeah, yeah   wordT,,,.y"))
	//fmt.Println(writeStringToFile("text.txt", "\nanother line"))
	//fmt.Println(readFirstLine("text.txt"))
	//fmt.Println(copyFile("text.txt", "copy_text.txt"))
	//fmt.Println(readAndWriteToFile("copy_text.txt"))
	//fmt.Println(concatenateFiles("text.txt", "copy_text.txt", "concatenate.txt"))
	//fmt.Println(fileExists("text.txt"))
	//fmt.Println(countUniqueWords("text.txt"))
	//fmt.Println(reverseReadFile("text.txt"))
	
}