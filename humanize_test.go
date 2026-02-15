package humanize

import (
	"testing"
)

func TestWhitespaces(t *testing.T) {
	hm := NewHumanizer()
	result := hm.HumanizeString("Hello\u200b\xa0World!  ")

	if result != "Hello World!" {
		t.Errorf("Expected: 'Hello World!', got: '%s'", result)
	}
}

func TestDashes(t *testing.T) {
	hm := NewHumanizer()
	result := hm.HumanizeString("I — super — man – 💪")

	if result != "I - super - man - 💪" {
		t.Errorf("Expected: 'I - super - man - 💪', got: '%s'", result)
	}
}

func TestQuotes(t *testing.T) {
	hm := NewHumanizer()
	result := hm.HumanizeString("Angular “quote” «marks» looks„ like Christmas «« tree")

	if result != `Angular "quote" "marks" looks" like Christmas "" tree` {
		t.Errorf(`Expected: 'Angular "quote" "marks" looks" like Christmas "" tree', got: '%s'`, result)
	}
}

func TestEllipsis(t *testing.T) {
	hm := NewHumanizer()
	result := hm.HumanizeString("Go on…")

	if result != "Go on..." {
		t.Errorf("Expected: 'Go on...', got: '%s'", result)
	}
}
