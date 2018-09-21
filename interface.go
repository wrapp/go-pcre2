package pcre2

import "errors"


var (
	// ErrInvalidRegexp is returned when the provided Regexp is
	// not backed by a proper C pointer to pcre2_code
	ErrInvalidRegexp = errors.New("invalid regexp")
	// ErrInvalidUTF8String is returned when the input string cannot
	// be decoded into runes
	ErrInvalidUTF8String = errors.New("invalid utf8 string")
)

// ErrCompile is returned when compiling the regular expression fails.
type ErrCompile struct {
	message string
	offset  int
	pattern string
}
