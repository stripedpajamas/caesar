package runes

// UpperMin is the lower bound of an uppercase alphabetic ASCII character
const UpperMin = 'A'

// UpperMax is the upper bound of an uppercase alphabetic ASCII character
const UpperMax = 'Z'

// LowerMin is the lower bound of a lowercase alphabetic ASCII character
const LowerMin = 'a'

// LowerMax is the upper bound of a lowercase alphabetic ASCII character
const LowerMax = 'z'

// IsLetter return true if the input rune is an alphabetic ASCII character
func IsLetter(r rune) bool {
	return IsLower(r) || IsUpper(r)
}

// IsUpper returns true if the input rune is an
// uppercase alphabetic ASCII character
func IsUpper(r rune) bool {
	return (r >= UpperMin && r <= UpperMax)
}

// IsLower returns true if the input rune is a
// lowercase alphabetic ASCII character
func IsLower(r rune) bool {
	return (r >= LowerMin && r <= LowerMax)
}

// ToUpper returns the uppercase value of the rune passed in
func ToUpper(r rune) rune {
	if IsUpper(r) {
		return r
	}
	return r - 32
}

// ToUpper returns the lowercase value of the rune passed in
func ToLower(r rune) rune {
	if IsLower(r) {
		return r
	}
	return r + 32
}
