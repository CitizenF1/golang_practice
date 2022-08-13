package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func numJewelsInStones(jewels string, stones string) int {
	cnt := 0
	jwl := make(map[byte]bool, len(jewels))
	for i := 0; i < len(jewels); i++ {
		jwl[jewels[i]] = true
	}
	for i := 0; i < len(stones); i++ {
		if _, isExt := jwl[stones[i]]; isExt {
			cnt++
		}
	}
	return cnt
}

func average(salary []int) float64 {
	t, m, M, l := 0, math.MaxInt, math.MinInt, 0
	for _, s := range salary {
		t += s
		if s < m {
			m = s
		} // found new min
		if M < s {
			M = s
		} // found new Max
		l++
	}
	// fmt.Println(t, m, M, l)
	return float64(t-(m+M)) / float64(l-2)
}

func shuffle(nums []int, n int) []int {
	var ans = make([]int, len(nums))
	j := 0
	for i := 0; i < len(nums)/2; i++ {
		ans[j] = nums[i]
		ans[j+1] = nums[i+len(nums)/2]
		j += 2
	}
	return ans
}

func sortSentence(s string) string {
	words := strings.Split(s, " ")
	res := make([]string, len(words))

	for i := 0; i < len(res); i++ {
		fmt.Println(i)
		res[words[i][len(words[i])-1]-'1'] = words[i][:len(words[i])-1]
		fmt.Println(res)
	}
	// fmt.Println(res)

	return ""
}

func minPartitions(n string) int {
	res := 0
	for _, v := range n {
		ch, _ := strconv.Atoi(string(v))
		if ch > res {
			res = ch
		}
	}
	return res
}

func minPartitionsSol(n string) int {

	max := 0

	for i, _ := range n {
		if max < int(n[i]-'0') {
			max = int(n[i] - '0')
		}
	}

	return max
}

func finalValueAfterOperations(operations []string) int {
	count := 0
	for _, v := range operations {
		if v == "++X" || v == "X++" {
			count++
		}
		if v == "--X" || v == "X--" {
			count--
		}
	}
	return count
}

type NumArray struct {
}

func Constructor(nums []int) NumArray {
	var sum []int

	current := 0
	for _, v := range nums {
		current += v
		sum = append(sum, current)
	}
	return NumArray{}
}

// func (this *NumArray) SumRange(left int, right int) int {

// }

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func getConcatenation(nums []int) []int {
	ans := nums

	ans = append(ans, nums...)

	return ans
}

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

func romanToInt(s string) int {
	var result int
	// roman := map[string]int{
	// 	"I": 1,
	// 	"V": 5,
	// 	"X": 10,
	// 	"L": 50,
	// 	"C": 100,
	// 	"D": 500,
	// 	"M": 1000,
	// }

	return result
}

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func buildArray(nums []int) []int {
	var ans []int
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[nums[i]], "i", i, nums[i])

		ans = append(ans, nums[nums[i]])
	}
	return ans
}

func containsDuplicate(nums []int) bool {
	fmt.Println(nums)
	sort.Ints(nums)
	fmt.Println(nums, "after")

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			fmt.Println(nums[i+1], "i+1", nums[i])
			return true
		}
	}
	return false
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	resultlist := ListNode{}

	// for _, v := range list1 {
	// 	for _, v := range list2 {

	// 	}
	// }

	return &resultlist
}

func lengthOfLastWord(s string) int {
	length := 0

	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(i, "+")
		if string(s[i]) == " " {
			if length == 0 {
				continue
			}
			break
		}
		length++
		fmt.Println(length, "++")
	}
	return length
}

func maximumWealth(accounts [][]int) int {
	var resul int
	for _, v := range accounts {
		tmp := 0
		for _, e := range v {
			tmp += e
			if resul > tmp {
				resul = tmp
			}
		}
	}
	return resul
}

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
