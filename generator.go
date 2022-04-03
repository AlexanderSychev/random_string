package random_string

import (
	"math/rand"
	"strings"
)

// Generator is common object which provides generation of random string
type Generator struct {
	// All available runes for generated string
	charset []rune
	// Length of generated string
	length int
}

// Generate creates string with random characters
func (g Generator) Generate() string {
	var output strings.Builder
	for i := 0; i < g.length; i++ {
		random := rand.Intn(len(g.charset))
		randomChar := g.charset[random]
		output.WriteRune(randomChar)
	}

	return output.String()
}

// NewGenerator creates Generator instance with received length and characters set runes slice
func NewGenerator(length int, charset []rune) Generator {
	return Generator{length: length, charset: charset}
}

// NewGeneratorFromString creates Generator instance with received length and characters set string
func NewGeneratorFromString(length int, charset string) Generator {
	return NewGenerator(length, []rune(charset))
}
