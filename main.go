package main

import "fmt"

func main() {
	//fmt.Println(GCD(10, 3))
	fmt.Println(BinarySearch(5, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
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
//BinarySearch(9, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 0,1,2,3,4,5,6,7,8,9
//                5,6,7,8,9   = 5  (10-0)/2
//                         8,9   = 8  (10-6)/2+6
//                            9   = 9  (10-9)/2+9
//BinarySearch(0, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 0,1,2,3,4,5,6,7,8,9
// 0,1,2,3,4,5               = 5   (10-0)/2
// 0,1,2                        = 2   (5-0)/2
// 0,1                           = 1   (2-0)/2
// 0                              = 0   (1-0)/2
//BinarySearch(5, []int{0, 1, 2, 3, 4, 6, 7, 8, 9}
// 0,1,2,3,4,6,7,8,9
//             4,6,7,8,9   = 4  (9-0)/2
//                      8,9   = 7  (9-5)/2+5
//                   7         = 6  (9-7)/2+5
//                6            = 5  (6-5)/2+5
//                              = -1  s=e
func BinarySearch(key int, a []int) int {

	var si = 0
	var ei = len(a)
	//si与ei分别对应Start Index和End Index，会随着计算缩小间隔

	for si < ei {
		var mid = (si + ei) >> 1 //同si + (ei-si)/2，右移一位等于/2
		if key > a[mid] {
			si = mid + 1 //+1的作用是中间位置偏右一位，因为ei等于a的长度而si最多是数组的长度，+1是为了所查询的数字超过数组最右值（最大）时让si=ei成立
		} else if key < a[mid] {
			ei = mid
		} else {
			return mid
		}
	}
	return -1
}
