package algo

func isPalindrome(s string) bool {
	n := len(s)
	half := n / 2
	for i := 0; i < half; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

func isPalindromeIndex(s string, begin, end int) bool {
	if begin < 0 {
		return false
	}
	for begin < end {
		if s[begin] != s[end] {
			return false
		}
		begin++
		end--
	}
	return true
}

func longestPalindrome(s string) (result string) {
	if len(s) < 2 {
		return s
	}

	mL := 1
	n := len(s)
	result = s[:1]

	// traverse the string
	for i := 1; i < n; i++ {
		// challenge the longest case

		// version 1 but slower
		// for nowL := i + 1; nowL > mL; nowL-- {
		// 	if subS := s[i-nowL+1 : i+1]; isPalindrome(subS) {
		// 		mL = nowL
		// 		result = subS
		// 		continue
		// 	}
		// }

		// version two, each i only compute twice
		if isPalindromeIndex(s, i-mL, i) {
			// max length +1 case
			result = s[i-mL : i+1]
			mL = mL + 1
		} else if isPalindromeIndex(s, i-mL-1, i) {
			// max length +2 case
			result = s[i-mL-1 : i+1]
			mL = mL + 2
		}
	}
	return
}

// #6 Zigzag
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	slices := make([][]rune, numRows)
	for i, _ := range slices {
		slices[i] = []rune{}
	}
	adding := 1
	zindex := 0
	top := numRows - 1
	for _, v := range s {
		if zindex == 0 {
			adding = 1
		} else if zindex == top {
			adding = -1
		}
		slices[zindex] = append(slices[zindex], v)
		zindex += adding
	}
	var result []rune
	for _, v := range slices {
		result = append(result, v...)
	}
	return string(result)
}
