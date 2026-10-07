package generator

import "testing"

func TestGenerateLength(t *testing.T) {
	pw, err := Generate(Options{Length: 20, Lower: true, Upper: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(pw) != 20 {
		t.Errorf("expected length 20, got %d", len(pw))
	}
}

func TestGenerateNoCharset(t *testing.T) {
	_, err := Generate(Options{Length: 10})
	if err == nil {
		t.Error("expected error when no charset is enabled")
	}
}

func TestGenerateInvalidLength(t *testing.T) {
	_, err := Generate(Options{Length: 0, Lower: true})
	if err == nil {
		t.Error("expected error for zero length")
	}
}

func TestGenerateOnlyNumbers(t *testing.T) {
	pw, err := Generate(Options{Length: 10, Numbers: true})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	for _, c := range pw {
		if c < '0' || c > '9' {
			t.Errorf("expected only digits, got %q", pw)
			break
		}
	}
}

func TestGenerateUniqueness(t *testing.T) {
	a, _ := Generate(Options{Length: 32, Lower: true, Upper: true, Numbers: true})
	b, _ := Generate(Options{Length: 32, Lower: true, Upper: true, Numbers: true})
	if a == b {
		t.Error("two generated passwords should not be identical")
	}
}
