package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	//fmt.Println(GCD(10, 3))
	//fmt.Println(BinarySearch(5, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
	arr := make([]int, 0)
	for i := 0; i < 100; i++ {
		arr = append(arr, rand.Intn(100))
	}
	s := time.Now()
	fmt.Println(arr)
	fmt.Println("------------------------------------------", len(arr))
	ShellSort(arr)
	fmt.Println(arr)
	e := time.Since(s).Nanoseconds()
	fmt.Printf("%f", float64(e)/1000000)
	key, _ := strconv.Atoi(os.Args[1])
	fmt.Println(BinarySearch(key, []int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12}))
	fmt.Println("=============================")
	fmt.Println(BinarySearch2(key, []int{1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12}))
}

//欧几里德算法求最小公约数:两数相除取余数，余数为0则被除数是最小公约数，不为0则被除数在除以余数继续算，直到余数为0
func GCD(x, y int) int {
	if y == 0 {
		return x
	}
	r := x % y
	return GCD(y, r)
}

//二分查找算法：查询key在数组中的位置（数组必须是有序的）
func BinarySearch2(key int, a []int) int {

	var si = 0
	var ei = len(a) - 1
	//si与ei分别对应Start Index和End Index，会随着计算缩小间隔

	for si <= ei {
		var mid = (si + ei) >> 1 //同si + (ei-si)/2，右移一位等于/2
		if key > a[mid] {
			si = mid + 1 //+1是为了避免表达式(ei-si)的值小于2时出现的死循环，这样至少每次都能上涨一位
		} else if key < a[mid] {
			ei = mid - 1 //-1是为了避免key为元素比最小值还小时造成无限与数组第一位元素做判断的死循环
		} else {
			return mid
		}
		fmt.Println(mid, si, ei)
	}
	return -1
}

//比较a > b
func Compare(a, b int) bool {
	return a > b
}

//希尔排序
func ShellSort(arr []int) []int {
	l := len(arr)
	h := 1
	for h < l/3 {
		h = 3*h + 1 //1，4，13，40，，，，增量为3倍+1
	}
	//这时h为接近l且不大于l的最大增量值
	for h > 0 {
		//假设排序长度为100的数组，则需要外层循环四次
		//第一次循环增量值为40，则依次比较40>0,41>1....99>59>19
		//第二次循环增量值为13，依次计算13>0,14>1...99>86>73>60>47>34>21>8
		//第三次循环增加值为4，依次计算4>0,5>1.....99>95>91>87.....>19>15>11>7>3
		//最后一次循环增量值为1
		for i := h; i < l; i++ {
			for j := i; j >= h && !Compare(arr[j] > arr[j-h]); j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j] //交换数组值
			}
		}
		h = h / 3 //增量每次递减
	}
	return arr
}
