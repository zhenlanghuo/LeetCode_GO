package main

import "fmt"

func main() {
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
}

func removeDuplicateLetters(s string) string {

	bytes := []byte(s)

	byteLastIndex := make([]int, 26)
	for i := 0; i < 26; i++ {
		byteLastIndex[i] = -1

	}

	for i := 0; i < len(bytes); i++ {
		byteLastIndex[bytes[i]-'a'] = i
	}

	stack := make([]byte, 0, len(bytes))
	byteInStack := make([]bool, 26)

	push := func(b byte) {
		stack = append(stack, b)
		byteInStack[b-'a'] = true
	}

	pop := func() {
		//fmt.Println(stack)
		topByte := stack[len(stack)-1]
		byteInStack[topByte-'a'] = false
		stack = stack[:len(stack)-1]
		//fmt.Println(stack)
	}

	for i := 0; i < len(bytes); i++ {
		//fmt.Println(i, stack)
		// 如果已经在栈中存在, 就不需要入栈了
		if byteInStack[bytes[i]-'a'] {
			continue
		}

		for len(stack) != 0 {
			//fmt.Println(i, stack)
			topByte := stack[len(stack)-1]
			if bytes[i] > topByte {
				//push(bytes[i])
				break
			}

			// 如果栈顶字符在后面还有出现, 弹出栈顶字符
			if byteLastIndex[topByte-'a'] > i {
				pop()
				continue
			}
			break
		}

		push(bytes[i])
		//fmt.Println(i, stack)
	}

	return string(stack)
}
