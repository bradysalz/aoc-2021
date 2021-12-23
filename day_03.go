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

func dayThreePartTwo() int {
	var data string = "input/day_03.txt"
	maxLength, err := getLengthOfInputWord(data)
	if err != nil {
		panic(err)
	}

	words, err := readFileToNumericArray(data)
	if err != nil {
		panic(err)
	}

	oxygen := calculateOxygenRating(words, maxLength)
	c02 := calculateC02Rating(words, maxLength)
	return oxygen * c02

}

func calculateOxygenRating(words []int, maxLength int) int {
	// Loop over each bit
	for b := maxLength - 1; b > -1; b-- {
		if len(words) <= 1 {
			break
		}

		// Find the most common bit
		bitCounter := 0
		for _, word := range words {
			bitFlag := word & (1 << b)
			if bitFlag > 0 {
				bitCounter += 1
			} else {
				bitCounter -= 1
			}
		}

		// Filter the array
		var temp []int
		for _, word := range words {
			if bitCounter >= 0 {
				if hasBitSet(word, b) {
					temp = append(temp, word)
				}
			} else {
				if !hasBitSet(word, b) {
					temp = append(temp, word)
				}
			}
		}
		words = temp
	}
	return words[0]
}

// too lazy to make this an easy refactor
func calculateC02Rating(words []int, maxLength int) int {
	// Loop over each bit
	for b := maxLength - 1; b > -1; b-- {
		if len(words) <= 1 {
			break
		}

		// Find the most common bit
		bitCounter := 0
		for _, word := range words {
			bitFlag := word & (1 << b)
			if bitFlag > 0 {
				bitCounter += 1
			} else {
				bitCounter -= 1
			}
		}

		// Filter the array
		var temp []int
		for _, word := range words {
			if bitCounter >= 0 {
				if !hasBitSet(word, b) {
					temp = append(temp, word)
				}
			} else {
				if hasBitSet(word, b) {
					temp = append(temp, word)
				}
			}
		}
		words = temp
	}
	return words[0]
}
