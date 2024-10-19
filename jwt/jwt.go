package jwt

import "github.com/golang-jwt/jwt/v5"

type Payload struct {
	PlayerID    string `json:"player_id"`
	ServerGroup string `json:"server_group"`
}

var (
	parser = jwt.NewParser()
)

func CreateToken(key []byte, payload Payload) (string, error) {
	claims := jwt.MapClaims{
		"player_id":    payload.PlayerID,
		"server_group": payload.ServerGroup,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CheckSignature(tokenString string, key []byte) (bool, error) {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ParseUnverified(tokenString string) (Payload, error) {
	claims := jwt.MapClaims{
		"player_id":    "",
		"server_group": "",
	}
	_, _, err := parser.ParseUnverified(tokenString, claims)
	if err != nil {
		return Payload{}, err
	}
	return Payload{
		PlayerID:    claims["player_id"].(string),
		ServerGroup: claims["server_group"].(string),
	}, nil
}

func ParseToken(tokenString string, key []byte) (Payload, error) {
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
