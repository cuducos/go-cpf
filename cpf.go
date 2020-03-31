package cpf

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//These are auxiliar type and variable to build a set
type void struct{}

var member void


func checksum(ds []int64) int64 {
	var s int64
	for i, n := range ds {
		s += n * int64(len(ds)+1-i)
	}
	r := 11 - (s % 11)
	if r == 10 {
		return 0
	}
	return r
}

//IsValid checks whether CPF number is valid or not
func  IsValid(n string) bool {
	u := Unmask(n)

	if len(u) != 11 {
		return false
	}

	ds := make([]int64, 11)
	s := make(map[int64]void)
	for i, v := range strings.Split(u, "") {
		c, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return false
		}
		ds[i] = c
		s[c] = member
	}

	//If all digits are the same, the CPF is not valid
	if len(s) == 1 {
		return false
	}

	return checksum(ds[:9]) == ds[9] && checksum(ds[:10]) == ds[10]
}

//Mask returns the CPF number formatted
func  Mask(n string) string {
	u := Unmask(n)
	if len(u) != 11 {
		return n
	}
	return fmt.Sprintf("%s.%s.%s-%s", u[:3], u[3:6], u[6:9], u[9:])
}

//Unmask removes any non-digit (numeric) from the CPF number
func  Unmask(n string) string {
	return regexp.MustCompile(`\D`).ReplaceAllString(n, "")
}
