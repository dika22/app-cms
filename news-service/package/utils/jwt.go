package utils

import (
	"fmt"
	"news-service/package/config"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cast"
)

func ValidateJWT(tokenStr string, cfg *config.Config) (userID int, role int, err error) {
    var jwtKey = []byte(cfg.JwtSecretKey) 
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        // Pastikan metode signing benar
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return jwtKey, nil
    })

    if err != nil {
        return 0, 0, err
    }

    // Ambil claims
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Ambil user_id
        if uid, ok := claims["user_id"].(float64); ok {
            userID = cast.ToInt(uid)
        } else {
            return 0, 0, fmt.Errorf("invalid user_id")
        }

        // Ambil role
        if r, ok := claims["role"].(float64); ok {
            role = cast.ToInt(r)
        } else {
            return 0, 0, fmt.Errorf("invalid role")
        }

        return userID, role, nil
    }

    return 0, 0, fmt.Errorf("invalid token")
}
