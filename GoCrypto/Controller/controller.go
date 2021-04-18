package controller 

import (
	"log"
	"net/http"
	"context"
	"fmt"
	"encoding/json"
	"io/ioutil"


	"Users/Downloads/GoCrypto/Controller"
	"Users/Downloads/GoCrypto/Model"
	"github.com/gorilla/mux"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"github.com/labstack/echo"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user model.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	var res model.ResponseResult
	if err != nil {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

	/*
	Rest code will be one after connection to database
	*/

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	/*
	Rest code will be one after connection to database
	*/

}


func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte("secret"), nil
	})
	var result model.User
	var res model.ResponseResult
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result.Username = claims["username"].(string)
		result.FirstName = claims["firstname"].(string)
		result.LastName = claims["lastname"].(string)

		json.NewEncoder(w).Encode(result)
		return
	} else {
		res.Error = err.Error()
		json.NewEncoder(w).Encode(res)
		return
	}

}

func CreateEvent(c echo.context) error {
	e := new(Event)
	err := c.Bind(&e)
	if err != nil {
		return err
	}

	now := time.Now()
	e.CreationDate = now
	e.UpdationDate = now

	/*
	Connection to database yet to done
	*/

	return c.JSON(http.StatusCreated, e)
}
