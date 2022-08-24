package palindrome

//字符串回文串
func isPalindrome(str string) bool {

	rs := []rune(str)

	rs2 := make([]rune, 0)
	for i := len(rs) - 1; i >= 0; i-- {
		rs2 = append(rs2, rs[i])
	}

	str2 := string(rs2)

	return str == str2
}

//func longestPalindromeRune(str string) string {
//	if len(str) < 2 {
//		return str
//	}
//
//	rs := []rune(str)
//	longestStr := []rune{rs[0]}
//	for i := 1; i < len(rs); i++ {
//		l := i - 1
//		r := i + 1
//
//		longestStr = longestHelperRune(rs, l, r, longestStr)
//
//		l = i - 1
//		r = i
//
//		longestStr = longestHelperRune(rs, l, r, longestStr)
//
//	}
//
//	return string(longestStr)
//}
//
//func longestHelperRune(rs []rune, l, r int, longestStr []rune) []rune {
//	for l >= 0 && r < len(rs) {
//		if rs[l] == rs[r] {
//			if len(rs[l:r+1]) > len(longestStr) {
//				longestStr = rs[l : r+1]
//			}
//		} else {
//			break
//		}
//
//		l--
//		r++
//	}
//	return longestStr
//}

func longestPalindrome(str string) string {
	if len(str) < 2 {
		return str
	}

	longestStr := str[:1]
	for i := 1; i < len(str); i++ {
		l := i - 1
		r := i + 1

		longestStr = longestHelper(str, l, r, longestStr)

		l = i - 1
		r = i

		longestStr = longestHelper(str, l, r, longestStr)

	}

	return longestStr
}

func longestHelper(rs string, l, r int, longestStr string) string {
	for l >= 0 && r < len(rs) {
		if rs[l] == rs[r] {
			if len(rs[l:r+1]) > len(longestStr) {
				longestStr = rs[l : r+1]
			}
		} else {
			break
		}

		l--
		r++
	}
	return longestStr
}
