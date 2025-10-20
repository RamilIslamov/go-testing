package panic_and_error

import "testing"

func TestValidateName_WithError(t *testing.T) {
	err := ValidateName("")

	if err == nil {
		t.Fatal("ожидали ошибку, но получили nil")
	} else if err.Error() != "name is empty" {
		t.Fatalf("unexpected error: %s", err.Error())
	}
}

func TestValidateName_NonEmpty(t *testing.T) {
	if err := ValidateName("Hexlet"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestMustAt_WithPanic_OutOfBoundaries(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("ожидали панику, но её не было")
		} else if r != "index out of range" {
			t.Errorf("got %q, want %q", r, "index out of range")
		}
	}()
	MustAt(2, []string{"string1", "string2"})
}

func TestMustAt_WithPanic_NegativeIndex(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("ожидали панику, но её не было")
		} else if r != "index out of range" {
			t.Errorf("got %q, want %q", r, "index out of range")
		}
	}()
	MustAt(-1, []string{"string1", "string2"})
}

func TestMustAt_WithoutPanic(t *testing.T) {
	res := MustAt(0, []string{"string1", "string2"})
	if res != "string1" {
		t.Errorf("got %q, want string1", res)
	}
}
