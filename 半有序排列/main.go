package main

func main() {

}

func semiOrderedPermutation(nums []int) int {
	n := len(nums)
	mnIndex, mxIndex := 0, 0

	for i, num := range nums {
		if num == 1 {
			mnIndex = i
		}

		if num == n {
			mxIndex = i
		}
	}

	if mnIndex < mxIndex {
		return mnIndex + n - 1 - mxIndex
	}

	return mnIndex + n - 1 - mxIndex - 1
}
