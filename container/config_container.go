/*
 * Copyright (c) 2019
 * Created at 5/20/19 10:06 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package container

import (
	"github.com/oktopriima/mark-i/app/config"
	"github.com/oktopriima/mark-i/libraries/httpresponse"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.uber.org/dig"
	"gopkg.in/mgo.v2"
)

func BuildConfigProvider(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(func() []byte {
		return []byte("mark-one-apps")
	}); err != nil {
		panic(err)
	}

	if err = container.Provide(func() config.Config {
		return config.NewConfig()
	}); err != nil {
		panic(err)
	}

	if err = container.Provide(func(cfg config.Config) (error, *gorm.DB) {
		return config.NewMysqlConfig(cfg)
	}); err != nil {
		panic(err)
	}

	if err = container.Provide(func(cfg config.Config) (error, *mgo.Database) {
		return config.NewMongoDBConfig(cfg)
	}); err != nil {
		panic(err)
	}

	if err = container.Provide(func() httpresponse.RequestReader {
		return httpresponse.NewRequestReader()
	}); err != nil {
		panic(err)
	}

	if err = container.Provide(func() *gin.Engine {
		return gin.New()
	}); err != nil {
		panic(err)
	}

	return container
}
