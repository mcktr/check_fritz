package thresholds

import "testing"

func TestIsSet(t *testing.T) {
	result := IsSet(nil)

	if result != false {
		t.Errorf("IsSet was incorrect, got: %t, want: %t", result, false)
	}

	f := 7.6

	result = IsSet(&f)

	if result != true {
		t.Errorf("IsSet was incorrect, got: %t, want: %t", result, true)
	}
}

func TestCheckLower(t *testing.T) {
	result := CheckLower(3, 2.9)

	if result != true {
		t.Errorf("CheckLower was incorrect, got: %t, want: %t.", result, true)
	}

	result = CheckLower(3, 3.1)

	if result != false {
		t.Errorf("CheckLower was incorrect, got: %t, want: %t.", result, false)
	}

	result = CheckLower(-1, 2.9)

	if result != false {
		t.Errorf("CheckLower was incorrect, got: %t, want: %t.", result, false)
	}
}

func TestCheckUpper(t *testing.T) {
	result := CheckUpper(3, 3.1)

	if result != true {
		t.Errorf("CheckUpper was incorrect, got: %t, want: %t.", result, true)
	}

	result = CheckUpper(3, 2.9)

	if result != false {
		t.Errorf("CheckLower was incorrect, got: %t, want: %t.", result, false)
	}

	result = CheckUpper(-1, 3.1)

	if result != false {
		t.Errorf("CheckUpper was incorrect, got: %t, want: %t.", result, true)
	}
}
