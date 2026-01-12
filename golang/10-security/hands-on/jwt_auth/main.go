package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret Key (Keep this safe! Env val in real life)
var secretKey = []byte("super-secret-key")

// CreateToken generates a JWT
func CreateToken(username string) (string, error) {
	// 1. Create Claims (Payload)
	claims := jwt.MapClaims{
		"username": username,
		"role":     "admin",
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Expires in 1 hour
	}

	// 2. Create Token with Signing Method (HS256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3. Sign it with the secret
	return token.SignedString(secretKey)
}

// VerifyToken checks signature and expiry
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// Parse takes the token string and a "KeyFunc" that provides the verify key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validating the algo is crucial!
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func main() {
	// 1. Issue Token
	user := "phoenix_user"
	tokenStr, _ := CreateToken(user)
	fmt.Println("Generated Token:", tokenStr)

	// 2. Validate Token
	fmt.Println("--- Verifying ---")
	claims, err := VerifyToken(tokenStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("✅ Welcome %s! Your role is %s.\n", claims["username"], claims["role"])

	// 3. Tamper Check
	// If we change one char, it should fail
	badToken := tokenStr + "fake"
	_, err = VerifyToken(badToken)
	if err != nil {
		fmt.Println("🛡️ Tampered token rejected:", err)
	}
}
