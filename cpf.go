package cpf

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Cpf string

func (c Cpf) String() string {
	return string(c)
}

func (c Cpf) Validate() bool {

	toInt := func(chars []string) []int {
		digits := make([]int, len(chars))
		for index, value := range chars {
			converted, _ := strconv.ParseInt(value, 10, 32)
			digits[index] = int(converted)
		}
		return digits
	}

	repeatedNumbers := func(digits []int) bool {
		hashTable := map[int]bool{}
		for _, digit := range digits {
			hashTable[digit] = true
		}
		return len(hashTable) == 1

	}

	check := func(digits []int, expected int) bool {
		sum := 0
		for index, number := range digits {
			weight := (len(digits) + 1) - index
			sum += number * weight
		}
		result := 11 - (sum % 11)
		if result == 10 {
			return 0 == expected
		}
		return result == expected
	}

	unmasked := c.Unmask()
	if len(unmasked) != 11 {
		return false
	}

	digits := toInt(strings.Split(unmasked, ""))
	if repeatedNumbers(digits) {
		return false
	}

	check1 := check(digits[:9], digits[9])
	check2 := check(digits[:10], digits[10])
	return check1 && check2

}

func (c Cpf) Mask() string {
	unmasked := c.Unmask()
	if len(unmasked) < 11 {
		return string(c)
	}
	return fmt.Sprintf("%s.%s.%s-%s", unmasked[:3], unmasked[3:6], unmasked[6:9], unmasked[9:])
}

func (c Cpf) Unmask() string {
	pattern := regexp.MustCompile(`\D`)
	unmasked := pattern.ReplaceAllString(c.String(), "")
	return unmasked
}
