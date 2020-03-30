package cpf

import (
	"fmt"
	"testing"
)

func TestMask(t *testing.T) {
	for _, tc := range []struct {
		cpf      string
		expected string
	}{
		{"11111111111", "111.111.111-11"},
		{"123456", "123456"},
		{"11223344556677889900", "11223344556677889900"},
	} {
		if got := Mask(tc.cpf); tc.expected != got {
			t.Errorf("Mask(\"%s\") = %v; expected %s", tc.cpf, got, tc.expected)
		}
	}
}

func TestUnmask(t *testing.T) {
	if got := Unmask("111.111.111-11"); "11111111111" != got {
		t.Errorf("Unmask(\"111.111.111-11\") = %v; want 11111111111", got)
	}
}

func TestIsValid(t *testing.T) {
	for _, tc := range []struct {
		cpf      string
		expected bool
	}{
		{"23858488135", true},
		{"238.584.881-35", true},
		{"123", false},
		{"111.111.111-11", false},
		{"123.456.769/01", false},
		{"ABC.DEF.GHI-JK", false},
	} {
		if got := IsValid(tc.cpf); tc.expected != got {
			t.Errorf("IsValid(%v) = %v; expected %v", tc.cpf, got, tc.expected)
		}
	}
}

func ExampleIsValid_validUnmasked() {
	fmt.Println(IsValid("23858488135"))
	// Output: true
}

func ExampleIsValid_validMasked() {
	fmt.Println(IsValid("238.584.881-35"))
	// Output: true
}

func ExampleIsValid_invalid() {
	fmt.Println(IsValid("111.111.111-11"))
	// Output: false
}

func ExampleMask_valid() {
	fmt.Println(Mask("11111111111"))
	// Output: 111.111.111-11
}

func ExampleMask_invalid() {
	fmt.Println(Mask("42"))
	// Output: 42
}

func ExampleUnmask() {
	fmt.Println(Unmask("111.111.111-11"))
	// Output: 11111111111
}
