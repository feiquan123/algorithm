package soundex

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSoundex(t *testing.T) {
	Convey("Soundex test", t, func() {
		cases := []struct {
			s       string
			expCode string
		}{
			{
				s:       "",
				expCode: "",
			},
			{
				s:       "_A",
				expCode: "A000",
			},
			{
				s:       "GAUSS",
				expCode: "G200",
			},
			{
				s:       "lossrock",
				expCode: "L262",
			},
			{
				s:       "+_-(){}Zhong中国",
				expCode: "Z520",
			},
		}
		for _, c := range cases {
			So(Soundex(c.s), ShouldEqual, c.expCode)
		}
	})
}

func BenchmarkSoundex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soundex("+_-(){}Zhong中国")
	}
}
