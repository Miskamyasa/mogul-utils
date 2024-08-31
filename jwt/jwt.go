package jwt

import "github.com/golang-jwt/jwt/v5"

type Payload struct {
	PlayerID    string `json:"player_id"`
	ServerGroup string `json:"server_group"`
}

func CreateJWTToken(key []byte, payload Payload) (string, error) {
	claims := jwt.MapClaims{
		"player_id":    payload.PlayerID,
		"server_group": payload.ServerGroup,
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
		PlayerID:    claims["player_id"].(string),
		ServerGroup: claims["server_group"].(string),
	}, nil
}
