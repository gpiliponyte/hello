package howdy

import "testing"

func TestHello(t *testing.T) {
    want := "Howdy"
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}