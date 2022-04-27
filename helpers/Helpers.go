package helpers

import (
	"unicode"
)

// func Verify(pass string, reqPass string) bool {
// 	byteHash := []byte(pass)
// 	err := bcrypt.CompareHashAndPassword(byteHash, []byte(reqPass))

// 	if err != nil {
// 		return false
// 	}

// 	return true
// }

// func GetTokenID(c echo.Context) int {
// 	userToken := c.Get("user").(*jwt.Token)
// 	claims := userToken.Claims.(jwt.MapClaims)
// 	idToken := claims["id"].(float64)
// 	id := int(idToken)

// 	return id
// }

// func VerifyToken(c echo.Context, user models.User) bool {
// 	reqToken := c.Request().Header.Get("Authorization")
// 	token := "Bearer " + user.Token
// 	if token != reqToken {
// 		return false
// 	}

// 	return true
// }

func IsInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
