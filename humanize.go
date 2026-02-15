package humanize

import (
	"regexp"
	"unicode/utf8"
	"unsafe"
)

type CharRange struct {
	Start rune
	End   rune
}

type Humanizer struct {
	charsToRemove  map[rune]struct{}
	ranges         []CharRange
	regexpTrailing *regexp.Regexp
}

func NewHumanizer() *Humanizer {
	h := Humanizer{
		charsToRemove:  make(map[rune]struct{}),
		ranges:         make([]CharRange, 0),
		regexpTrailing: regexp.MustCompile(`[ \t\x0B\f]+$`),
	}
	h.addCharString("\xad\u06dd\u070f\u200b\u200c\u200e\u200f\ufeff\U000110bd\U000e0001\u00AD\u180E\u2060\uFEFF")
	h.addRange('\x00', '\x08')
	h.addRange('\x0e', '\x1b')
	h.addRange('\x7f', '\x9f')
	h.addRange('\u0600', '\u0604')
	h.addRange('\u202a', '\u202e')
	h.addRange('\u2060', '\u2064')
	h.addRange('\u206a', '\u206f')
	h.addRange('\ufff9', '\ufffb')
	h.addRange('\U0001d173', '\U0001d17a')
	h.addRange('\U000e0020', '\U000e007f')
	h.addRange('\u200B', '\u200F')
	h.addRange('\u202A', '\u202E')
	h.addRange('\u2066', '\u2069')

	return &h
}

func (h *Humanizer) HumanizeString(text string) string {
	return string(h.Humanize(unsafe.Slice(unsafe.StringData(text), len(text))))
	// return string(h.Humanize([]byte(text)))
}

func (h *Humanizer) Humanize(data []byte) []byte {
	result := make([]byte, 0, len(data))

	for len(data) > 0 {
		r, size := utf8.DecodeRune(data)
		tmp := data[:size]
		data = data[size:]
		if r == utf8.RuneError {
			// Replace invalid UTF-8 with a space
			result = append(result, ' ')
			continue
		}
		if h.shouldRemove(r) {
			continue
		}

		switch r {
		case '\u00A0':
			// Replace non-breaking space with a space
			result = append(result, ' ')
			continue
		case '—', '–':
			// Replace hyphens with a dash
			result = append(result, '-')
			continue
		case '“', '”', '«', '»', '„':
			// Replace double quotes with a double quote
			result = append(result, '"')
			continue
		case '‘', '’', 'ʼ':
			// Replace single quotes with a single quote
			result = append(result, '\'')
			continue
		case '…':
			// Replace ellipsis with a ellipsis
			result = append(result, `...`...)
			continue
		default:
			// Keep the character
			result = append(result, tmp...)
		}
	}

	return h.regexpTrailing.ReplaceAll(result, []byte{})
}

// MARK: Internal

func (h *Humanizer) addCharString(s string) *Humanizer {
	for _, char := range s {
		h.charsToRemove[char] = struct{}{}
	}

	return h
}

func (h *Humanizer) addRange(start, end rune) *Humanizer {
	if start > end {
		start, end = end, start
	}
	h.ranges = append(h.ranges, CharRange{Start: start, End: end})

	return h
}

func (h *Humanizer) shouldRemove(r rune) bool {
	if _, ok := h.charsToRemove[r]; ok {
		return true
	}

	for _, rang := range h.ranges {
		if r >= rang.Start && r <= rang.End {
			return true
		}
	}

	return false
}
