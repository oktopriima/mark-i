/*
 * Copyright (c) 2019
 * Created at 5/28/19 7:32 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package main

import (
	"fmt"

	"github.com/oktopriima/mark-i/app/config"
	"github.com/oktopriima/mark-i/libraries/middleware"
	"github.com/oktopriima/mark-i/router"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

var _ = dig.Name

func main() {
	var err error

	myrole := make(map[string][]string)
	myrole[config.ADMIN] = []string{config.ADMIN}
	myrole[config.MERCHANT] = []string{config.ADMIN, config.MERCHANT}
	myrole[config.CONSUMER] = []string{config.ADMIN, config.MERCHANT, config.CONSUMER}

	middleware.InitRole(myrole)
	middleware.InitJWTMiddlewareCustom([]byte(config.SIGNATURE), jwt.SigningMethodHS512)

	gin.SetMode(gin.DebugMode)

	c := NewBuildContainer()

	if err = c.Invoke(router.InvokeRoute); err != nil {
		panic(err)
	}

	if err = c.Provide(NewRoute); err != nil {
		panic(err)
	}

	if err = c.Invoke(func(s *ServerRoute) {
		s.Run()
	}); err != nil {
		fmt.Println(err)
	}
}
