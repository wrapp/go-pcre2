package pcre2_test

import (
	"regexp"
	"testing"

	"github.com/lestrrat/go-pcre2"
)

type matchStringer interface {
	MatchString(string) bool
}

func benchMatchString(b *testing.B, re matchStringer) {
	patterns := []string{`Hello World!`, `Hello Friend!`, `Hello 友達!`}
	for _, pat := range patterns {
		if !re.MatchString(pat) {
			b.Errorf("Expected to match, failed")
			return
		}
	}

	patterns = []string{`Goodbye World!`, `Hell no!`, `HelloWorld!`}
	for _, pat := range patterns {
		if re.MatchString(pat) {
			b.Errorf("Expected to NOT match, matched")
			return
		}
	}
}

func BenchmarkGoRegexpMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re, err := regexp.Compile(`^Hello (.+)!$`)
		if err != nil {
			b.Errorf("compile failed: %s", err)
			return
		}
		benchMatchString(b, re)
	}
}

func BenchmarkPCRE2RegexpMatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		re, err := pcre2.Compile(`^Hello (.+)!$`)
		if err != nil {
			b.Errorf("compile failed: %s", err)
			return
		}
		benchMatchString(b, re)
	}
}