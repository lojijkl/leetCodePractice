package Test824

import (
	"strings"
	"testing"
)

func Test(t *testing.T) {
	s := "I speak Goat Latin"
	//s:="HZ sg L"
	t.Log(toGoatLatin(s))
}

func toGoatLatin(S string) string {
	ss := strings.Split(S, " ")
	for i := 0; i < len(ss); i++ {
		if !maps[string(ss[i][0])] {
			ss[i] = ss[i][1:] + string(ss[i][0]) + "ma"
		} else if strings.ContainsAny(ss[i], "aeiouAEIOU") {
			ss[i] += "ma"
		}
		for j := i + 1; j > 0; j-- {
			ss[i] += "a"
		}
	}
	return strings.Join(ss, " ")
}

var maps = map[string]bool{
	"a": true,
	"e": true,
	"i": true,
	"o": true,
	"u": true,
	"A": true,
	"E": true,
	"I": true,
	"O": true,
	"U": true,
}
