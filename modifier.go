package emoji

import "fmt"

// Modifier is a string representation of an emoji modifier sequence.
type Modifier int

const (
	SkinToneNone        Modifier = iota // no explicit skin tone set
	SkinToneLight                       // SkinToneLight represents a light skin tone 👋🏻.
	SkinToneMediumLight                 // SkinToneMediumLight represents a medium light skin tone 👋🏼.
	SkinToneMedium                      // SkinToneMedium represents a medium skin tone 👋🏽.
	SkinToneMediumDark                  // SkinToneMediumDark represents a medium dark skin tone 👋🏾.
	SkinToneDark                        // SkinToneDark represents a dark skin tone 👋🏿.
)

func NewModifier(text string) (Modifier, error) {
	m := SkinToneNone
	return m, m.UnmarshalText([]byte(text))
}

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

func (m Modifier) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

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

func (m Modifier) Unicode() rune {
	switch m {
	case SkinToneLight:
		return 0x1F3FB
	case SkinToneMediumLight:
		return 0x1F3FC
	case SkinToneMedium:
		return 0x1F3FD
	case SkinToneMediumDark:
		return 0x1F3FE
	case SkinToneDark:
		return 0x1F3FF
	default:
		return 0
	}
}
