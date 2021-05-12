package middleware

import (
	"bee.com/conf"
	"bee.com/dao"
	"bee.com/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	//"github.com/dgrijalva/jwt-go"
)

var (
	//TokenExpired     = errors.New("Token is expired")

)


type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(conf.C.JWT.SigningKey),
	}
}

// 创建token
func (j *JWT) CreateToken(claims dao.CustomClaims)  (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(j.SigningKey)
}



func JWTAuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		token := c.Request.Header.Get("token")
		if token == "" {
			utils.ResError(c, "未登录或非法访问")
			c.Abort()
			return
		}

		//j := NewJWT()
		//claims, err := j.ParseToken(token)
		//if err != nil {
		//
		//}
	}
}

func (j *JWT) ParseToken(tokenString string) (*dao.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dao.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("That's not even a token")
			} else if ve.Errors & jwt.ValidationErrorExpired != 0{
				return nil, errors.New("Token is expired")
			} else if ve.Errors & jwt.ValidationErrorNotValidYet != 0{
				return nil, errors.New("Token not active yet")
			} else {
				return nil, errors.New("Couldn't handle this token")
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*dao.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("Couldn't handle this token")
	} else {
		return nil,  errors.New("Couldn't handle this token")
	}
}