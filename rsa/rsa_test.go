package rsa

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGCD(t *testing.T) {
	Convey("GCD test", t, func() {
		cases := []struct {
			m, n, expR uint64
		}{
			{
				m:    1,
				n:    3,
				expR: 1,
			},
			{
				m:    0,
				n:    3,
				expR: 3,
			},
			{
				m:    20,
				n:    15,
				expR: 5,
			},
			{
				m:    100,
				n:    0x50,
				expR: 20,
			},
		}

		for _, c := range cases {
			So(GCD(c.m, c.n), ShouldEqual, c.expR)
		}
	})
}

func TestRSA(t *testing.T) {
	// rsa2.EncryptPKCS1v15()
	// rsa2.DecryptPKCS1v15()

	Convey("given a rsa", t, func() {
		rsa := New(3000)
		t.Logf("rsa: %+v", rsa)

		Convey("PublicKey test", func() {
			k := rsa.PublicKey()
			So(k, ShouldNotEqual, "")
			t.Log("PublicKey:" + k)
		})

		Convey("Encode&Dcode test", func() {
			cases := []string{
				string(byte(7)),
				"assdfs",
				"asdffygvbmnasldbfhbajksjdfljbasdbf asjdf",
			}

			for _, c := range cases {
				es := string(rsa.Encode([]byte(c), nil))
				ds, err := rsa.Decode([]byte(es), nil)
				t.Logf("%s%v-- encode -->%q%v-- decode -->%q%v\n", c, []byte(c), es, []byte(es), ds, []byte(ds))
				So(err, ShouldBeNil)
				So(string(ds), ShouldEqual, c)
			}
		})
	})
}

func BenchmarkGCD(b *testing.B) {
	for i := -100; i < b.N; i++ {
		j := uint64(i)
		GCD(j, j+10)
	}
}

func BenchmarkTestRSAEncode(b *testing.B) {
	rsa := New(0)
	bb := []byte("asdffygvbmnasldbfhbajksjdfljbasdbf asjdf")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rsa.Encode(bb, &PublicKey{19787, 17117})
	}
}

func BenchmarkTestRSADecode(b *testing.B) {
	rsa := New(0)
	bb := []byte("L9\x02i84\xa6\x02i8\x1e\x02\x02i8\x1e\xb2\x02i8\x1e\xb2\x02i8*d\x02i8\x10\xc6\x02i8\x05\x0e\x02i8\aN\x02i8)z\x02i82%\x02i8L9\x02i84\xa6\x02i8=U\x02i8\x1e\x02\x02i8\aN\x02i8\x1e\xb2\x02i8+D\x02i8\aN\x02i8L9\x02i8:\x12\x02i8\"\xdb\x02i84\xa6\x02i8:\x12\x02i8\x1e\x02\x02i8\x1e\xb2\x02i8=U\x02i8:\x12\x02i8\aN\x02i8L9\x02i84\xa6\x02i8\x1e\x02\x02i8\aN\x02i8\x1e\xb2\x02i8)\xe0\x02i8L9\x02i84\xa6\x02i8:\x12\x02i8\x1e\x02\x02i8\x1e\xb2")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rsa.Decode(bb, &PrivateKey{19787, 11813})
	}
}
