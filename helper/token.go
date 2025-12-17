package helper

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"technical-test/config"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

type JwtData struct {
	UserId string
	Exp    string
}

type JwtKeys struct {
	UserId string
	Exp    string
}

func GetJwtKeys() JwtKeys {
	return JwtKeys{
		UserId: "user_id",
		Exp:    "exp",
	}
}

func GenerateToken(env *config.EnvironmentVariable, userId string) (token string, err error) {

	// Load env
	tokenSecretKey := env.Token.SecretKey

	atClaims := jwt.MapClaims{}
	atClaims[GetJwtKeys().UserId] = userId
	atClaims[GetJwtKeys().Exp] = time.Now().Add(time.Hour * 24 * 1).Unix()

	log.Info().Interface("atClaims", atClaims).Msg("claims")

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err = at.SignedString([]byte(tokenSecretKey))

	if err != nil {
		return
	}

	return
}
func TokenExpired(token *jwt.Token) bool {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return true
	}

	exp := int64(claims[GetJwtKeys().Exp].(float64))
	expTime := time.Unix(exp, 0)
	now := time.Now()

	return now.After(expTime)
}

func TokenValid(r *http.Request, env *config.EnvironmentVariable) error {
	token, err := VerifyToken(r, env)
	if err != nil {
		return err
	}

	if token == nil || TokenExpired(token) {
		return errors.New("token is expired or invalid")
	}

	return nil
}

func VerifyToken(r *http.Request, env *config.EnvironmentVariable) (*jwt.Token, error) {
	tokenString := ExtractToken(r)

	if tokenString == "" {
		err := errors.New("token is required")
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(env.Token.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ExtractToken(r *http.Request) string {
	tokenBearer := r.Header.Get("Authorization")

	strArr := strings.Split(tokenBearer, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractTokenMetadata(r *http.Request, env *config.EnvironmentVariable) (*JwtData, error) {
	token, err := VerifyToken(r, env)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {

		userId := fmt.Sprintf("%s", claims[GetJwtKeys().UserId])
		exp := fmt.Sprintf("%s", claims[GetJwtKeys().Exp])

		return &JwtData{
			UserId: userId,
			Exp:    exp,
		}, nil
	}

	return nil, err
}

func GetDataFromToken(r *http.Request, env *config.EnvironmentVariable) (userId string, err error) {

	jwtData, err := ExtractTokenMetadata(r, env)
	if err != nil {
		log.Error().Err(err).Msg("failed to extract data from token")
		return
	}

	userId = jwtData.UserId
	log.Info().Interface("jwtData", jwtData).Str("userId", userId).Msg("Jwt Data")

	if userId == "" {
		err = errors.New("user_id from token is empty")
		return
	}

	return
}

func GenerateResetPasswordToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}

	return hex.EncodeToString(b)
}
