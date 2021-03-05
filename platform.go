package emoji

// Constants that represent different emoji platforms.
type Platform int

const (
	PlatformNone Platform = iota
	PlatformApple
	PlatformGoogle
	PlatformTwitter
	PlatformFacebook
)

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
