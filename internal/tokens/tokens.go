package tokens

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("my_super_secret_key") //Please never store secret in the code for the production setup

const tokenName = "jpjwt"

type Claims struct {
	UserId string `json:"user_id"`
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

func GenerateToken(userId int, roles []string) (string, time.Time, error){
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		UserId: strconv.Itoa(userId),
		Roles: roles,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expirationTime, nil
}

func GenerateAndAddTokenToCookie(userId int, roles [] string, w http.ResponseWriter) error {
	token, expirationTime, err := GenerateToken(userId, roles)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name: tokenName,
		Value: token,
		Expires: expirationTime,
		Path: "/",
	})

	return nil
}

func UpdateJWTandSaveToCookie(w http.ResponseWriter, r * http.Request) error {
	userId, err  := strconv.Atoi(r.Context().Value("userId").(string))
	if err != nil {
		log.Println("can't retrieve userId from request's context", err)
		return err
	}

	roles := r.Context().Value("roles").([]string)

	err = GenerateAndAddTokenToCookie(userId, roles, w)
	if err != nil {
		log.Println("can't save token into cookie ", err)
		return err
	}

	return nil
}

// user should contain at least one required role
func IsAuthenticated(r *http.Request, roles ... string) (*http.Request, bool) {
	cookie, err := r.Cookie(tokenName)
	if err != nil {
		log.Println("Error extracting cokie", err)
		return r, false
	}

	claims := &Claims{}

	_, err = jwt.ParseWithClaims(cookie.Value, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		log.Println("failed to parse jwt token ", err)
		return r, false
	}

	if len(roles) > 0 {
		if len(claims.Roles) == 0 {
			return r, false
		}

		userRoles := make(map[string]bool)
		for _, userRole := range roles {
			userRoles[userRole] = true 
		}

		containsRole := false 
		for _, role := range claims.Roles {
			if _, ok := userRoles[role]; ok {
				containsRole = true
				break
			}
		}

		if !containsRole {
			return r, false
		}
	}

	newR := r.WithContext(context.WithValue(r.Context(), "userId", claims.UserId))
	newR = newR.WithContext(context.WithValue(newR.Context(), "roles", claims.Roles))

	return newR, true
}

