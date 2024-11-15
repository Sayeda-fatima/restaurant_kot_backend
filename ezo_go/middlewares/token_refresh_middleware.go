package middlewares

// import (
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/labstack/echo/v4"
// )

// func TokenRefreshMiddleware(next echo.HandlerFunc) echo.HandlerFunc{
// 	return func(c echo.Context) error{

// 		if c.Get("user") == nil{
// 			return next(c)
// 		}

// 		user := c.Get("user").(*jwt.Token)
// 		claims := user.Claims.(jwt.MapClaims)
// 		exp := claims["exp"]
// 		id := claims["user_id"]

// 		if time.Unix(exp, 0).Sub(time.Now()) < 15*time.Minute{

// 			refreshToken, err := c.Cookie(refreshTokenString)

// 		}
// 	}
// }