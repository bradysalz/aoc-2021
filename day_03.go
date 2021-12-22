package main

func dayThreePartOne() int {
	var data string = "input/day_03.txt"
	vals, err := readFileToStringArray(data)
	check(err)

	var numBits int
	bitCount := make([]int, 0)
	for i, row := range vals {
		if i == 0 {
			numBits = len(row)
			bitCount = make([]int, numBits)
		}

		for c, char := range row {
			s := string(char)
			if s == "1" {
				bitCount[c] += 1
			}
			if s == "0" {
				bitCount[c] -= 1
			}
		}
	}

	var gamma int
	var epsilon int
	for b, val := range bitCount {
		var bValue int
		if val > 0 {
			bValue = 1
		} else {
			bValue = 0
		}

		gamma += bValue << (numBits - b - 1)
	}

	epsilon = calculateInverseValue(gamma, numBits)
	return gamma * epsilon
}
