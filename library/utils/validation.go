package utils

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

func GenerateID() string {
	return uuid.New().String()
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func IsValidName(name string) bool {
	return len(name) >= 3 && len(name) <= 100
}

var (
	ErrMinLength    = errors.New("string length is less than minimum")
	ErrMaxLength    = errors.New("string length is greater than maximum")
	ErrPatternMatch = errors.New("string does not match pattern")
	ErrEnumValue    = errors.New("string value is not in enum")
)

var (
	ErrMinimum          = errors.New("number is less than minimum")
	ErrExclusiveMinimum = errors.New("number is less than or equal to exclusive minimum")
	ErrMaximum          = errors.New("number is greater than maximum")
	ErrExclusiveMaximum = errors.New("number is greater than or equal to exclusive maximum")
	ErrMultipleOf       = errors.New("number is not a multiple of given value")
)

func ValidateMinLength(value string, min int) error {
	if len(value) < min {
		return fmt.Errorf("%w: expected >= %d, got %d", ErrMinLength, min, len(value))
	}
	return nil
}

func ValidateMaxLength(value string, max int) error {
	if len(value) > max {
		return fmt.Errorf("%w: expected <= %d, got %d", ErrMaxLength, max, len(value))
	}
	return nil
}

func ValidatePattern(value string, pattern string) error {
	re := regexp.MustCompile(pattern)
	if !re.MatchString(value) {
		return fmt.Errorf("value '%s' does not match pattern '%s'", value, pattern)
	}
	return nil
}

func ValidateEnum(value string, enum []string) error {
	if enum == nil || len(enum) == 0 {
		return nil
	}
	for _, v := range enum {
		if v == value {
			return nil
		}
	}
	return fmt.Errorf("%w: allowed values are [%s]", ErrEnumValue, strings.Join(enum, ", "))
}

func ValidateMinimum(value float64, minimum float64) error {
	if value < minimum {
		return fmt.Errorf("%w: expected >= %v, got %v", ErrMinimum, minimum, value)
	}
	return nil
}

func ValidateExclusiveMinimum(value float64, exclusiveMinimum float64) error {
	if value <= exclusiveMinimum {
		return fmt.Errorf("%w: expected > %v, got %v", ErrExclusiveMinimum, exclusiveMinimum, value)
	}
	return nil
}

func ValidateMaximum(value float64, maximum float64) error {
	if value > maximum {
		return fmt.Errorf("%w: expected <= %v, got %v", ErrMaximum, maximum, value)
	}
	return nil
}

func ValidateExclusiveMaximum(value float64, exclusiveMaximum float64) error {
	if value >= exclusiveMaximum {
		return fmt.Errorf("%w: expected < %v, got %v", ErrExclusiveMaximum, exclusiveMaximum, value)
	}
	return nil
}

func ValidateMultipleOf(value float64, multipleOf float64) error {
	quotient := value / multipleOf
	if !isWholeNumber(quotient) {
		return fmt.Errorf("%w: expected multiple of %v, got %v", ErrMultipleOf, multipleOf, value)
	}
	return nil
}

func isWholeNumber(f float64) bool {
	const epsilon = 1e-9
	_, frac := math.Modf(f)
	return frac < epsilon || frac > 1-epsilon
}
