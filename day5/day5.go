package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isPolymer(a, b byte) bool {
	if unicode.IsUpper(rune(a)) && unicode.IsLower(rune(b)) {
		return unicode.ToUpper(rune(b)) == rune(a)
	} else if unicode.IsLower(rune(a)) && unicode.IsUpper(rune(b)) {
		return unicode.ToLower(rune(b)) == rune(a)
	}
	return false
}
func reduce1(charList []byte) (int, []byte) {
	i := 0
	j := 1
	for j < len(charList) {
		if isPolymer(charList[j], charList[i]) {
			copy(charList[i:], charList[j+1:])
			charList = charList[:len(charList)-2]
			if i > 0 {
				i--
				j--
			}
		} else {
			i++
			j++
		}
	}
	return len(charList), charList
}
func main() {
	fh, err := os.Open("./day5Puzzle.txt")
	if err != nil {
		fmt.Println(err)
	}
	var charList []byte
	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		charList = scanner.Bytes()
	}
	// part 1
	length, newline2 := reduce1(charList)
	fmt.Println(length)

	// part 2
	var arr = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	line := string(newline2)
	minLength := 1000000
	for _, v := range arr {
		newLine := strings.Replace(line, v, "", -1)
		newLine = strings.Replace(newLine, strings.ToLower(v), "", -1)
		newLength, _ := reduce1([]byte(newLine))
		if newLength < minLength {
			minLength = newLength
		}
	}
	fmt.Println(minLength)

}
