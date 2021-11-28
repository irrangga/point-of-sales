package middleware

import (
	"log"
	"net/http"
	"time"

	"point-of-sales/config/db"
	"point-of-sales/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:           "point-of-sales",
		Key:             []byte("secret key"),
		Timeout:         time.Hour,
		MaxRefresh:      time.Hour,
		IdentityKey:     IdentityKey,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
		Authenticator:   Authenticator,
		Authorizator:    Authorizator,
		Unauthorized:    Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware
}

var IdentityKey = "username"

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*models.Merchant); ok {
		return jwt.MapClaims{
			"email":     v.Email,
			IdentityKey: v.Username,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &models.Merchant{
		Email:    claims["email"].(string),
		Username: claims[IdentityKey].(string),
	}
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var login models.Login
	if err := c.ShouldBind(&login); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	username := login.Username
	password := login.Password

	usr := models.Merchant{}
	db.DB.First(&usr, "username = ? && password = ?", username, password)

	if username == usr.Username && password == usr.Password {
		return &models.Merchant{
			Email:    usr.Email,
			Username: usr.Username,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func Authorizator(data interface{}, c *gin.Context) bool {
	claims := jwt.ExtractClaims(c)
	if v, ok := data.(*models.Merchant); ok && v.Username == claims[IdentityKey] {
		return true
	}

	return false
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusUnauthorized, models.ResponseErrorCustom{
		Status:  http.StatusUnauthorized,
		Message: message,
	})
}
