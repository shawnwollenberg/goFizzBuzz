package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	want := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}
	if got := fizzbuzz(15); !reflect.DeepEqual(want, got) {
		t.Errorf("fizzbuzz() = %q, want %q", got, want)
	}

}
