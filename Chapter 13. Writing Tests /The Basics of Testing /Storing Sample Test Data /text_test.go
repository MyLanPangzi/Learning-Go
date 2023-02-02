package text

import (
	"testing"
)

func TestCountCharacters(t *testing.T) {
	count, err := CountCharacters("testdata/sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	if count != 11 {
		t.Fatal("expected 11 got ", count)
	}
	count, err = CountCharacters("testdata/notfound.txt")
	if err == nil {
		t.Fatal("expected error")
	}
}
