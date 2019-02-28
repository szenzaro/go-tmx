package tmx

import (
	"fmt"
	"testing"
)

func TestTmxRead(t *testing.T) {
	path := "test-data/example-document.tmx"
	tmx, err := Read(path)
	if err != nil {
		t.Errorf("Expected to be able to read file at %v, got %v", path, err)
	}
	fmt.Println(tmx.Body.Tu[0].Tuid)
}

func TestTranslation(t *testing.T) {
	path := "test-data/example-document.tmx"
	tmx, err := Read(path)
	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		descr    string
		n        int // input
		expected int // expected result
	}{
		{"Number of tu", len(tmx.Body.Tu), 2},
		{"Notes in first tu", len(tmx.Body.Tu[0].Note), 1},
		{"Props in first tu", len(tmx.Body.Tu[0].Prop), 2},
		{"Tuvs in first tu", len(tmx.Body.Tu[0].Tuv), 2},
	}

	for _, tt := range tests {
		if tt.n != tt.expected {
			t.Errorf("%s: expected %d, actual %d", tt.descr, tt.expected, tt.n)
		}
	}
}
