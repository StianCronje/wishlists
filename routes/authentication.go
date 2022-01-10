package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"wishlists/database"
	"wishlists/models"

	"github.com/dgrijalva/jwt-go/v4"
	"golang.org/x/crypto/bcrypt"
)

const secretKey string = "w1shl1st_4pp_s3cr3t"

func Register(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, err.Error())
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), 14)
	user := models.User {
		Name: r.FormValue("name"),
		Email: r.FormValue("email"),
		Password: hashedPassword,
	}

	database.DB.Create(&user)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(user)
}

func Login(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	var user models.User

	database.DB.Where("email = ?", r.FormValue("email")).First(&user)

	if user.ID == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, "Invalid username or password.")
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(r.FormValue("password")))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, "Invalid username or password.")
		return
	}

	expiresTime := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.At(expiresTime),
	})

	secret := []byte(secretKey)
	tokenString, err := token.SignedString(secret)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Could not log in.")
		return
	}

	cookie := http.Cookie{
		Name: "jwt_token",
		Value: tokenString,
		Expires: expiresTime,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure: true,
		Path: "/",
	}

	http.SetCookie(rw, &cookie)

	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Success")
	return
}

func Logout(rw http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name: "jwt_token",
		Value: "",
		Expires: time.Now().Add(time.Hour * -1),
		HttpOnly: true,
	}

	http.SetCookie(rw, &cookie)
	
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintf(rw, "Success")
	return
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt_token")
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(rw, "Unauthenticated.")
		return
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(rw, "Unauthenticated. %v", err)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User
	database.DB.Where("id = ?", claims.Issuer).First(&user)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(user)
	return
}