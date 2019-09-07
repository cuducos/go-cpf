package cpf

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//auxiliar type and variable to build a set
type void struct{}

var member void

//Cpf type
type Cpf string

func (c Cpf) String() string {
	return string(c)
}

func checksum(ds []int64) int64 {
	var s int64
	for i, n := range ds {
		s += n * int64(len(ds) + 1 - i)
	}
	r := 11 - (s % 11)
	if r == 10 {
		return 0
	}
	return r
}

//Validate check if Cpf is in a valid format
func (c Cpf) Validate() bool {
	u := c.Unmask()

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
  
	if len(m) == 1 {
		return false
	}

	return checksum(ds[:9]) == ds[9] && checksum(ds[:10]) == ds[10]
}

//Mask return the formated value
func (c Cpf) Mask() string {
	u := c.Unmask()
	if len(u) < 11 {
		return string(c)
	}
	return fmt.Sprintf("%s.%s.%s-%s", u[:3], u[3:6], u[6:9], u[9:])
}

//Unmask remove format and return the raw data
func (c Cpf) Unmask() string {
	return regexp.MustCompile(`\D`).ReplaceAllString(string(c), "")
}
