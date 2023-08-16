package controllers

import (
	"fmt"
	"math"
	"time"
)

func GenerateTOTP(secret string, currentTime time.Time) string {
	hmacValue := GenerateHash(secret, currentTime)
	offset := int(hmacValue[19] & 0x0f)
	binary := ((int(hmacValue[offset]) & 0x7f) << 24) |
		((int(hmacValue[offset+1] & 0xff)) << 16) |
		((int(hmacValue[offset+2] & 0xff)) << 8) |
		(int(hmacValue[offset+3] & 0xff))

	digits := 6
	mod := int(math.Pow10(digits))
	otp := binary % mod

	code := fmt.Sprintf("%0*d", digits, otp)

	return code
}

// Verify a TOTP code
func ValidateTOTP(code, secret string, currentTime time.Time) bool {
	generatedCode := GenerateTOTP(secret, currentTime)

	return code == generatedCode
}
