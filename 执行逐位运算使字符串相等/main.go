package main

func main() {

}

func makeStringsEqual(s string, target string) bool {
	if s == target {
		return true
	}

	return hasOne(s) && hasOne(target)
}

func hasOne(s string) bool {
	for _, c := range s {
		if c == '1' {
			return true
		}
	}
	return false
}
