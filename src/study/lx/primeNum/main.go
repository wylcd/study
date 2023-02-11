package main

import "fmt"

// 统计素数个数
// 遍历
func primeNum(num int) {
	pnNum := 0
	pn := map[int]int{}
	for i := 2; i <= num; i++ {
		if judgePrime(i) {
			pnNum++
			pn[pnNum] = i
		}
	}
	fmt.Println(pnNum, pn)
}

func judgePrime(i int) bool {
	for f := 2; f*f <= i; f++ { // 求素数，相当于判断两数乘积，如6，等同于2x3和3x2，只需要判断2x3可得6，
		// 便知6不是素数，而根号6*根号6也等于6，所以只需要判断在根号6以前，有没有6的约数，便可以判断6是不是素数，可以节约判断次数。
		if i%f == 0 {
			return false
		}
	}
	return true
}

// 埃式算法（埃筛法）
// 素数 非素数（合数）
func eratosthenes(n int) {
	pn := map[int]int{}
	isPrime := make([]bool, n, n)
	pnNum := 0
	for i := 2; i <= n; i++ {
		if !isPrime[i-1] {
			if !judgePrime(i) { // 当此数不是素数，标记此数，跳过。
				isPrime[i-1] = true
				break
			}
			pn[pnNum] = i
			pnNum++
			for f := i * i; f <= n; f += i { // 当一个数是素数时，它的倍数就不是素数，获取到查询上限下，所有的此素数倍数，将其标记为非素数，
				// 然后在判读时就可直接跳过，节约判断次数。
				isPrime[f-1] = true
			}
		}
	}
	fmt.Println(pnNum, pn)
}

func main() {
	primeNum(100)
	eratosthenes(100)
}
