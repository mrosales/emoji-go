package emoji

//go:generate go run ./internal/importer data.go

import (
	"fmt"
)

type Info struct {
	// Name is the canonical short name.
	Name string `json:"name"`
	// Category is the category of emoji.
	Category string `json:"category"`
	// PlainText is defined if the character has a canonical plaintext representation.
	PlainText string `json:"plain_text,omitempty"`
	// AlternateNames is a list of alternative string names or keywords
	AlternateNames []string `json:"alternate_names,omitempty"`
	// ImageData is an embedded struct with details about the actual image.
	ImageData
	// SkinVariations is the map of alternative emojis for different skin tones.
	SkinVariations map[Modifier]ImageData `json:"skin_variations,omitempty"`
}

type ImageData struct {
	// A hyphen separated sequence of hex-encoded codepoints.
	// Includes zero-width-joiner character for multi code sequences.
	Unified string `json:"unified"`
	// Character is the actual emoji character consisting of one or more codepoints.
	Character string `json:"character"`
	// SheetX is the X index of the image in the sprite sheet.
	SheetX int `json:"sheet_x"`
	// SheetY is the Y index of the image in the sprite sheet.
	SheetY int `json:"sheet_y"`
	// AddedIn is the unicode revision that the emoji was added in
	AddedIn string `json:"added_in,omitempty"`
	// PlatformSupport defines the supported platforms.
	PlatformSupport map[Platform]bool `json:"platform_support,omitempty"`
	// Obsoletes is set if the emoji replaces another emoji from an older Unicode revision.
	Obsoletes string `json:"obsoletes,omitempty"`
	// ObsoletedBy is set if the emoji is replaced by another emoji from a newer Unicode revision.
	ObsoletedBy string `json:"obsoleted_by,omitempty"`
}

// ImageForModifier returns the ImageData for the given emoji modifier sequence.
// Currently only skin tone modifications are supported.
func (i Info) ImageForModifier(mod Modifier) ImageData {
	switch mod {
	case SkinToneLight, SkinToneMediumLight, SkinToneMedium, SkinToneMediumDark, SkinToneDark:
		if modified, ok := i.SkinVariations[mod]; ok {
			return modified
		}
	default:
		break
	}
	return i.ImageData
}

// String representation of the modifier sequence.
func (i Info) String() string {
	return fmt.Sprintf("%s (%s - %s)", i.Character, i.Name, i.Unified)
}
