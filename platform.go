package emoji

import (
	"fmt"
	"strings"
)

// Constants that represent different emoji platforms.
type Platform int

const (
	PlatformNone Platform = iota
	PlatformApple
	PlatformGoogle
	PlatformTwitter
	PlatformFacebook
)

func (p Platform) MarshalText() ([]byte, error) {
	return []byte(p.String()), nil
}

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
