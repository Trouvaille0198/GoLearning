package stdpkglearning

import (
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

// generate 生成随机数组
func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

// generate2 生成随机数组
func generate2(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = rand.Intn(n)
	}
	return nums
}

// bubbleSort 冒泡排序
func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func getChar(nums []int) (res string) {
	for _, num := range nums {
		res += strconv.Itoa(num)
	}
	return
}

func getChar2(nums []int) string {
	var builder strings.Builder
	for _, num := range nums {
		builder.WriteString(strconv.Itoa(num))
	}
	return builder.String()
}

func Run() {
	// f, _ := os.OpenFile("cpu2.pprof", os.O_CREATE|os.O_RDWR, 0644)
	// defer f.Close()
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	n := 10000
	for i := 0; i < 50; i++ {
		nums := generate2(n)
		bubbleSort(nums)
		_ = getChar2(nums)
	}
}

func RunWithGoroutine() {
	// f, _ := os.OpenFile("cpu4.pprof", os.O_CREATE|os.O_RDWR, 0644)
	// defer f.Close()
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()

	n := 10000
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			nums := generate2(n)
			bubbleSort(nums)
			_ = getChar2(nums)
			wg.Done()
		}()
	}
	wg.Wait()
}
