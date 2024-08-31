package utils

import "github.com/golang-jwt/jwt/v5"

type Payload struct {
	PlayerID string `json:"player_id"`
}

func CreateJWTToken(key []byte, payload Payload) (string, error) {
	claims := jwt.MapClaims{
		"player_id": payload.PlayerID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseJWTToken(tokenString string, key []byte) (Payload, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return Payload{}, err
	}
	return Payload{
		PlayerID: claims["player_id"].(string),
	}, nil
}
