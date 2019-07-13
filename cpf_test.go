package cpf

import "testing"

func TestMask(t *testing.T) {
	cpf := Cpf("11111111111")
	expect := "111.111.111-11"
	got := cpf.Mask()

	if expect != got {
		t.Errorf("Expected %s to be %s but got %s", cpf, expect, got)
	}
}

func TestUnmask(t *testing.T) {
	cpf := Cpf("111.111.111-11")
	expect := "11111111111"
	got := cpf.Unmask()

	if expect != got {
		t.Errorf("Expected %s to be %s but got %s", cpf, expect, got)
	}

}

func TestValidate(t *testing.T) {

	testTable := []struct {
		cpf      Cpf
		expected bool
	}{
		{Cpf("23858488135"), true},
		{Cpf("238.584.881-35"), true},
		{Cpf("123"), false},
		{Cpf("111.111.111-11"), false},
		{Cpf("123.456.769/01"), false},
	}

	for _, testCase := range testTable {
		assertValidate(t, testCase.cpf, testCase.expected)
	}

}

func assertValidate(t *testing.T, cpf Cpf, expected bool) {
	t.Helper()
	got := cpf.Validate()

	if got != expected {
		t.Errorf("expected %q to be %t but got %t", cpf, expected, got)
	}
}
