package main

func main() {

}

func addMinimum(word string) int {
	ans := 0

	left := 0
	for left < len(word) {
		if word[left] == 'a' {
			if left+1 < len(word) && word[left+1] == 'b' {
				if left+2 < len(word) && word[left+2] == 'c' {
					left += 3
				} else {
					ans++
					left += 2
				}
			} else if left+1 < len(word) && word[left+1] == 'c' {
				ans += 1
				left += 2
			} else {
				ans += 2
				left++
			}
		} else if word[left] == 'b' {
			if left+1 < len(word) && word[left+1] == 'c' {
				left += 2
				ans++
			} else {
				left++
				ans += 2
			}
		} else if word[left] == 'c' {
			ans += 2
			left++
		}
	}

	return ans
}
