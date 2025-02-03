package helper

import (
	"GoCart/internal/domain"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

func SetupAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
}

func (a *Auth) CreateHashPassword(p string) (string, error) {

	if len(p) < 6 {
		return "", errors.New("length of password should be atleast 6")
	}

	hashP, err := bcrypt.GenerateFromPassword([]byte(p), 10)

	if err != nil {
		log.Errorf("Error from genrating password is %v", err)
		return "", errors.New("password hash failed")
	}

	return string(hashP), nil
}

func (a *Auth) GenerateToken(id uint, email, role string) (string, error) {
	if id == 0 || email == "" || role == "" {
		return "", errors.New("id,email or role can't be empty")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"expiry":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenStr, err := token.SignedString([]byte(a.Secret))
	if err != nil {
		return "", errors.New("unable to sign the token")
	}

	return tokenStr, nil
}

func (a *Auth) VerifyPassword(pP string, hP string) error {
	if len(pP) < 6 {
		return errors.New("passowrd cant be shorter that 6")
	}
	err := bcrypt.CompareHashAndPassword([]byte(hP), []byte(pP))
	if err != nil {
		return errors.New("password does not match")
	}

	return nil
}

func (a *Auth) VerifyToken(t string) (domain.User, error) {
	tokenArray := strings.Split(t, " ")
	if len(tokenArray) != 2 {
		return domain.User{}, nil
	}
	tokenStr := tokenArray[1]
	if tokenArray[0] != "Bearer" {
		return domain.User{}, errors.New("invalid token")
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown signing method %v", token.Header)
		}
		return []byte(a.Secret), nil
	})
	if err != nil {
		return domain.User{}, errors.New("token is expired")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expClaim, ok := claims["expiry"].(float64)
		if !ok {
			return domain.User{}, errors.New("invalid or missing expiration claim")
		}
		if float64(time.Now().Unix()) > expClaim {
			return domain.User{}, errors.New("token is expired")
		}

		user := domain.User{}
		user.ID = uint(claims["user_id"].(float64))
		user.Email = claims["email"].(string)
		user.UserType = claims["role"].(string)
		return user, nil
	}

	return domain.User{}, errors.New("token verification failed")
}

func (a *Auth) Authorize(ctx *fiber.Ctx) error {
	authHeaders := ctx.GetReqHeaders()["Authorization"]
	if len(authHeaders) == 0 {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  "missing Authorization header",
		})
	}
	user, err := a.VerifyToken(authHeaders[0])

	if err == nil && user.ID > 0 {
		ctx.Locals("user", user)
		return ctx.Next()
	}

	return ctx.Status(401).JSON(&fiber.Map{
		"message": "authorization failed",
		"reason":  err,
	})
}

func (a *Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")
	return user.(domain.User)
}

func (a Auth) GenerateCode() (int, error) {

	return RandomNumbers(6)
}
