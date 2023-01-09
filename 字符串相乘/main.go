package main

import "fmt"

func main() {
	fmt.Println(multiply("230", "30"))
}

func multiply(num1 string, num2 string) string {
	result := make([]byte, 0)

	num1Bytes := []byte(num1)
	reverse(num1Bytes)

	num2Bytes := []byte(num2)
	reverse(num2Bytes)

	for i, numb2 := range num2Bytes {
		start := i
		for _, numb1 := range num1Bytes {
			add(&result, start, byte((numb2-'0')*(numb1-'0')))
			start++
		}
	}

	fmt.Println(result)

	for i, _ := range result {
		result[i] += '0'
	}

	reverse(result)
	i := 0
	for ; i < len(result); i++ {
		if result[i] != '0' {
			break
		}
	}
	result = result[i:]

	if len(result) == 0 {
		return "0"
	}

	return string(result)
}

func add(result *[]byte, start int, value byte) {
	if value == 0 {
		if start >= len(*result) {
			*result = append(*result, 0)
		}
	}

	for value > 0 {
		if start >= len(*result) {
			*result = append(*result, 0)
		}
		(*result)[start] += value % 10
		value = value/10 + (*result)[start]/10
		(*result)[start] = (*result)[start] % 10
		start++
	}
}

func reverse(bytes []byte) {
	start, end := 0, len(bytes)-1
	for start < end {
		bytes[start], bytes[end] = bytes[end], bytes[start]
		start++
		end--
	}
}
