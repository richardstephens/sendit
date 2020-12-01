package utils

import (
	"github.com/google/uuid"
	"strings"
)

func isWindowsDeviceFile(s string) bool {
	switch s {
	case "CON", "AUX", "COM1", "COM2", "COM3", "COM4", "LPT1",
			"LPT2", "LPT3", "PRN", "NUL":
				return true
	}
	return false
}

func SecureFilename(s string) string {
	lastrune := '\n'
	sanitized := strings.Map(func(r rune) rune {
		if (r >= 'A' && r <= 'Z') || (r >='a' && r <= 'z') || (r >= '0' && r <= '9')  || r == '-' {
			return r
		}
		if r == ' ' || r == '_'  {
			return '_'
		}
		if lastrune != '.' && r == '.' {
			return '.'
		}
		lastrune = r
		return -1
	}, s)
	if isWindowsDeviceFile(sanitized) || sanitized == "" {
		return "upload-" + uuid.New().String()
	} else {
		if len(sanitized) < 100 {
			return sanitized
		} else {
			return sanitized[:50] +"__truncated__" + sanitized[len(sanitized)-50:]
		}
	}
}
