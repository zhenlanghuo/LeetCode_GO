package main

func main() {

}

func categorizeBox(length int, width int, height int, mass int) string {

	isHeavy := false
	isBulky := false

	if mass >= 100 {
		isHeavy = true
	}


	if length >= 1e4 || width >= 1e4 || height >= 1e4 {
		isBulky = true
	}

	if length * width * height >= 1e9 {
		isBulky = true
	}

	if isHeavy && isBulky {
		return "Both"
	}

	if !isHeavy && !isBulky {
		return "Neither"
	}

	if isHeavy {
		return "Heavy"
	}

	return "Bulky"
}
