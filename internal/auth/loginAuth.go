package auth

import "encoding/base64"

func HashPassword(password string) string {
	return base64.StdEncoding.EncodeToString([]byte(password))
}
