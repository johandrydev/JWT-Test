package token

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

var (
	privateKey *rsa.PrivateKey
	// PublicKey is used to validate the token
	PublicKey *rsa.PublicKey
)

func init() {
	privateBytes, err := ioutil.ReadFile("./private.rsa")
	if err != nil {
		log.Fatal("Private file could not be read", err)
	}

	publicBytes, err := ioutil.ReadFile("./public.rsa.pub")
	if err != nil {
		log.Fatal("I can't read the public file", err)
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		log.Fatal("Could not parse privateKey", err)
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		log.Fatal("Could not parse PublicKey", err)
	}
}

// Claim struct for create token access
type Claim struct {
	User interface{} `json:"user"`
	jwt.StandardClaims
}

// GenerateJWT Creates the user's token
func GenerateJWT(user interface{}, time time.Time) (string, error) {
	claims := Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(privateKey)
	return tokenString, err
}

// GenerateTokenPair Creates the user's access and refresh tokenpa
func GenerateTokenPair(user interface{}) (map[string]string, error) {
	token, err := GenerateJWT(user, time.Now().Add(time.Minute*5))
	if err != nil {
		return nil, err
	}

	tokenRf, err := GenerateJWT(user, time.Now().Add(time.Hour*24))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  token,
		"refresh_token": tokenRf,
	}, nil
}

// ValidateToken Middleware for validate token
func ValidateToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := request.ParseFromRequestWithClaims(
			r,
			request.OAuth2Extractor,
			&Claim{},
			func(t *jwt.Token) (interface{}, error) {
				return PublicKey, nil
			},
		)

		if err != nil {
			fmt.Println(err)
			code := http.StatusUnauthorized
			switch err.(type) {
			case *jwt.ValidationError:
				vError := err.(*jwt.ValidationError)
				switch vError.Errors {
				case jwt.ValidationErrorExpired:
					w.WriteHeader(code)
					w.Write([]byte("Your token has expired"))
					return
				case jwt.ValidationErrorSignatureInvalid:
					w.WriteHeader(code)
					w.Write([]byte("Token signature does not match"))
					return
				default:
					w.WriteHeader(code)
					w.Write([]byte("Your token is not valid"))
					return
				}
			}
		}

		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Your token is not valid"))
			return
		}
	}
}
