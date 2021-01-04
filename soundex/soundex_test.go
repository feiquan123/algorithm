package soundex

import "testing"

func TestSoundex(t *testing.T) {
	s := ""
	t.Logf("%q->%q", s, Soundex(s))

	s = "_A"
	t.Logf("%q->%q", s, Soundex(s))

	s = "GAUSS"
	t.Logf("%q->%q", s, Soundex(s))

	s = "Feng"
	t.Logf("%q->%q", s, Soundex(s))

	s = "Zho中国"
	t.Logf("%q->%q", s, Soundex(s))
}
