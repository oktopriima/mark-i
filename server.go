/*
 * Copyright (c) 2019
 * Created at 5/20/19 11:18 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package main

import (
	"log"

	"github.com/oktopriima/mark-i/app/config"
	"github.com/gin-gonic/gin"
)

type ServerRoute struct {
	cfg    config.Config
	engine *gin.Engine
}

func NewRoute(cfg config.Config, engine *gin.Engine) *ServerRoute {
	return &ServerRoute{cfg, engine}
}

func (s *ServerRoute) Run() {
	if err := s.engine.Run(s.cfg.GetString(`server.address`)); err != nil {
		log.Fatal(err)
	}
}
