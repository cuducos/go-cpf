package cpf

import "testing"

func TestMask(t *testing.T) {
	if got := Cpf("11111111111").Mask(); "111.111.111-11" != got {
		t.Errorf("Cpf(\"11111111111\").Mask() = %v; want 111.111.111-11", got)
	}
}

func TestUnmask(t *testing.T) {
	if got := Cpf("111.111.111-11").Unmask(); "11111111111" != got {
		t.Errorf("Cpf(\"111.111.111-11\").Unmask() = %v; want 11111111111", got)
	}

}

func TestValidate(t *testing.T) {
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
		if got := tc.cpf.Validate(); tc.expected != got {
			t.Errorf("Cpf(%v).Validate() = %v; want %v", tc.cpf, got, tc.expected)
		}
	}

}
