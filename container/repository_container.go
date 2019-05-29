/*
 * Copyright (c) 2019
 * Created at 5/20/19 10:31 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package container

import (
	"github.com/oktopriima/mark-i/domain/services"
	"go.uber.org/dig"
)

func BuildRepositoryProvider(container *dig.Container) *dig.Container {
	var err error

	if err = container.Provide(services.NewUserServices); err != nil {
		panic(err)
	}

	if err = container.Provide(services.NewUserSocialServices); err != nil {
		panic(err)
	}

	if err = container.Provide(services.NewRoleServices); err != nil {
		panic(err)
	}

	if err = container.Provide(services.NewRoleUserService); err != nil {
		panic(err)
	}

	return container
}
