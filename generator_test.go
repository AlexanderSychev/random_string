package random_string

import (
	"reflect"
	"strings"
	"testing"
)

func TestGenerator_String(t *testing.T) {
	t.Log("Test \"String\" method of \"Generator\" type")

	length := 10
	charset := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}
	expected := "Generator(length = 10 characters, charset = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'])"

	gen := Generator{
		length: length,
		charset: charset,
	}

	sGen := gen.String()
	if sGen != expected {
		t.Fatalf("Expected \"%s\", got \"%s\"", expected, sGen)
	}
}

func TestGenerator_Generate(t *testing.T) {
	t.Log("Tests for \"Generate\" method of \"Generator\" type")

	length := 10
	charset := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}
	nonCharset := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}

	gen := Generator{
		length: length,
		charset: charset,
	}

	t.Logf("  Test #1: Generated string must have length %d and contain runes %s", length, string(charset))
	str01 := gen.Generate()
	if len(str01) != length {
		t.Fatalf("  Expected length is %d, got %d", length, len(str01))
	}
	for _, r := range nonCharset {
		if strings.ContainsRune(str01, r) {
			t.Fatalf("  Unexpected rune '%q' at generated string \"%s\"", r, str01)
		}
	}
	t.Log("  Succeed")

	t.Logf("  Test #2: Generated string must have length %d and contain runes %s", length, string(charset))
	str02 := gen.Generate()
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

func TestNewGenerator(t *testing.T) {
	t.Log("Test \"NewGenerator\" function")

	length := 10
	charset := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}

	generator := NewGenerator(length, charset)
	expected := Generator{
		length: length,
		charset: charset,
	}

	if !reflect.DeepEqual(generator, expected) {
		t.Fatalf("Expected %s got %s", generator, expected)
	}
}

func TestNewGeneratorFromString(t *testing.T) {
	t.Log("Test \"NewGeneratorFromString\" function")

	length := 10
	charset := "ABCDEFGH"

	generator := NewGeneratorFromString(length, charset)
	expected := Generator{
		length: length,
		charset: []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'},
	}

	if !reflect.DeepEqual(generator, expected) {
		t.Fatalf("Expected %s got %s", generator, expected)
	}
}
