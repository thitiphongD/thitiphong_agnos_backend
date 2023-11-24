package helpers

import "unicode"

func PasswordStrengthSteps(initPassword string) int {
	steps := 0

	if len(initPassword) < 6 {
		steps += 6 - len(initPassword)
	} else if len(initPassword) >= 20 {
		steps += len(initPassword) - 19
	}

	var hasLower, hasUpper, hasDigit bool
	for _, char := range initPassword {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	if !hasLower {
		steps++
	}
	if !hasUpper {
		steps++
	}
	if !hasDigit {
		steps++
	}

	for i := 0; i < len(initPassword)-2; i++ {
		if initPassword[i] == initPassword[i+1] && initPassword[i] == initPassword[i+2] {
			steps++
			break
		}
	}

	return steps
}
