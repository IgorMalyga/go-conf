package middleware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"go-conf/config"
	"go-conf/models"
	"go-conf/validators"
)

var identityKey = "Email"
var db_con = db.GetDB()

func AuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Email: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals validators.LoginValidator
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			var us models.User
			err := db_con.First(&us, "email = ?", loginVals.Email)

			if err.Error == nil && us.CheckPassword(loginVals.Password) {
				return &us, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*models.User); ok { //in v stores user identity (email)
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	return authMiddleware
}
