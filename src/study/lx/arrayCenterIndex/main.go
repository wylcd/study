package main

import "fmt"

// 数组中心索引。
// 给定一个整数数组nums，给定一个返回数组中心下标的方法。
// 中心下标是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。如果数组不存在中心下标，则返回-1。数组有多个中心下标，返回靠左边那个。

// 思路1：左边+中心值=右边+中心值，把数组中的每个值都假设为中心值，将这个值与左边和右边相加的和作比较，如果相等，就表明这个值是中心值。
func arrayCenterIndex() {
	//nums := [...]int{1, 1, 1, 1, 1, 1, 1, 6}
	//nums := [...]int{1, 7, 3, 6, 5, 6}
	//nums := [...]int{6, -6, 0}
	nums := [...]int{0, 6, -6}
	// 数组和
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 对应值
	val := 0
	// 13 1, 12 2, 11 3, 10 4, 9 5, 8 6, 7 7
	// 28 1, 21 8, 20 11, 17 17
	// 中心值
	centerVal := -1
	for i := 0; i < len(nums); i++ {
		val += nums[i]
		if val == sum {
			centerVal = i
			break
		}
		sum -= nums[i]
	}
	fmt.Println(centerVal)
}

// 思路2：数组和=2x左边+中心值 或 2x右边+中心值，假设数组每个值都是中心值，将中心值左边的和乘以2加上中心值，看是否等于数组和。
func arrayCenterIndex2() {
	//nums := [...]int{1, 1, 1, 1, 1, 1, 1, 6}
	//nums := [...]int{1, 7, 3, 6, 5, 6}
	//nums := [...]int{6, -6, 0}
	nums := [...]int{0, 6, -6}
	// 数组和
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum == nums[0] {
		fmt.Println(0)
		return
	}
	for i := 1; i < len(nums); i++ {
		leftVal := 0
		for j := 0; j < i; j++ {
			leftVal += nums[j]
		}
		if leftVal*2+nums[i] == sum {
			fmt.Println(i)
			break
		}
	}
}

func main() {
	arrayCenterIndex()
	arrayCenterIndex2()
}
