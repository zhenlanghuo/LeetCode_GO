package main

import (
	"fmt"
	"sync"
	"container/heap"
)

func main() {
	m := sync.Map{}
	m.Store(1, 1)
	m.Store(2, 2)
	m.Store(3, 3)
	for i := 0; i < 4; i++ {
		m.Load(1)
	}
	m.Delete(3)
	m.Store(4, 4)
	m.Delete(2)

	fmt.Println(m.Load(3))

	fmt.Println(1)
}
