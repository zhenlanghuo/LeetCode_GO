package main

import "fmt"

func main()  {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	martx := make([][]byte, 0)

	last := 0
	count := 0
	for {
		martx = append(martx, make([]byte, numRows))
		if count%(numRows-1) == 0 {
			for i := last; i < len(s) && i < last + numRows ; i++ {
				martx[count][i-last] = s[i]
			}
			last += numRows
		} else {
			martx[count][numRows-1-count%(numRows-1)] = s[last]
			last++
		}
		count++
		if last >= len(s) {
			break
		}
	}

	bytes := make([]byte, 0)
	for i:=0;i<numRows;i++ {
		for j:=0;j<len(martx);j++ {
			if martx[j][i] != 0 {
				bytes = append(bytes, martx[j][i])
			}
		}
	}

	return string(bytes)
}
