package main

import "fmt"

// 一个有序数组，原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 删除数组中重复值，且不使用额外数组。
var nums = make([]int, 7, 15)

// 直接使用切割获取。
func delArrayRepeat() {
	nums = []int{0, 1, 2, 2, 3, 3, 4}
	x := nums[0]
	for i := 1; i < len(nums); i++ {
		if x == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			continue
		}
		x = nums[i]
	}
	fmt.Println(len(nums), nums)
}

// 逻辑：由于是有序数组，所以后面的数一定大于等于前面的数，所以遇到重复元素时，记录当前元素的位置，找到当前元素下的第一个非重复元素，
// 并将该元素赋值给当前元素的下一个元素，以此类推，获取到一个从大到小的数组，然后返回该数组。
func delArrayRepeat2() {
	nums = []int{0, 1, 2, 2, 3, 3, 4, 4, 6, 6}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] >= nums[len(nums)-1] {
				fmt.Println(i+1, nums[:i+1])
				return
			}
			if i+1 == j && nums[i] < nums[j] { // 2和3
				break
			}
			if nums[i] < nums[j] { // 2和3, 2和4 , 3和6,
				nums[i+1] = nums[j]
				break
			} else { // 2和3相等
				continue
			}
		}
	}
}

// 只获取删除重复元素后的长度。
func delArrayRepeat3() {
	nums = []int{0, 1, 2, 2, 3, 3, 4, 4, 6, 6}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] { // 0 1, 1 2, 2 4,3 6,4 8
			j++               // 1 1,2 2,3 4,4 6,5 8
			nums[j] = nums[i] // []int{0, 1, 2, 3, 4, 6, 4, 4, 6, 6}
		}
	}
	fmt.Println(j + 1)
}

func main() {
	//delArrayRepeat()
	//delArrayRepeat2()
	delArrayRepeat3()
}
