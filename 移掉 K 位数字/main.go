package 移掉_K_位数字

func main() {

}

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0)
	n := len(num)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && len(stack)+n-i > n-k && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) < n-k {
			stack = append(stack, num[i])
		}
	}

	if len(stack) == 0 {
		return "0"
	}

	i := 0
	for i < len(stack) && stack[i] == '0' {
		i++
	}
	stack = stack[i:]

	return string(stack)
}
