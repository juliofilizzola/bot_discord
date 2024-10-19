package discord

import "strings"

func GetColorByString(input string) int {
	switch {
	case strings.Contains(input, "hot") || strings.Contains(input, "hotfix"):
		return ColorRed
	case strings.Contains(input, "fix"):
		return ColorOrange
	case strings.Contains(input, "feat") || strings.Contains(input, "feature"):
		return ColorGreen
	default:
		return ColorWhite
	}
}
