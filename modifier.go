package emoji

import "fmt"

// Modifier is a string representation of an emoji modifier sequence.
type Modifier int

const (
	// SkinToneNone means no modifier was set.
	SkinToneNone Modifier = iota
	// SkinToneLight represents a light skin tone 👋🏻.
	SkinToneLight
	// SkinToneMediumLight represents a medium light skin tone 👋🏼.
	SkinToneMediumLight
	// SkinToneMedium represents a medium skin tone 👋🏽.
	SkinToneMedium
	// SkinToneMediumDark represents a medium dark skin tone 👋🏾.
	SkinToneMediumDark
	// SkinToneDark represents a dark skin tone 👋🏿.
	SkinToneDark
)

// NewModifier creates a modifier from a string.
// An empty string is interpreted as `SkinToneNone`.
func NewModifier(text string) (Modifier, error) {
	m := SkinToneNone
	return m, m.UnmarshalText([]byte(text))
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// This function also determines how a modifier is unmarshaled from JSON.
func (m *Modifier) UnmarshalText(text []byte) error {
	switch string(text) {
	case "", "none":
		*m = SkinToneNone
	case "1F3FB", "light":
		*m = SkinToneLight
	case "1F3FC", "medium_light":
		*m = SkinToneMediumLight
	case "1F3FD", "medium":
		*m = SkinToneMedium
	case "1F3FE", "medium_dark":
		*m = SkinToneMediumDark
	case "1F3FF", "dark":
		*m = SkinToneDark
	default:
		*m = SkinToneNone
		return fmt.Errorf("unrecognized modifier sequence %s", string(text))
	}
	return nil
}

// MarshalText implements the encoding.TextMarshaler interface.
// This function also determines how a modifier is marshaled to JSON.
func (m Modifier) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

// String implements fmt.Stringer and returns a string representation of the modifier.
func (m Modifier) String() string {
	switch m {
	case SkinToneNone:
		return "none"
	case SkinToneLight:
		return "light"
	case SkinToneMediumLight:
		return "medium_light"
	case SkinToneMedium:
		return "medium"
	case SkinToneMediumDark:
		return "medium_dark"
	case SkinToneDark:
		return "dark"
	default:
		return "unknown"
	}
}

// Unicode returns the sequence of unicode runes that represent the modifier.
func (m Modifier) Unicode() []rune {
	switch m {
	case SkinToneLight:
		return []rune{0x1F3FB}
	case SkinToneMediumLight:
		return []rune{0x1F3FC}
	case SkinToneMedium:
		return []rune{0x1F3FD}
	case SkinToneMediumDark:
		return []rune{0x1F3FE}
	case SkinToneDark:
		return []rune{0x1F3FF}
	default:
		return nil
	}
}
