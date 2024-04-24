package helper

import (
	"DynamicStockManagmentSystem/internal/api/rest/responses"
	"DynamicStockManagmentSystem/internal/domain"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

type Auth struct {
	Secret string
}

func SetupAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
}

func (a Auth) CreateHashedPassword(p string) (string, error) {
	if len(p) < 6 {
		return "", errors.New("password length should be at least 6 characters long")
	}

	hashP, err := bcrypt.GenerateFromPassword([]byte(p), 10)

	if err != nil {
		return "", errors.New("password hash failed")
	}

	return string(hashP), nil
}

func (a Auth) GenerateToken(id string, userName string, firstName string, lastName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":       id,
		"firstName": firstName,
		"lastName":  lastName,
		"username":  userName,
		"exp":       time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(a.Secret))

	if err != nil {
		return "", errors.New("unable to signed the token")
	}

	return tokenStr, nil
}

func (a Auth) VerifyPassword(pP string, hP string) error {
	if len(pP) < 6 {
		return errors.New("password length should be at least 6 characters long")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hP), []byte(pP))

	if err != nil {
		return errors.New("wrong password")
	}

	return nil
}

func (a Auth) VerifyToken(t string) (domain.User, error) {
	tokenArr := strings.Split(t, " ")
	if len(tokenArr) != 2 {
		return domain.User{}, nil
	}

	tokenStr := tokenArr[1]

	if tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New("invalid token")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unknown signing method %v", token.Header)
		}
		return []byte(a.Secret), nil
	})

	if err != nil {
		return domain.User{}, errors.New("invalid signing method")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return domain.User{}, errors.New("token is expired")
		}

		user := domain.User{}
		user.ID, err = primitive.ObjectIDFromHex(claims["_id"].(string))
		if err != nil {
			return domain.User{}, errors.New("invalid user ID format")
		}
		user.FirstName = claims["firstName"].(string)
		user.LastName = claims["lastName"].(string)
		user.Username = claims["username"].(string)

		return user, nil
	}

	return domain.User{}, errors.New("token verification failed")
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {
	authHeaders := ctx.GetReqHeaders()["Authorization"]
	var authHeader string
	if len(authHeaders) > 0 {
		authHeader = authHeaders[0]
	} else {
		return responses.NewErrorResponse(ctx, http.StatusUnauthorized, "authorization header missing")
	}

	user, err := a.VerifyToken(authHeader)

	if err == nil {
		ctx.Locals("user", user)
		return ctx.Next()
	} else {
		return responses.NewErrorResponse(ctx, http.StatusUnauthorized, "authorization failed")
	}
}

func (a Auth) GetCurrentUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user")
	return user.(domain.User)
}
