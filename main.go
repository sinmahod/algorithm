package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//fmt.Println(GCD(10, 3))
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
