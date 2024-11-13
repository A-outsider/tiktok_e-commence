package password

import "regexp"

// CheckPassword 校验密码安全性
func CheckPassword(password string) bool {
	var (
		hasMinLength   = len(password) >= 6
		hasUpperCase   = regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLowerCase   = regexp.MustCompile(`[a-z]`).MatchString(password)
		hasNumber      = regexp.MustCompile(`[0-9]`).MatchString(password)
		hasSpecialChar = regexp.MustCompile(`[\~\!\?\@\#\$\%\^\&\*\_\-\+\=\(\)\[\]\{\}\>\<\/\\\|\"\'\.\,\:\;]`).MatchString(password)
	)
	return hasMinLength && hasUpperCase && hasLowerCase && (hasNumber || hasSpecialChar)
}
