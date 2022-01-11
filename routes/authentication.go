package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"wishlists/database"
	"wishlists/helpers"
	"wishlists/models"

	"golang.org/x/crypto/bcrypt"
)

const secretKey string = "w1shl1st_4pp_s3cr3t"

func Register(rw http.ResponseWriter, r *http.Request) {
	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	user := models.User{
		Name: data.Name,
		Email: data.Email,
		Password: string(hashedPassword),
	}

	database.DB.Create(&user)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(user)
}

func Login(rw http.ResponseWriter, r *http.Request) {
	var data models.User
	err :=json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	var user models.User

	database.DB.Where("email = ?", data.Email).First(&user)

	if user.ID == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, "Invalid username or password.")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, "Invalid username or password.")
		return
	}

	expiresTime := time.Now().Add(time.Hour * 24)
	tokenString, err := helpers.CreateUserToken(user.ID, expiresTime)

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
	rw.Write([]byte{})
	return
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(models.User)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(user)
	return
}