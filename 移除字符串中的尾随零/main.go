package main

func main() {

}

func removeTrailingZeros(num string) string {

	bytes := []byte(num)

	for i := len(bytes) - 1; i >= 0; i-- {
		if bytes[i] != '0' {
			return string(bytes[0 : i+1])
		}
	}
	return ""
}
