package discord

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetColorByStringReturnsRedForHotfix(t *testing.T) {
	result := GetColorByString("hotfix/branch")
	assert.Equal(t, ColorRed, result)
}

func TestGetColorByStringReturnsRedForHot(t *testing.T) {
	result := GetColorByString("hot/branch")
	assert.Equal(t, ColorRed, result)
}

func TestGetColorByStringReturnsOrangeForFix(t *testing.T) {
	result := GetColorByString("fix/branch")
	assert.Equal(t, ColorOrange, result)
}

func TestGetColorByStringReturnsGreenForFeature(t *testing.T) {
	result := GetColorByString("feature/branch")
	assert.Equal(t, ColorGreen, result)
}

func TestGetColorByStringReturnsGreenForFeat(t *testing.T) {
	result := GetColorByString("feat/branch")
	assert.Equal(t, ColorGreen, result)
}

func TestGetColorByStringReturnsWhiteForUnknownString(t *testing.T) {
	result := GetColorByString("unknown/branch")
	assert.Equal(t, ColorWhite, result)
}
