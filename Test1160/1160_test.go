package Test1160

import (
	"testing"
)

func Test(t *testing.T) {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	cnt := countCharacters(words, chars)
	t.Log(cnt)
}

func countCharacters(words []string, chars string) int {
	maps := map[int32]int{}
	for _, v := range chars {
		maps[v]++
	}
	var cnt int
	for _, v := range words {
		tempV := ""
		temps := map[int32]int{}
		for _, w := range v {
			if maps[w] < 1 {
				tempV = ""
				break
			} else {
				tempV += string(w)
				temps[w]++
				if temps[w] > maps[w] {
					tempV = ""
					break
				}
			}
		}
		cnt += len(tempV)

	}
	return cnt
}
