package hello

import (
	"testing"
)

func TestGreet(t *testing.T){
	got := Greet()
	want := "hello world"
	if got != want{
		t.Errorf("Expected: %s but got: %s",want,got)
	}
}