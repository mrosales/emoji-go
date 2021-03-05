package emoji

import (
	"fmt"
	"strings"
)

// Platform defines constants that represent different emoji platforms.
type Platform int

const (
	// PlatformNone represents no specific platform.
	PlatformNone Platform = iota
	// PlatformApple is the Apple Platform.
	PlatformApple
	// PlatformGoogle is the Google and Android platform.
	PlatformGoogle
	// PlatformTwitter is the Twitter platform.
	PlatformTwitter
	// PlatformFacebook is the Facebook platform.
	PlatformFacebook
)

// MarshalText implements the encoding.TextMarshaler interface.
// This function also determines how a modifier is marshaled to JSON.
func (p Platform) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// This function also determines how a modifier is unmarshaled from JSON.
func (p *Platform) UnmarshalText(text []byte) error {
	switch strings.ToLower(string(text)) {
	case "none":
		*p = PlatformNone
	case "apple":
		*p = PlatformApple
	case "google":
		*p = PlatformGoogle
	case "twitter":
		*p = PlatformTwitter
	case "facebook":
		*p = PlatformFacebook
	default:
		*p = PlatformNone
		return fmt.Errorf("unrecognized platform \"%s\"", string(text))
	}
	return nil
}

// String implements fmt.Stringer and returns a string representation of the platform.
func (p Platform) String() string {
	switch p {
	case PlatformNone:
		return "None"
	case PlatformApple:
		return "Apple"
	case PlatformGoogle:
		return "Google"
	case PlatformTwitter:
		return "Twitter"
	case PlatformFacebook:
		return "Facebook"
	default:
		return "Unknown"
	}
}
