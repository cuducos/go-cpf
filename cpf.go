package cpf

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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

type Cpf string

func (c Cpf) String() string {
	return string(c)
}

func (c Cpf) Validate() bool {
	u := c.Unmask()
	if len(u) != 11 {
		return false
	}

	var (
		ds = make([]int64, 11)
		m = map[int64]bool{}
	)
	for i, v := range strings.Split(u, "") {
		c, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return false
		}
		ds[i] = c
		m[c] = true
	}
	if len(m) == 1 {
		return false
	}

	return checksum(ds[:9]) == ds[9] && checksum(ds[:10]) == ds[10]
}

func (c Cpf) Mask() string {
	u := c.Unmask()
	if len(u) < 11 {
		return string(c)
	}
	return fmt.Sprintf("%s.%s.%s-%s", u[:3], u[3:6], u[6:9], u[9:])
}

func (c Cpf) Unmask() string {
	return regexp.MustCompile(`\D`).ReplaceAllString(string(c), "")
}
