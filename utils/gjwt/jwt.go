package gjwt

import (
	"github.com/Heqiaomu/hqmGframe/model/config"
	"github.com/Heqiaomu/hqmGframe/utils/gerror"
	"github.com/pkg/errors"
	"time"
)
import "github.com/dgrijalva/jwt-go"

type JWTKey struct {
	Key []byte
}

type CustomClaims struct {
	UserId    string        `json:"userid"`
	Name      string        `json:"username"`
	ClientIP  string        `json:"clientip"`
	Timestamp int64         `json:"timestamp"`
	Expire    time.Duration `json:"expire"`
	jwt.StandardClaims
}

func KeyIn() *JWTKey {
	if config.GetJWT().JwtTokenSignKey == "" {
		return &JWTKey{Key: []byte("www.heqiaomu123.com")}
	}
	return &JWTKey{Key: []byte(config.GetJWT().JwtTokenSignKey)}
}

func (j *JWTKey) Token(claims CustomClaims) (string, error) {
	// 生成jwt格式的header、claims 部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 继续添加秘钥值，生成最后一部分
	return token.SignedString(j.Key)
}

func (j *JWTKey) ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.Key, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New(gerror.ErrorsTokenMalFormed)
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New(gerror.ErrorsTokenNotActiveYet)
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// 如果 TokenExpired ,只是过期（格式都正确），我们认为他是有效的，接下可以允许刷新操作
				token.Valid = true
				goto labelHere
			} else {
				return nil, errors.New(gerror.ErrorsTokenInvalid)
			}
		}
	}
labelHere:
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New(gerror.ErrorsTokenInvalid)
	}
}

func (j *JWTKey) Refresh(tokenStr string) (string, error) {
	if CustomClaims, err := j.ParseToken(tokenStr); err == nil {
		CustomClaims.ExpiresAt = time.Now().Add(CustomClaims.Expire).Unix()
		return j.Token(*CustomClaims)
	} else {
		return "", err
	}
	return "", nil
}
