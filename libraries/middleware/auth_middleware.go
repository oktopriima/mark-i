/*
 * Copyright (c) 2019
 * Created at 5/26/19 12:24 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type TokenStructure struct {
	UserID int64  `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

type TokenResponse struct {
	AccessToken string  `json:"access_token"`
	TokenType   string  `json:"token_type"`
	ExpiredIn   float64 `json:"expired_in"`
	ExpiredAt   int64   `json:"expired_at"`
}

type customAuth struct {
	signature []byte
}

type CustomAuth interface {
	GenerateToken(data TokenStructure) (response TokenResponse, err error)
}

func NewCustomAuth(signature []byte) CustomAuth {
	return &customAuth{signature}
}

func (cAuth *customAuth) GenerateToken(data TokenStructure) (response TokenResponse, err error) {
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)

	expiredIn := time.Hour * (24 * 7)
	expiredAt := time.Now().Add(time.Hour * (24 * 7))

	myCrypt, err := bcrypt.GenerateFromPassword([]byte("my-password"), 8)
	if err != nil {
		return
	}

	claims["user_id"] = data.UserID
	claims["email"] = data.Email
	claims["role"] = data.Role
	claims["hash"] = string(myCrypt)
	claims["exp"] = expiredIn

	tokenString, _ := token.SignedString(cAuth.signature)

	response.AccessToken = tokenString
	response.TokenType = "Bearer"
	response.ExpiredAt = expiredAt.Unix()
	response.ExpiredIn = expiredIn.Seconds()

	return
}
