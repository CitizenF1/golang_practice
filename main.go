package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type Storage struct {
	m    *sync.Mutex
	root string
}

func CreateStorage() *Storage {
	root := "./storage"
	os.Mkdir(root, 0755)
	return &Storage{
		m:    &sync.Mutex{},
		root: root,
	}
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

func main() {
	isPolindrome(123)
	///==========BYTE request
	// data, err := os.ReadFile("./image.png")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // fmt.Println(data)

	// req, err := http.NewRequest("POST", "http://localhost:7360", bytes.NewBuffer(data))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer resp.Body.Close()
	// io.Copy(os.Stdout, resp.Body)
	//====================
	// do not modify below here, readline is our function
	// that properly reads in the input for you

	//lettercount
	// fmt.Println(LetterCount("Hello apple pie"))
	// fmt.Printf("%v\n", StringMerge("abc1*kyoo"))
	// fmt.Println(StringMerge("abc1*kyoo"))

}

func BinarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		// fmt.Println(mid, "mid")
		guess := list[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0
}

// func main() {

// client, err := rpc.DialHTTP("tcp", "192.168.0.76"+":7350")
// if err != nil {
// 	log.Fatal("dialing:", err)
// }
// var reply int
// err = client.Call("http_handler_path", "{some : data}", &reply)
// if err != nil {
// 	log.Fatal("arith error:", err)
// }
// fmt.Print("Arith: %d*%d=%d", reply)

// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
// body := strings.NewReader(`"{"some": "data"}"`)
// req, err := http.NewRequest("POST", "http://192.168.0.76:7350/v2/rpc/http_handler_path?http_key=defaulthttpkey", body)
// if err != nil {
// 	// handle err
// 	fmt.Println(err, "req")
// }
// req.Header.Set("Content-Type", "application/json")
// req.Header.Set("Accept", "application/json")

// resp, err := http.DefaultClient.Do(req)
// if err != nil {
// 	// handle err
// 	fmt.Println(err, "resp")
// }
// defer resp.Body.Close()

// socket, err := socket.Dial(context.Background(), "tcp", "localhost:8081")
// if err != nil {
// 	fmt.Println(err)
// }
// fmt.Println(socket.LocalAddr())
// client := &http.Client{}
// // Perform the request	// conn, err := net.Dial("tcp", "192.168.0.159:7351")
// if err != nil {
// 	fmt.Println(err)
// 	return
// }

// defer conn.Close()
// io.Copy(os.Stdout, conn)
// fmt.Println("\nDone")0", "text/plain", bytes.NewBufferString("Hello Go!"))
// if err != nil {
// 	log.Fatalf("Failed get: %s", err)
// }
// defer resp.Body.Close()
//==============================================
// c := evtwebsocket.Conn{
// 	OnConnected: func(w *websocket.Conn) {
// 		fmt.Println("Connected")
// 	},
// 	OnMessage: func(msg []byte) {
// 		fmt.Printf("Received message: %s\n", msg)
// 	},
// 	OnError: func(err error) {
// 		fmt.Printf("** ERROR **\n%s\n", err.Error())
// 	},
// }
// // Connect
// c.Dial("http://localhost:7350")
// c.Send([]byte("TEST"), nil)
//===================================================

// }

//========================================================================
// func main() {
// RemoveFolder("test5")
// fmt.Println(os.Getwd())

// for i := 0; i < 10; i++ {
// 	for k := 0; k < 9-i; k++ {
// 		fmt.Print("-")
// 	}
// 	for starts := 0; starts < 2*i+1; starts++ {
// 		fmt.Print("*")
// 	}
// 	fmt.Println("-")
// }

// CreateStorage().MoveFile("test4", ".storage/", "test3")
// RemoveFile("storage/test1", "file.txt")
// CreateFile(storage, "fl.txt")
// fmt.Println(listFolder(("./storage/")))

// err := filepath.Walk("/file.txt",
// 	func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		fmt.Println(path)
// 		return err
// 	})
// if err != nil {
// 	log.Println(err)
// }

// f, err := os.Open("./storage")
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// files, err := f.Readdir(0)
// if err != nil {
// 	fmt.Println(err)
// 	return
// }

// for _, v := range files {
// 	fmt.Println(v.Name(), v.IsDir())
// }
// dir, err := os.Getwd()
// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println(dir)
// fmt.Println(FilePathWalkDir("./"))
// fmt.Println(OSReadDir("./storage/test1/file.txt"))

// }

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func OSReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

func listFolder(path string) (string, error) {
	var files string
	fileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for i := range fileInfo {
		files += fileInfo[i].Name() + " " + strconv.FormatBool(fileInfo[i].IsDir()) + "\n"
	}
	return files, err
}

func CreateFile(location string, fileName string) error {
	file, err := os.Create(location + "/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	return nil
}

func (this *Storage) MoveFile(fileName string, oldPath string, newPath string) error {
	err := os.Rename(filepath.Join(oldPath, fileName), filepath.Join(newPath, fileName))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// func twoSum(nums []int, target int) []int {
// 	mapNums := map[int]int{}
// 	for i := range nums {
// 		dif := target - nums[i]
// 		val, isExist := mapNums[dif]
// 		if isExist == true {
// 			return []int{val, i}
// 		} else {
// 			mapNums[nums[i]] = 1
// 		}
// 	}
// 	return nil
// }

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

// }

//LEGASY CODE

// url := "https://www.google.com/"
// fmt.Printf("HTML code of %s ...\n", url)
// resp, err := http.Get(url)
// // handle the error if there is one
// if err != nil {
// 	panic(err)
// }
// // do this now so it won't be forgotten
// defer resp.Body.Close()
// // reads html as a slice of bytes
// html, err := ioutil.ReadAll(resp.Body)
// if err != nil {
// 	panic(err)
// }
// // show the HTML code as a string %s
// fmt.Printf("%s\n", html)
