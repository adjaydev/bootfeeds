package parser

import (
	"fmt"
	"time"
)

func ParseTimestamp(str string) (time.Time, error) {
	layouts := []string{
		time.RFC1123,           // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC1123Z,          // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC3339,           // "2006-01-02T15:04:05Z07:00"
		time.RFC3339Nano,       // "2006-01-02T15:04:05.999999999Z07:00"
		"2006-01-02 15:04:05",  // Common SQL format
		"02 Jan 2006 15:04:05", // "02 Jan 2025 15:04:05"
		"02-Jan-2006 15:04:05", // "02-Jan-2025 15:04:05"
		"02/01/2006 15:04:05",  // "02/01/2025 15:04:05"
	}

	var t time.Time
	var err error
	for _, layout := range layouts {
		t, err = time.Parse(layout, str)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unknown timestamp format: %s", str)
}
