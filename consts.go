package pcre2

/*
#define PCRE2_CODE_UNIT_WIDTH 32
#cgo pkg-config: libpcre2-32
#include <pcre2.h>
*/
import "C"

const (
	PCRE2JITComplete = C.PCRE2_JIT_COMPLETE
	PCRE2JITPartialHard = C.PCRE2_JIT_PARTIAL_HARD
	PCRE2JITPartialSoft= C.PCRE2_JIT_PARTIAL_SOFT
	)
