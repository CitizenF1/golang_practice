package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//ZigZog Convert
func ZigZagConvert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	len := len(s)
	result := make([]byte, len)
	inter := numRows*2 - 2
	index := 0
	for i := 0; i < numRows; i++ {
		for j := 0; i+j < len; j += inter {
			result[index] = s[j+i]
			index++
			if i != 0 && i != numRows-1 && j+inter-i < len {
				result[index] = s[j+inter-i]
				index++
			}
		}
	}
	return string(result)
}

func IsDecimalPalindromeNumber(n int) bool {
	if n < 0 {
		n = -n
	}
	s := strconv.Itoa(n)
	bound := (len(s) / 2) + 1
	for i := 0; i < bound; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func isPolindrome(x int) bool {
	if x < 0 {
		x = -x
	}
	s := strconv.Itoa(x)
	rev := []rune(s)
	for i, j := 0, len(rev)-1; i < len(rev)/2; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	intstr, _ := strconv.Atoi(string(rev))

	return x == intstr

	// bound := (len(s) / 2) + 1
	// fmt.Println(bound)
}

func Reverse(s string) string {
	rev := []rune(s)
	for i, j := 0, len(rev)-1; i < len(rev)/2; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return string(rev)
}

func StringChalleng(num int) string {
	hours := num / 60
	mins := num % 60
	out := strconv.Itoa(hours) + ":" + strconv.Itoa(mins)
	token := "02fqhv4i9ac"
	revOut := Reverse(out)
	revTok := Reverse(token)

	// code goes here
	return revOut + ":" + revTok
}

func StringChallenge(str string) string {
	if len(str) == 0 {
		return ""
	}
	if !strings.Contains(str, "*") {
		return ""
	}
	item := strings.Split(str, "*")
	if len(item) != 2 {
		return ""
	}
	res := ""
	for i := 0; i < len(item[0]); i++ {
		res += string(item[0][i]) + string(item[1][i])
	}
	revRes := Reverse(res)
	token := "02fqhv4i9ac"
	revTok := Reverse(token)
	// code goes here
	return revRes + ":" + revTok

}

func StrChall(num int) string {
	hours := num / 60
	mins := num % 60
	out := strconv.Itoa(hours) + ":" + strconv.Itoa(mins)
	token := "9zwf5rap14b"
	revOut := Reverse(out)
	revTok := Reverse(token)
	return revOut + ":" + revTok
}

func LetterCount(str string) string {
	if len(str) == 0 {
		return "-1"
	}
	words := strings.Split(str, " ")
	kvOut := make(map[int]int, len(words))
	breakDown := true
	for key, value := range words {
		kvIn := make(map[string]int)
		for i := 0; i < len(value); i++ {
			if val, isExist := kvIn[string(value[i])]; isExist {
				kvIn[string(value[i])] = val + 1
			} else {
				kvIn[string(value[i])] = 1
			}
		}
		max := 0
		for _, value := range kvIn {
			if value > max {
				max = value
			}
		}
		kvOut[key] = max
		if breakDown && max != 1 {
			breakDown = false
		}
	}
	if breakDown {
		return "-1"
	}
	var indexSort []int
	for key := range kvOut {
		indexSort = append(indexSort, key)
	}
	sort.Slice(indexSort, func(i, j int) bool {
		return indexSort[i] < indexSort[j]
	})
	max := 0
	keyIndex := 0
	for key := range indexSort {
		if kvOut[key] > max {
			max = kvOut[key]
			keyIndex = key
		}
	}
	maxCount := 0
	for _, value := range kvOut {
		if value == max {
			maxCount++
		}
	}
	if maxCount == 1 {
		return words[keyIndex]
	} else {
		for key := range indexSort {
			if kvOut[key] == max {
				max = kvOut[key]
				keyIndex = key
				break
			}
		}
		return words[keyIndex]
	}
}

func StringMerge(str string) string {
	if len(str) == 0 {
		return ""
	}
	if !strings.Contains(str, "*") {
		return ""
	}
	items := strings.Split(str, "*")
	if len(items) != 2 {
		return ""
	}
	res := ""
	for i := 0; i < len(items[0]); i++ {
		res += string(items[0][i]) + string(items[1][i])
	}
	token := "9zwf5rap14b"
	fmt.Println(res)
	for _, val := range res {
		// op := 1
		for _, el := range token {
			if string(val) == string(el) {
				result := el
				fmt.Println(string(result))
			}
		}
	}
	// for _, char := range res {
	// 	op := 1
	// 	for _, el := range token {
	// 		if string(el) == string(char) {
	// 			op = 0
	// 			break
	// 		}
	// 		if op == 1 {
	// 			res += string(char)
	// 			break
	// 		}
	// 		if len(res) == 0 {
	// 			return "EMPTY"
	// 		}
	// 		fmt.Println(string(el))
	// 	}
	// }
	return res
}
