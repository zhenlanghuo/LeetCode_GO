package main

import "fmt"

type UidAmount struct {
	uid          string
	amount       int
	minHeadIndex int
}

func NewUidAmount() *UidAmount {
	return &UidAmount{}
}

func main() {
	//array := []UidAmount{{uid: "1"}, {uid: "2"}, {uid: "3"}}
	//result := make([]*UidAmount, 0)
	//for _, v := range array {
	//	result = append(result, &v)
	//}
	//
	//for i := 0; i < len(result); i++ {
	//	fmt.Println(*result[i])
	//}

	array := []int{1, 2, 3, 4, 5, 6}
	//for _, v := range array {
	//	fmt.Println(v)
	//	array = array[:len(array)-1]
	//}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
		array = array[:len(array)-1]
	}

	fmt.Println(array)

	a := &UidAmount{}
	b := new(UidAmount)
	c := UidAmount{}
}
