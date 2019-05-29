/*
 * Copyright (c) 2019
 * Created at 5/20/19 10:36 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package container

import (
	"github.com/oktopriima/mark-i/app/context/auth"
	"github.com/oktopriima/mark-i/app/context/role"
	"github.com/oktopriima/mark-i/app/context/user"
	"go.uber.org/dig"
)

func BuildUsecaseProvider(container *dig.Container) *dig.Container {
	/** register your use case here with this format */
	var err error

	if err = container.Provide(auth.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(user.NewUsecase); err != nil {
		panic(err)
	}

	if err = container.Provide(role.NewUsecase); err != nil {
		panic(err)
	}
	return container
}
