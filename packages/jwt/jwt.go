package jwt

import (
	"errors"
	"os"
	"refrigerator/business"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(secret []byte, member business.Member) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["nbf"] = time.Now().Unix()
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	claims["permission"] = strconv.Itoa(member.Permission)
	claims["id"] = member.ID
	claims["iss"] = os.Getenv("JWT_ISSUER")
	claims["sub"] = "PYPL_TKN"

	s, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return s, nil
}

func VerifyJWT(secret []byte, memberToken string) (business.Member, error) {
	token, err := jwt.Parse(memberToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("jwt token specified cant be parsed")
		}
		return secret, nil
	})
	if err != nil {
		return business.Member{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return business.Member{}, errors.New("jwt token is invalid")
	}

	if claims["iss"] != os.Getenv("JWT_ISSUER") {
		return business.Member{}, errors.New("jwt token is unacceptable")
	}

	var member business.Member
	member.Permission, _ = strconv.Atoi(claims["permission"].(string))
	member.ID = claims["id"].(string)
	return member, nil
}
