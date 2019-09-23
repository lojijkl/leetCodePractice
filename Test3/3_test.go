package Test3

import (
	"testing"
)

func Test(t *testing.T) {
	t.Log(lengthOfLongestSubstring(" "))
	t.Log(lengthOfLongestSubstring("abcabcbb"))
	t.Log(lengthOfLongestSubstring("bbbbb"))
	t.Log(lengthOfLongestSubstring("pwwkew"))
	t.Log(lengthOfLongestSubstring("cdd"))
	t.Log(lengthOfLongestSubstring("aab"))
	t.Log(lengthOfLongestSubstring("dvdf"))
	t.Log(lengthOfLongestSubstring("abba"))
}

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	maps := map[byte]int{}
	max := 0
	start := 0
	i := 0
	for ; i < len(s); i++ {
		if maps[s[i]] > 0 {
			if i-start > max {
				max = i - start
			}
			if start < maps[s[i]] {
				start = maps[s[i]]
			}
		}
		maps[s[i]] = i + 1
	}

	if len(s)-start > max {
		max = len(s) - start
	}
	return max
}
