package filesec

import (
	"os"
	"regexp"
	"runtime"
)

// MatchAnyRegex checks str against a list of possible regex, if any match true is returned
func MatchAnyRegex(str []byte, regex []string) bool {
	for _, reg := range regex {
		if matched, _ := regexp.Match(reg, str); matched {
			return true
		}
	}

	return false
}

func uid() int {
	if useFakeUID {
		return fakeUID
	}

	return os.Geteuid()
}

func runtimeOs() string {
	if useFakeOS {
		return fakeOS
	}

	return runtime.GOOS
}
