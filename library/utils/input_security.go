package utils

import (
	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/microcosm-cc/bluemonday"
)

var (
	spaceRegex = regexp.MustCompile(`\s+`)
	tagRegex   = regexp.MustCompile(`<.*?>`)
)

func ContainsRole(roles []string, role string) bool {
	role = strings.ToLower(role)
	for _, r := range roles {
		if strings.ToLower(r) == role {
			return true
		}
	}
	return false
}

func SanitizeString(input string) string {
	clean := govalidator.StripLow(input, true)
	clean = tagRegex.ReplaceAllString(clean, "")
	clean = spaceRegex.ReplaceAllString(clean, " ")
	clean = strings.TrimSpace(clean)
	policy := bluemonday.StrictPolicy()
	clean = policy.Sanitize(clean)
	return clean
}
