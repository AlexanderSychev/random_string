package random_string

// Generate generates a string of random digits and latin characters
func Generate(length int) string {
	generator := NewGeneratorFromString(length, DefaultCharset)
	return generator.Generate()
}

// GenerateDigits generates a string of random digits
func GenerateDigits(length int) string {
	generator := NewGeneratorFromString(length, DigitsCharset)
	return generator.Generate()
}

// GenerateLowerCase generates a string of random lowercase latin characters
func GenerateLowerCase(length int) string {
	generator := NewGeneratorFromString(length, LatinLowerCaseCharset)
	return generator.Generate()
}

// GenerateUpperCase generates a string of random uppercase latin characters
func GenerateUpperCase(length int) string {
	generator := NewGeneratorFromString(length, LatinUpperCaseCharset)
	return generator.Generate()
}
