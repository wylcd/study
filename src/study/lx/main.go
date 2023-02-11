package main

import (
	"fmt"
)

// defer先入后出
func b() (i int) {
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return i //或者直接写成return
}

func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("defer2:", i)
	}()
	defer func() {
		i++
		fmt.Println("defer1:", i)
	}()
	return &i
}

//func main() {
//	fmt.Println("return:", b())
//
//	fmt.Println("returns:", *(c()))
//}

// 方法和有参函数，会在运行到对应defer时，代入参数，先获取到结果，主函数结束时输出。
// 无参函数，只会在主函数结束时，再执行函数。
func deferTest(a int) { //无返回值函数
	defer fmt.Println("1、a =", a)                    //方法 1
	defer func(v int) { fmt.Println("2、a =", v) }(a) //有参函数 1
	defer func() { fmt.Println("3、a =", a) }()       //无参函数 2
	a++
}

// 有参函数
func deferTest2() (r int) {
	// 此时有参函数传入的r不是返回值的r，是使用值传递，复制了一个r = r´，将r´代入了有参函数，在函数中修改值传递的参数，与原本参数无关。
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func deferTest3() {
	if true {
		defer fmt.Printf("1")
	} else {
		defer fmt.Printf("2")
	}
	fmt.Printf("3")
}

func test4() {
	i := 1
	j := 2
	i, j = j, i
	fmt.Println(i, j)

	a := []string{"a", "b", "c"}
	for b := range a {
		fmt.Println(b)
	}
}

// 回文判定
func huiWen() {
	str := "天连水碧水连天"
	runeStr := []rune(str)
	for i := 0; i < (len(runeStr)-1)/2; i++ {
		if runeStr[i] != runeStr[len(runeStr)-i-1] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
	return
}

func DeferFunc4() (t int) {
	defer func(i int) {
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	t = 1
	return 2
}

func nilMap() {
	m := map[int]int{}
	for key, value := range m {
		fmt.Println(key, value)
	}
	fmt.Println(m)
}

func main() {
	//deferTest(1)

	//r := deferTest2()
	//fmt.Println(r)
	//DeferFunc4()

	//test2()
	//a := 12^2	// ^:异或 1100 ^ 0010 = 1110 = 14
	//fmt.Println(a)
	//test4()

	// 回文判定。
	//huiWen()

	//a := make([]int, 5, 10)
	//b := new([5]string)
	//
	//fmt.Printf("%T\n", a)
	//fmt.Printf("%T", b)

	//c := map[int]int{
	//	1:1,
	//	2:2,
	//	3:3,
	//	4:4,
	//	5:5,
	//}
	//
	//for key, value := range c {
	//	fmt.Printf("key:%d, value:%d\n", key, value)
	//}

	//fmt.Println(2^2)
	//nilMap()

	// go中，函数传递都是值传递。
	//a := make([]int64, 0)
	//a = append(a, 1)
	//testSlice(a)
	//fmt.Println(a)

	//b := make(map[int64]int64, 2)
	//b[1] = 1
	//testMap(b)
	//fmt.Println(b)

	//a := []int64{1, 2, 3}
	//b := a[0:2]
	//b = append(b, 50)
	//fmt.Println(a)
	//b = append(b, 60)
	//b[0] = 20
	//fmt.Println(a)

	fmt.Println("我被酒色所伤，竟如此憔悴，从今日起，戒酒！")
	fmt.Println("死生契阔，与子成说；执子之手，与子偕老。")
	fmt.Println("隐约雷鸣，阴霾天空，但盼风雨来，能留你在此。")
	fmt.Println("隐约雷鸣，阴霾天空，纵使天无雨，我亦留此地。")
	fmt.Println("github push test")
	fmt.Println("github pull test")
}

func testSlice(a []int64) {
	a = append(a, 2)
	fmt.Println("testSlice:", a)
}

func testMap(b map[int64]int64) {
	b = map[int64]int64{2: 2} // 在函数中重新覆盖了map的值，但是主函数中输出的map值没有变更，表明当前还是值传递。
	fmt.Println("testMap:", b)
}
