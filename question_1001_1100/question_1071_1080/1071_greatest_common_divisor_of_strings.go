package question_1071_1080

// 1071. 字符串的最大公因子
// https://leetcode-cn.com/problems/greatest-common-divisor-of-strings
// Topics: 字符串

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[0:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
