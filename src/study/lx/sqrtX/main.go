package main

import (
	"fmt"
)

// 求数x的平方根。

// 方法一：二分法。
func binarySearch(x int) {
	value := 0
	if x == 0 {
		fmt.Println(value)
		return
	}
	upLimit := x
	lowerLimit := 0
	for {
		if (upLimit+lowerLimit)/2*(upLimit+lowerLimit)/2 > x {
			upLimit = (upLimit + lowerLimit) / 2
		} else if (upLimit+lowerLimit)/2*(upLimit+lowerLimit)/2 < x {
			if ((upLimit+lowerLimit)/2+1)*((upLimit+lowerLimit)/2+1) >= x {
				if x-(upLimit+lowerLimit)/2*(upLimit+lowerLimit)/2 > ((upLimit+lowerLimit)/2+1)*((upLimit+lowerLimit)/2+1)-x {
					value = (upLimit+lowerLimit)/2 + 1
				} else {
					value = (upLimit + lowerLimit) / 2
				}
				break
			}
			if (upLimit+lowerLimit)/2 == lowerLimit {
				value = upLimit
				break
			}
			lowerLimit = (upLimit + lowerLimit) / 2
		} else {
			value = (upLimit + lowerLimit) / 2
		}
	}
	fmt.Println(value)
}

// 方法二：牛顿迭代。
// n*n=x，n/x=n。
// 16可以表示为2x8,4x4，2和8的平均值比2或者8都更接近于16的平方根4。
func binarySearch2(x int) {
	sqrt(x, x)
}
func sqrt(i, x int) {
	res := (i + x/i) / 2
	if res == i {
		fmt.Println(i)
	} else {
		sqrt(res, x)
	}
}

func main() {
	binarySearch(50)
	//binarySearch2(50)
}
