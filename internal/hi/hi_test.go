package hi

import "testing"

func TestHi(t *testing.T) {
	want := "Hi! Hello, world."
	if got := Hi(); got != want {
		t.Errorf("Hi() = %q, want %q", got, want)
	}
}
