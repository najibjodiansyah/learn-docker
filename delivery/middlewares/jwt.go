package middlewares

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte("R4HASIA"),
	})
}

func CreateToken(userId int, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = userId
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //token akan expired dalam waktu 1 jam
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("R4HASIA"))
}

func ExtractTokenUserId(e echo.Context) (int,error) {
	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userId := int(claims["id"].(float64))
		return userId,nil
	}
	return -1,fmt.Errorf("token tidak valid")
}

		// user := c.Get("user").(*jwt.Token)

        // if !user.Valid {
        //     return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
        // }

        // claims := user.Claims.(jwt.MapClaims)

        // //use float64 because its default data that provide by JWT, we cant use int/int32/int64/etc.
        // //MUST CONVERT TO FLOAT64, OTHERWISE ERROR (not ok)!
        // userID, ok := claims["id"].(float64)

        // fmt.Println("inject jwt with testing", int(userID))

        // if !ok {
        //     return c.JSON(http.StatusForbidden, common.ForbiddedRequest())
        // }