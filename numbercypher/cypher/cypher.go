package cypher

import (
	"errors"
	"strconv"
)

func CreateNumberCypher(input string) (string, error) {
	temp := []int{0}
	for _, char := range input {
		lastnumber := temp[len(temp)-1]
		switch char {
		case 'L':
			temp = append(temp, lastnumber-1)
		case 'R':
			temp = append(temp, lastnumber+1)
		case '=':
			temp = append(temp, lastnumber)
		default:
			return "", errors.New("input error")
		}
	}

	adjust := temp[0]
	for _, v := range temp {
		if v < adjust {
			adjust = v
		}
	}
	for i, _ := range temp {
		temp[i] -= adjust
	}

	count := 0
	for i := 0; i < len(input)-1; i++ {
		if input[i] == '=' {
			count++
		} else {
			break
		}
	}
	if input[count] == 'R' {
		for i := 0; i <= count; i++ {
			temp[i] = 0
		}
	}

	count = 0
	for i := len(input) - 1; i > 0; i-- {
		if input[i] == '=' {
			count++
		} else {
			break
		}
	}
	if input[len(input)-count-1] == 'L' {
		for i := 0; i <= count; i++ {
			temp[len(temp)-count-1] = 0
		}
	}

	for i := 0; i < len(input); i++ {
		if input[i] == 'L' {
			pos := make([]int, 0)
			pos = append(pos, i)
			j := i + 1
			for j < len(input) && input[j] == '=' {
				pos = append(pos, j)
				j++
			}
			if j < len(input) && input[j] != 'L' {
				lastpos := 0
				diff := 0
				for _, v := range pos {
					if v+1 < len(input) {
						diff = temp[v+1]
						temp[v+1] = 0
					}
					lastpos = v + 2
				}
				for lastpos < len(temp) && input[lastpos-1] != 'L' {
					temp[lastpos] = temp[lastpos] - diff
					lastpos++
				}
			}

		}
	}

	resultStr := ""
	for _, digit := range temp {
		resultStr += strconv.Itoa(digit)
	}
	return resultStr, nil
}
