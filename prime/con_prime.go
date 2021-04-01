package prime

import "context"

// 返回生成自然数序列的通道: 2,3,4,5...
func generateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			case ch <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}

// 通道过滤器: 删除能被素数整除的数
func primeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}

func ConGenPrime(ctx context.Context, n int) []int {
	r := make([]int, 0, 1024)
	ch := generateNatural(ctx)
	for i := 0; i < 10; i++ {
		prime := <-ch
		r = append(r, prime)
		ch = primeFilter(ctx, ch, prime)
	}
	return r
}
