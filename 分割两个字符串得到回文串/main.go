package main

import "fmt"

func main() {
	//fmt.Println(checkPalindromeFormation("abdef", "fecab"))
	//fmt.Println(checkPalindromeFormation("x", "y"))
	//fmt.Println(checkPalindromeFormation("ulacfd", "jizalu"))
	//fmt.Println(checkPalindromeFormation("pvhmupgqeltozftlmfjjde", "yjgpzbezspnnpszebzmhvp"))
	fmt.Println(checkPalindromeFormation("aejbaalflrmkswrydwdkdwdyrwskmrlfqizjezd", "uvebspqckawkhbrtlqwblfwzfptanhiglaabjea"))
}

func checkPalindromeFormation(a string, b string) bool {
	return check(a, b) || check(b, a)
}

func checkPalindrome(a string) bool {
	left, right := 0, len(a)-1
	for left < right {
		if a[left] != a[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func check(a, b string) bool {
	n, m := len(a), len(b)

	left, right := 0, m-1
	for left < n && right > left {
		if a[left] == b[right] {
			left++
			right--
			if left > right || left == right {
				return true
			}
		} else {
			break
		}
	}

	//fmt.Println(left, right)
	//fmt.Println(a[0:left])
	//fmt.Println(b[right+1:])
	//fmt.Println(b[left : right+1])

	return checkPalindrome(b[left:right+1]) || checkPalindrome(a[left:right+1])
}
