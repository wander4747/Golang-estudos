package address

import "strings"

// TypeAddress = check type address
func TypeAddress(address string) string {
	validType := []string{"rua", "avenida", "rodovia"}

	firstWord := strings.Split(strings.ToLower(address), " ")[0]

	addressTypeValid := false

	for _, typeAddress := range validType {
		if typeAddress == firstWord {
			addressTypeValid = true
		}
	}

	if addressTypeValid {
		return strings.Title(firstWord)
	}

	return "Invalid type"
}
