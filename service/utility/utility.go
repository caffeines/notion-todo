package utility

// add IsValid date format function

import (
	"fmt"
	"regexp"
	"strings"
)

// IsValidDateFormat checks if the date is in DD-MM-YYYY format

func IsValidDateFormat(date string) bool {
	// Regular expression to match DD-MM-YYYY format
	const datePattern = `^(0[1-9]|[12][0-9]|3[01])-(0[1-9]|1[0-2])-(\d{4})$`
	re := regexp.MustCompile(datePattern)
	return re.MatchString(date)
}

func ConvertDateToYYYYMMDD(date string) string {
	// Check if the date is in DD-MM-YYYY format
	if !IsValidDateFormat(date) {
		return ""
	}
	// Split the date into components
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return ""
	}
	// Rearrange the components to YYYY-MM-DD format
	return fmt.Sprintf("%s-%s-%s", parts[2], parts[1], parts[0])
}
