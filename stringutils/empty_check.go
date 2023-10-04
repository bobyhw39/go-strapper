package stringutils

import (
	"fmt"
	genericutils "github.com/bobyhw39/go-strapper/generic"
	"regexp"
	"strconv"
)

func IsBlank(str string) bool {
	if str == "" {
		return true
	}
	if regexp.MustCompile(`^\s+$`).MatchString(str) {
		return true
	}
	return false
}

func IsPointerBlank(str *string) bool {
	if str == nil {
		return true
	}
	return IsBlank(*str)
}

func AssumeString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case *string:
		if genericutils.NilEmpty(v) == nil {
			return ""
		}
		return *v
	case int:
		return strconv.Itoa(v)
	case *int:
		if genericutils.NilEmpty(v) == nil {
			return "0"
		}
		return strconv.Itoa(*v)
	case float32:
		return fmt.Sprintf("%f", v)
	case *float32:
		if genericutils.NilEmpty(v) == nil {
			return "0.00"
		}
		return fmt.Sprintf("%f", *v)
	case float64:
		return fmt.Sprintf("%f", v)
	case *float64:
		if genericutils.NilEmpty(v) == nil {
			return "0.00"
		}
		return fmt.Sprintf("%f", *v)
	default:
		return ""
	}
}
