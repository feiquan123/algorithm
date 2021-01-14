package prime

import (
	"math"
)

// IsPrime judge whether the given number is prime
type IsPrime func(n uint64) bool

// IsPrimeByRangeEnum judge whether the given number is prime
// O(n)
func IsPrimeByRangeEnum(n uint64) bool {
	if n < 2 {
		return false
	}
	for i := uint64(2); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// IsPrimeBySqrtRangeEnum judge whether the given number is prime
// O( sqrt(n) )
func IsPrimeBySqrtRangeEnum(n uint64) bool {
	if n < 2 {
		return false
	}
	for i, l := uint64(2), uint64(math.Sqrt(float64(n))); i <= l; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// NPrime 获取一个[2-n] 内的素数
func NPrime(n uint64, isPrime IsPrime) (ps []uint64) {
	ps = make([]uint64, 0)
	if n < 2 {
		return
	}

	for i := uint64(2); i <= n; i++ {
		if isPrime(i) {
			ps = append(ps, i)
		}
	}

	return
}

// NPrimeEratosthenes 埃拉托斯特尼, 最优  + 倍数筛选
// O(n * log(log(n)))
func NPrimeEratosthenes(n uint64) (ps []uint64) {
	ps = make([]uint64, 0)
	if n < 2 {
		return ps
	}

	N := make([]bool, n+1) // false: 素数, true: 不是素数
	for i, l := uint64(2), uint64(math.Sqrt(float64(n))); i <= l; i++ {
		if N[i] == false {
			for j := uint64(2); i*j <= n; j++ {
				N[i*j] = true // 倍数筛选: 同一元素会重复删除, 最优
			}
		}
	}

	for i, l := uint64(2), n+1; i < l; i++ {
		if N[i] == false {
			ps = append(ps, i)
		}
	}
	return ps
}

// NPrimeEratosthenes2 埃拉托斯特尼 + 二次筛选法
// O(n * log(log(n)))
func NPrimeEratosthenes2(n uint64) (ps []uint64) {
	ps = make([]uint64, 0)
	if n < 2 {
		return ps
	}

	N := make([]bool, n+1) // false: 素数, true: 不是素数
	for i, l := uint64(2), uint64(math.Sqrt(float64(n))); i <= l; i++ {
		if N[i] == false {
			for j := i * i; j <= n; j += i {
				N[j] = true // 二次筛选法: 假设每次i是素数，则下一个起点是 i*i ,把后面 [i*i, i*(i+1), i*(i+2), ..., n] 全部移除
			}
		}
	}

	for i, l := uint64(2), n+1; i < l; i++ {
		if N[i] == false {
			ps = append(ps, i)
		}
	}
	return ps
}

// NPrimeEratosthenesLine 埃拉托斯特尼 + 线性筛选
// O(n)
func NPrimeEratosthenesLine(n uint64) (ps []uint64) {
	ps = make([]uint64, 0)
	if n < 2 {
		return ps
	}

	N := make([]bool, n+1)  // false: 素数, true: 不是素数
	P := make([]uint64, n+1) // 保存素数
	count := uint64(0)       // 素数的个数

	for i := uint64(2); i <= n; i++ {
		if N[i] == false {
			P[count] = i
			count++
		}
		for j := uint64(0); j < count && P[j]*i <= n; j++ {
			N[P[j]*i] = true
			a := make([]uint64, n+1)
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

	for i, l := uint64(2), n+1; i < l; i++ {
		if N[i] == false {
			ps = append(ps, i)
		}
	}
	return ps
}
