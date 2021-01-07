package prime

import (
	"math"
)

// IsPrime judge whether the given number is prime
type IsPrime func(n int64) bool

// IsPrimeByRangeEnum judge whether the given number is prime
// O(n)
func IsPrimeByRangeEnum(n int64) bool {
	if n < 2 {
		return false
	}
	for i := int64(2); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// IsPrimeBySqrtRangeEnum judge whether the given number is prime
// O( sqrt(n) )
func IsPrimeBySqrtRangeEnum(n int64) bool {
	if n < 2 {
		return false
	}
	for i, l := int64(2), int64(math.Sqrt(float64(n))); i <= l; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// NPrime 获取一个[2-n] 内的素数
func NPrime(n int64, isPrime IsPrime) (ps []int64) {
	ps = make([]int64, 0)
	if n < 2 {
		return
	}

	for i := int64(2); i <= n; i++ {
		if isPrime(i) {
			ps = append(ps, i)
		}
	}

	return
}

// NPrimeEratosthenes 埃拉托斯特尼, 最优  + 倍数筛选
// O(n * log(log(n)))
func NPrimeEratosthenes(n int64) (ps []int64) {
	ps = make([]int64, 0)
	if n < 2 {
		return ps
	}

	N := make([]bool, n+1) // false: 素数, true: 不是素数
	for i, l := int64(2), int64(math.Sqrt(float64(n))); i <= l; i++ {
		if N[i] == false {
			for j := int64(2); i*j <= n; j++ {
				N[i*j] = true // 倍数筛选: 同一元素会重复删除, 最优
			}
		}
	}

	for i, l := int64(2), n+1; i < l; i++ {
		if N[i] == false {
			ps = append(ps, i)
		}
	}
	return ps
}

// NPrimeEratosthenes2 埃拉托斯特尼 + 二次筛选法
// O(n * log(log(n)))
func NPrimeEratosthenes2(n int64) (ps []int64) {
	ps = make([]int64, 0)
	if n < 2 {
		return ps
	}

	N := make([]bool, n+1) // false: 素数, true: 不是素数
	for i, l := int64(2), int64(math.Sqrt(float64(n))); i <= l; i++ {
		if N[i] == false {
			for j := i * i; j <= n; j += i {
				N[j] = true // 二次筛选法: 假设每次i是素数，则下一个起点是 i*i ,把后面 [i*i, i*(i+1), i*(i+2), ..., n] 全部移除
			}
		}
	}

	for i, l := int64(2), n+1; i < l; i++ {
		if N[i] == false {
			ps = append(ps, i)
		}
	}
	return ps
}

// NPrimeEratosthenesLine 埃拉托斯特尼 + 线性筛选
// O(n)
func NPrimeEratosthenesLine(n int64) (ps []int64) {
	ps = make([]int64, 0)
	if n < 2 {
		return ps
	}

	N := make([]bool, n+1)  // false: 素数, true: 不是素数
	P := make([]int64, n+1) // 保存素数
	count := int64(0)       // 素数的个数

	for i := int64(2); i <= n; i++ {
		if N[i] == false {
			P[count] = i
			count++
		}
		for j := int64(0); j < count && P[j]*i <= n; j++ {
			N[P[j]*i] = true
			a := make([]int64, n+1)
			for i, v := range N {
				if v {
					a[i] = 1
				}
			}
			if i%P[j] == 0 { // 保证每个合数只会被它的最小质因数筛去，因此每个数只会被标记一次，所以时间复杂度是O(n)
				break
			}
		}
	}

	for i, l := int64(2), n+1; i < l; i++ {
		if N[i] == false {
			ps = append(ps, i)
		}
	}
	return ps
}
