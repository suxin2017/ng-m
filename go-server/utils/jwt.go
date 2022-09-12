package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	jwt.RegisteredClaims
	UserId uint `json:"userId"`
}

var hmacSampleSecret []byte

func Generate(userId uint) (tokenString string, err error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
		UserId: userId,
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err = token.SignedString(hmacSampleSecret)
	println(tokenString)
	// fmt.Println(tokenString, err)
	// if claims, ok := ParseToken(tokenString); ok == nil {
	// 	fmt.Println(claims.UserId)
	// 	return
	// } else {
	// 	fmt.Println(ok.Error())
	// }
	return
}

func ParseToken(tokenString string) (claims *UserClaims, err error) {
	parseToken, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})

	if claims, ok := parseToken.Claims.(*UserClaims); ok {
		fmt.Printf("%v\n", claims.UserId)
		return claims, err
	} else {
		return nil, err
	}

}
