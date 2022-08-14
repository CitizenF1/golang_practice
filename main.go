package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"sync"
	"time"
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

// range specification, note that min <= max
type IntRange struct {
	min, max int
}

// get next random value within the interval including min and max
func (ir *IntRange) NextRandom(r *rand.Rand) int {
	return r.Intn(ir.max-ir.min+1) + ir.min
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func RangeInt(n int) []string {
	var res []string
	// rr := map[int]string{}
	var done []string

	theBoys := []string{
		"rakhotep",
		"Randomllll",
		"Magauiya00",
		"drinksu",
		"Tima_b",
		"ganikazbek",
		"MarkusDosmagambetov",
		"Sakenever",
		"askarbai1",
		"azimkhankuat",
		"Baglan2",
		"rayodelsol01",
		"AzamatBK4",
		"ykabylbekov",
		"adilbekov23",
		"nurik_yes",
		"yerlawkaa",
		"CTZNF1",
		"OPTIMUS_GANK",
		"Bakbergen_M",
		"Deus_loVult7",
		"jc1an",
		"runreyvo",
		"Kaleeeeeeb",
		"mrsdrd",
	}

	var r int
	for r = 0; r <= n-1; r++ {
		// arr[r] = rand.Intn(max) + min
		done = append(done, theBoys[rand.Intn(len(theBoys))])
	}
	res = removeDuplicateStr(done)
	// result := getNonVoted(theBoys, res)
	// fmt.Println(removeDuplicateStr(done))

	// rand.Seed(time.Now().UnixNano())
	// p := rand.Perm(len(theBoys))
	// for _, r := range p[:2] {
	// 	fmt.Println(theBoys[r])
	// }
	return res
}

func doneboys(list1, list2 []string) string {
	var nonDuty []string
	for _, t := range list1 {
		voted := false
		for _, v := range list2 {
			if t == v {
				voted = true
				break
			}
		}
		if !voted {
			nonDuty = append(nonDuty, t)
		}
	}
	result := ""
	for _, n := range nonDuty {
		result += "|" + n + " "
	}
	return result
}

func restoreString(s string, indices []int) string {
	ans := make([]rune, len(s))
	for i, v := range s {
		ans[indices[i]] = v
	}
	return string(ans)
}

func main() {
	// fmt.Println(piscine.TrimAtoi("sdx1-fa2W3s4"))
	// var indices = []int{4, 5, 6, 7, 0, 2, 1, 3}
	// s := "codeleet"
	// restoreString(s, indices)
	// 	start := time.Date(2013, time.November, 17, 0, 0, 0, 0, time.UTC)
	// 	end := start.AddDate(9, 0, 0)
	// 	fmt.Println(start.Format("2006-01-02 15:04:05"), "-", end.Format("2006-01-02 15:04:05"))

	// 	for rd := rangeDate(start, end); ; {
	// 		date := rd()
	// 		if date.IsZero() {
	// 			break
	// 		}
	// 		rand.Seed(time.Now().UnixNano())
	// 		fmt.Println(date.Format("2006-01-02"))
	// 	}

	// 	[]
	// [2,7,4]

	seat := []int{3, 1, 5}
	students := []int{2, 7, 4}
	minMovesToSeat(seat, students)
}

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	var sum int
	for i := 0; i < len(seats); i++ {
		// fmt.Println(Abs(seats[i]-students[i]), "2")
		sum += Abs(seats[i] - students[i])
	}
	return sum
}

func Abs(i int) int {
	// fmt.Println(i, "1")
	if i < 0 {
		// fmt.Println(-i, "--")
		return -i
	}
	// fmt.Println(i, "++")
	return i
}

func randFloats(min, max float64) float64 {
	res := min + rand.Float64()*(max-min)
	return res
}

func rangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

func recaman(size int) []int {
	vec := []int{}
	vec[0] = 0
	vec[1] = 1

	for i := 2; i < size; i++ {
		if !elementExists(vec, i) && vec[i-1]-i > 0 {
			vec[i] = vec[i-1] - i
		} else {
			vec[i] = vec[i-1] + i
		}
	}
	return vec
}

func elementExists(vec []int, elem int) bool {
	for i := 0; i < len(vec); i++ {
		if vec[i] == elem {
			return true
		}
	}
	return false
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

func twoSum(nums []int, target int) []int {
	mapNums := map[int]int{}
	for i := range nums {
		dif := target - nums[i]
		val, isExist := mapNums[dif]
		if isExist == true {
			return []int{val, i}
		} else {
			mapNums[nums[i]] = 1
		}
	}
	return nil
}

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
