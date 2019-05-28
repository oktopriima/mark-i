/*
 * Copyright (c) 2019
 * Created at 5/20/19 10:02 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package main

import (
	"github.com/oktopriima/mark-i/container"
	"go.uber.org/dig"
)

func NewBuildContainer() *dig.Container {
	c := dig.New()
	c = container.BuildConfigProvider(c)
	c = container.BuildRepositoryProvider(c)
	c = container.BuildUsecaseProvider(c)
	c = container.BuildHttpHandlerProvider(c)

	return c
}
