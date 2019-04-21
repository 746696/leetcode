package question_31_40

import (
	"strconv"
)

func countAndSay(n int) string {
	tmp := ""
	result := "1"
	for i := 2; i <= n; i++ {
		c := 1
		t := result[0]
		for j := 1; j < len(result); j++ {
			if t == result[j] {
				c++
				continue
			}
			tmp += strconv.Itoa(c) + string(t)
			c = 1
			t = result[j]
		}
		result = tmp + strconv.Itoa(c) + string(t)
		tmp = ""
	}
	return result
}