package controller

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	knot "github.com/eaciit/knot/knot.v1"
)

type IController interface{}

type Api struct {
	controller *IController
}

func (c *Api) SetAPIJsonType(k *knot.WebContext) {
	k.Config.NoLog = true
	k.Config.OutputType = knot.OutputJson
	k.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// k.Config.Headers["Access-Control-Allow-Headers"] = "Content-Type,access-control-allow-origin, access-control-allow-headers"
}

func (c *Api) GetToken(k *knot.WebContext) interface{} {
	c.SetAPIJsonType(k)

	secret := k.QueryDef("secret", time.Now().String()) // serverside secret
	// exp, _ := strconv.ParseInt(k.QueryDef("exp", "1000"), 10, 0)
	claim := jwt.StandardClaims{
		// ExpiresAt: exp,
		Issuer: "i wanna be",
	}
	sign := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claim)
	token, err := sign.SignedString([]byte(secret))

	return map[string]interface{}{
		"secret": secret,
		"claims": claim,
		"token":  token,
		"err":    err,
	}
}

func (c *Api) CheckToken(k *knot.WebContext) interface{} {
	c.SetAPIJsonType(k)

	tokenString := k.QueryDef("token", "")
	secret := k.QueryDef("secret", "") // serverside secret
	// exp, _ := strconv.ParseInt(k.QueryDef("exp", "1000"), 10, 0)

	// claim := jwt.StandardClaims{
	// ExpiresAt: exp,
	// Issuer: "i wanna be",
	// }

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// fmt.Println(jwt.GetSigningMethod("HS256"))
		// fmt.Println(token.Method)
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	return map[string]interface{}{
		"token":       token,
		"err":         err,
		"secret":      secret,
		"tokenString": tokenString,
	}
}
