/*
 * Copyright (c) 2019
 * Created at 5/26/19 10:53 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package helper

import (
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/oktopriima/mark-i/libraries/middleware"
)

func ResponseToByte(response *http.Response) ([]byte, error) {
	var bodyBytes []byte
	bodyBytes, err := ioutil.ReadAll(response.Body)
	return bodyBytes, err
}

func RandString(n int) string {
	const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

func GetAuthenticatedUser(r *http.Request) (int64, error) {
	userID, err := middleware.ExtractToken(r, "user_id")
	if err != nil {
		return 0, err
	}
	return int64(userID.(float64)), nil
}
