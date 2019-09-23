package Test1002

import (
	"testing"
)

func Test(t *testing.T) {
	t.Log(commonChars([]string{"bella", "label", "roller"}))
	t.Log(commonChars([]string{"cool", "lock", "cook"}))
	t.Log(commonChars([]string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"}))
}
func Benchmark_Test(b *testing.B) {
	for i := 0; i < b.N; i++ {
		commonChars([]string{"bella", "label", "roller"})
		commonChars([]string{"cool", "lock", "cook"})
		commonChars([]string{"acabcddd", "bcbdbcbd", "baddbadb", "cbdddcac", "aacbcccd", "ccccddda", "cababaab", "addcaccd"})
	}
}

func commonChars(A []string) []string {
	maps := map[int32]int{}  //maps声明，主要放要返回的字符，及个数
	for _, v := range A[0] { //第一个值放到默认maps里面
		maps[v]++
	}
	temp := map[int32]int{}
	for i := 1; i < len(A); i++ {
		temp = map[int32]int{} //每个字符串里面的字符都放到temp里面，用于与maps的比较
		for _, data := range A[i] {
			if maps[data] > 0 { //如果maps存在的话计数，不存在的话不是有效数据
				temp[data]++
			}
		}
		for k, _ := range maps {
			if temp[k] == 0 { //如果temp里面不存在，表示maps里面有无效数据
				delete(maps, k)
			} else if maps[k] > temp[k] { //如果temp里面的数据更少，更新maps数据
				maps[k] = temp[k]
			}
		}
	}
	A = nil
	for k, v := range maps { //封装返回值
		for v > 0 {
			A = append(A, string(k))
			v--
		}
	}
	return A
}
