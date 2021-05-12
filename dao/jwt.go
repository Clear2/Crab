package dao

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type CustomClaims struct {
	 UUID uuid.UUID
	 ID uint64
	 UserName string
	 AuthorityId string
	 BufferTime int64
	 jwt.StandardClaims
}
