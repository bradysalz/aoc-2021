package main

import (
	"io/ioutil"
	"log"
	"strconv"
)

type packet struct {
	version  int
	id       int
	value    int64
	children []packet
}

func daySixteenPartOne(isPartTwo bool) int64 {
	content, err := ioutil.ReadFile("input/day_16.txt")
	if err != nil {
		panic(err)
	}
	command := string(content)

	var head packet
	bits := make([]int, 0)
	parsePackets(&head, &command, bits)
	evalChildren(&head)
	if isPartTwo {
		return head.value
	}
	return scorePackets(head)
}

func scorePackets(head packet) int64 {
	val := int64(head.version)
	for c := range head.children {
		val += scorePackets(head.children[c])
	}
	return val
}

func parsePackets(head *packet, command *string, bits []int) []int {
	// parse version
	if len(bits) < 3 {
		bits = append(bits, parseChar(command)...)
	}
	head.version = bitArrToInt(bits[:3])
	bits = bits[3:]

	// parse ID
	if len(bits) < 3 {
		bits = append(bits, parseChar(command)...)
	}
	head.id = bitArrToInt(bits[:3])
	bits = bits[3:]

	// start children header
	if head.children == nil {
		head.children = make([]packet, 0)
	}

	if head.id == 4 {
		// literal packet
		keepParsing := true
		value := make([]int, 0)
		for keepParsing {
			for len(bits) < 5 {
				bits = append(bits, parseChar(command)...)
			}
			value = append(value, bits[1], bits[2], bits[3], bits[4])

			if bits[0] == 0 {
				keepParsing = false
			}
			bits = bits[5:]
		}
		head.value = bitArrToBigInt(value)
	} else {
		// operator packet
		for len(bits) < 1 {
			bits = append(bits, parseChar(command)...)
		}
		subType := bits[0]
		bits = bits[1:]
		if subType == 1 {
			// subpackets by total packets
			for len(bits) < 11 {
				bits = append(bits, parseChar(command)...)
			}
			toRead := bitArrToInt(bits[:11])
			bits = bits[11:]
			for i := 0; i < toRead; i++ {
				var newChild packet
				bits = parsePackets(&newChild, command, bits)
				head.children = append(head.children, newChild)
			}
		} else {
			// subpackets by bit-length
			for len(bits) < 15 {
				bits = append(bits, parseChar(command)...)
			}
			toParse := bitArrToInt(bits[:15])
			bits = bits[15:]

			// algorithm here is to give the children all their bits and then
			// break when they've parsed at least that many
			for len(bits) < toParse {
				bits = append(bits, parseChar(command)...)
			}
			oldLen := len(bits)
			for toParse > 0 {
				var newChild packet
				bits = parsePackets(&newChild, command, bits)
				head.children = append(head.children, newChild)

				newLen := len(bits)
				toParse -= (oldLen - newLen)
				oldLen = newLen
			}
		}
		evalChildren(head)
	}
	return bits
}

func evalChildren(head *packet) {
	for _, c := range head.children {
		if c.value < 0 {
			log.Fatalln("something overflowed?")
		}
	}
	switch head.id {
	case 0:
		sum := int64(0)
		for _, c := range head.children {
			sum += c.value
		}
		head.value = sum
	case 1:
		prod := int64(1)
		for _, c := range head.children {
			prod *= c.value
		}
		head.value = prod
	case 2:
		head.value = head.children[0].value
		for _, c := range head.children {
			if head.value > c.value {
				head.value = c.value
			}
		}
	case 3:
		head.value = head.children[0].value
		for _, c := range head.children {
			if head.value < c.value {
				head.value = c.value
			}
		}
	case 5:
		if head.children[0].value > head.children[1].value {
			head.value = int64(1)
		} else {
			head.value = int64(0)
		}
	case 6:
		if head.children[1].value > head.children[0].value {
			head.value = int64(1)
		} else {
			head.value = int64(0)
		}
	case 7:
		if head.children[0].value == head.children[1].value {
			head.value = int64(1)
		} else {
			head.value = int64(0)
		}
	}
}

func bitArrToBigInt(bits []int) int64 {
	val := int64(0)
	for b := range bits {
		val += int64(bits[b]) << (len(bits) - 1 - b)
	}
	return val
}

func bitArrToInt(bits []int) int {
	val := 0
	for b := range bits {
		val += bits[b] << (len(bits) - 1 - b)
	}
	return val
}

func parseChar(command *string) []int {
	s := *command
	char := s[0]
	s = string(s[1:])
	*command = s

	bits := make([]int, 0)
	val, err := strconv.ParseInt(string(char), 16, 0)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 4; i++ {
		bits = append([]int{int(val & 1)}, bits...)
		val = val >> 1
	}
	return bits
}
