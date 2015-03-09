package uuid

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var isHexes = []bool{
	true,
	true,
	true,
	true,
	true,
	true,
	true,
	true,
	false,
	true,
	true,
	true,
	true,
	false,
	true,
	true,
	true,
	true,
	false,
	true,
	true,
	true,
	true,
	false,
	true,
	true,
	true,
	true,
	true,
	true,
	true,
	true,
	true,
	true,
	true,
	true,
}

func UUID(s string) (string, error) {
	if len(s) != 36 {
		return "", fmt.Errorf("uuid must have a length of 36 characters")
	}
	s = strings.ToLower(s)
	for i, isHex := range isHexes {
		if isHex {
			if !hex(s[i]) {
				return "", fmt.Errorf("%s is not a valid uuid", s)
			}
		} else {
			if s[i] != 45 {
				return "", fmt.Errorf("%s is not a valid uuid", s)
			}
		}
	}
	return s, nil
}

var charFuncs = []func(byte) bool{
	hex,
	hex,
	hex,
	hex,
	hex,
	hex,
	hex,
	hex,
	dash,
	hex,
	hex,
	hex,
	hex,
	dash,
	hex,
	hex,
	hex,
	hex,
	dash,
	hex,
	hex,
	hex,
	hex,
	dash,
	hex,
	hex,
	hex,
	hex,
	hex,
	hex,
	hex,
	hex,
	hex,
	hex,
	hex,
	hex,
}

func Funcs(s string) (string, error) {
	if len(s) != 36 {
		return "", fmt.Errorf("uuid must have a length of 36 characters")
	}
	s = strings.ToLower(s)
	for i, f := range charFuncs {
		if !f(s[i]) {
			return "", fmt.Errorf("%s is not a valid uuid", s)
		}
	}
	return s, nil
}

// From http://stackoverflow.com/a/20553306
const rawRegexp = `[a-f\d]{8}(-[a-f\d]{4}){3}-[a-f\d]{12}`

var compiledRegexp = regexp.MustCompile(rawRegexp)

func Regex(s string) (string, error) {
	s = strings.ToLower(s)
	if compiledRegexp.MatchString(s) {
		return s, nil
	}
	return "", fmt.Errorf("%s is not a valid uuid", s)
}

var okChars = unicode.RangeTable{
	R16: []unicode.Range16{
		unicode.Range16{Lo: 48, Hi: 57, Stride: 1},
		unicode.Range16{Lo: 97, Hi: 102, Stride: 1},
	},
}

func Runes(uuid string) (string, error) {
	if len(uuid) != 36 {
		return "", fmt.Errorf("uuid must have a length of 36 characters")
	}
	uuid = strings.ToLower(uuid)
	parts := strings.Split(uuid, "-")
	if len(parts) != 5 {
		return "", fmt.Errorf("Invalid number of uuid parts split by '-'")
	}
	if !okPartLen(parts) {
		return "", fmt.Errorf("Invalid UUID part length in %s", uuid)
	}

	for _, part := range parts {
		if badPartSyntax(part) {
			return "", fmt.Errorf("Invalid UUID syntax in %s", uuid)
		}
	}

	return uuid, nil
}

func okPartLen(parts []string) bool {
	return len(parts[0]) == 8 && len(parts[1]) == 4 && len(parts[2]) == 4 && len(parts[3]) == 4 && len(parts[4]) == 12
}

func badPartSyntax(part string) bool {
	for _, c := range part {
		if !unicode.In(rune(c), &okChars) {
			return true
		}
	}
	return false
}

func Bytes(uuid string) (string, error) {
	if len(uuid) != 36 {
		return "", fmt.Errorf("uuid must have a length of 36 characters")
	}
	uuid = strings.ToLower(uuid)
	parts := strings.Split(uuid, "-")
	if len(parts) != 5 {
		return "", fmt.Errorf("Invalid number of uuid parts split by '-'")
	}
	if !okPartLen(parts) {
		return "", fmt.Errorf("Invalid UUID part length in %s", uuid)
	}

	for _, part := range parts {
		for _, char := range part {
			if !hex(byte(char)) {
				return "", fmt.Errorf("invalid character")
			}
		}
	}

	return uuid, nil
}

func dash(char byte) bool {
	return char == 45
}

func hex(char byte) bool {
	if char < 48 || char > 102 {
		return false
	}
	return (char <= 57 || char >= 97)
}
