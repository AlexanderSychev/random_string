package random_string

import (
	"strings"
	"testing"
)

func funcsTestCase(t *testing.T, fn func(int) string, charset string, nonCharset []rune) {
	length := 10

	t.Logf("  Test #1: Generated string must have length %d and contain runes %s", length, charset)
	str01 := fn(length)
	if len(str01) != length {
		t.Fatalf("  Expected length is %d, got %d", length, len(str01))
	}
	for _, r := range nonCharset {
		if strings.ContainsRune(str01, r) {
			t.Fatalf("  Unexpected rune '%q' at generated string \"%s\"", r, str01)
		}
	}
	t.Log("  Succeed")

	t.Logf("  Test #1: Generated string must have length %d and contain runes %s", length, charset)
	str02 := fn(length)
	if len(str02) != length {
		t.Fatalf("  Expected length is %d, got %d", length, len(str02))
	}
	for _, r := range nonCharset {
		if strings.ContainsRune(str02, r) {
			t.Fatalf("  Unexpected rune '%q' at generated string \"%s\"", r, str02)
		}
	}
	t.Log("  Succeed")

	t.Log("  Test #3: Generated strings cannot be equal")
	if str01 == str02 {
		t.Fatalf("  Expected non-equal strings, got \"%s\" and \"%s\"", str01, str02)
	}
	t.Log("  Succeed")
}

func TestGenerate(t *testing.T) {
	t.Log("Tests for \"Generate\" function")
	funcsTestCase(t, Generate, DefaultCharset, []rune{'+', '-', '=', '*'})
}

func TestGenerateDigits(t *testing.T) {
	t.Log("Tests for \"GenerateDigits\" function")
	funcsTestCase(t, GenerateDigits, DigitsCharset, []rune(LatinUpperCaseCharset + LatinLowerCaseCharset))
}

func TestGenerateLowerCase(t *testing.T) {
	t.Log("Tests for \"GenerateLowerCase\" function")
	funcsTestCase(t, GenerateLowerCase, LatinLowerCaseCharset, []rune(LatinUpperCaseCharset + DigitsCharset))
}

func TestGenerateUpperCase(t *testing.T) {
	t.Log("Tests for \"GenerateUpperCase\" function")
	funcsTestCase(t, GenerateUpperCase, LatinUpperCaseCharset, []rune(DigitsCharset + LatinLowerCaseCharset))
}
