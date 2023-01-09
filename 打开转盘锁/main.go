package main

import "fmt"

func main() {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0000"))
	fmt.Println(openLock([]string{"8888"}, "0009"))

	//fmt.Println(bytesToInt([]byte("0100")))
}

func openLock(deadends []string, target string) int {

	if target == "0000" {
		return 0
	}

	usedMap := make(map[int]int)
	for _, deadend := range deadends {
		usedMap[bytesToInt([]byte(deadend))] = -1
	}
	usedMap[0] = 0

	targetInt := bytesToInt([]byte(target))
	//fmt.Println(targetInt)

	queue := make([][]byte, 0)
	queue = append(queue, []byte("0000"))
	ops := []int{-1, +1}
	for i := 0; i < len(queue); i++ {
		times := usedMap[bytesToInt(queue[i])]
		if times == -1 {
			continue
		}
		for j, b := range queue[i] {
			//fmt.Println(11111, string(queue[i]), queue[i])
			for _, op := range ops {
				bytes := make([]byte, 4)
				copy(bytes, queue[i])
				bytes[j] = byte((int(b-'0')+op+10)%10) + '0'
				bytesInt := bytesToInt(bytes)
				_, ok := usedMap[bytesInt]
				if ok {
					continue
				}
				//fmt.Println(string(bytes))
				if bytesInt == targetInt {
					//fmt.Println(queue)
					return times + 1
				}
				usedMap[bytesInt] = times + 1
				queue = append(queue, bytes)
			}
		}
	}

	return -1
}

func bytesToInt(bytes []byte) int {
	result := 0
	for _, b := range bytes {
		result = result*10 + int(b-'0')
	}
	return result
}
