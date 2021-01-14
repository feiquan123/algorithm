package prime

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNPrime(t *testing.T) {
	Convey("Prime test", t, func() {
		cases := []struct {
			n    uint64
			exPs []uint64
		}{
			{
				n:    0,
				exPs: []uint64{},
			},
			{
				n:    1,
				exPs: []uint64{},
			},
			{
				n:    2,
				exPs: []uint64{2},
			},
			{
				n:    11,
				exPs: []uint64{2, 3, 5, 7, 11},
			},
			{
				n:    16,
				exPs: []uint64{2, 3, 5, 7, 11, 13},
			},
		}

		Convey("NPrime", func() {
			Convey("IsPrimeByRangeEnum test", func() {
				for _, c := range cases {
					So(NPrime(c.n, IsPrimeByRangeEnum), ShouldResemble, c.exPs)
				}
			})

			Convey("IsPrimeBySqrtRangeEnum", func() {
				for _, c := range cases {
					So(NPrime(c.n, IsPrimeBySqrtRangeEnum), ShouldResemble, c.exPs)
				}
			})
		})

		Convey("NPrimeEratosthenes test", func() {
			for _, c := range cases {
				So(NPrimeEratosthenes(c.n), ShouldResemble, c.exPs)
			}
		})

		Convey("NPrimeEratosthenes2 test", func() {
			for _, c := range cases {
				So(NPrimeEratosthenes2(c.n), ShouldResemble, c.exPs)
			}
		})

		Convey("NPrimeEratosthenesLine test", func() {
			for _, c := range cases {
				So(NPrimeEratosthenesLine(c.n), ShouldResemble, c.exPs)
			}
		})
	})
}

var N = uint64(10000)

func BenchmarkNprimeIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NPrime(N, IsPrimeByRangeEnum)
	}
}

func BenchmarkNprimeIsPrimeBySqrtRangeEnum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NPrime(N, IsPrimeBySqrtRangeEnum)
	}
}

func BenchmarkNPrimeEratosthenes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NPrimeEratosthenes(N)
	}
}

func BenchmarkNPrimeEratosthenes2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NPrimeEratosthenes2(N)
	}
}

func BenchmarkNPrimeEratosthenesLine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NPrimeEratosthenesLine(N)
	}
}
