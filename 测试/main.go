package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//slice := make([]int, 2, 5)
	//for i := 0; i < len(slice); i++ {
	//	slice[i] = i
	//}
	//
	//fmt.Printf("slice: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", slice, slice, cap(slice), len(slice), unsafe.Pointer(&slice))
	//
	//slice = append(slice, []int{1,2,3,4}...)
	//
	////changeSlice(slice)
	//fmt.Printf("slice: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", slice, slice, cap(slice), len(slice), unsafe.Pointer(&slice))

	//slice := []int{1, 2, 3, 4, 5}
	//fmt.Println(slice[:len(slice)-2])

	i := -2
	fmt.Printf("%b\n", uint(i))
	fmt.Println(uint(i))

}

func changeSlice(s []int) {
	fmt.Printf("func s: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", s, s, cap(s), len(s), unsafe.Pointer(&s))
	s = append(s, 3)
	s = append(s, 4)
	s[1] = 111
	fmt.Printf("func s: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", s, s, cap(s), len(s), unsafe.Pointer(&s))
}
