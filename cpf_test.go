package cpf

import "testing"

func TestMask(t *testing.T) {
	for _, tc := range []struct {
		cpf      string
		expected string
	}{
		{"11111111111", "111.111.111-11"},
		{"123456", "123456"},
		{"11223344556677889900", "11223344556677889900"},
	} {
		if got := Cpf(tc.cpf).Mask(); tc.expected != got {
			t.Errorf("Cpf(\"%s\").Mask() = %v; want %s", tc.cpf, got, tc.expected)
		}
	}
}

func TestUnmask(t *testing.T) {
	if got := Cpf("111.111.111-11").Unmask(); "11111111111" != got {
		t.Errorf("Cpf(\"111.111.111-11\").Unmask() = %v; want 11111111111", got)
	}
}

func TestIsValid(t *testing.T) {
	for _, tc := range []struct {
		cpf      Cpf
		expected bool
	}{
		{Cpf("23858488135"), true},
		{Cpf("238.584.881-35"), true},
		{Cpf("123"), false},
		{Cpf("111.111.111-11"), false},
		{Cpf("123.456.769/01"), false},
		{Cpf("ABC.DEF.GHI-JK"), false},
	} {
		if got := tc.cpf.IsValid(); tc.expected != got {
			t.Errorf("Cpf(%v).IsValid() = %v; want %v", tc.cpf, got, tc.expected)
		}
	}
}
