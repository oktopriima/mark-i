/*
 * Copyright (c) 2019
 * Created at 5/20/19 10:53 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package container

import (
	"github.com/oktopriima/mark-i/httphandler/auth"
	"github.com/oktopriima/mark-i/httphandler/role"
	"github.com/oktopriima/mark-i/httphandler/user"
	"go.uber.org/dig"
)

func BuildHttpHandlerProvider(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(auth.NewHandler); err != nil {
		panic(err)
	}

	if err = container.Provide(user.NewHandler); err != nil {
		panic(err)
	}

	if err = container.Provide(role.NewHandler); err != nil {
		panic(err)
	}

	return container
}
