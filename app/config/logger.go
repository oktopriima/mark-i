/*
 * Copyright (c) 2019
 * Created at 5/27/19 1:49 AM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package config

import (
	"log"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func Logger(data ...interface{}) {
	log.Println(flagInfo(), data)
}

func Error(e error) error {
	log.Println(flagErr(), e)
	return e
}

func NewError(ds ...string) (e error) {
	d := ""
	for _, s := range ds {
		d += s + " "
	}
	e = errors.New(strings.Trim(d, " "))
	log.Println(flagErr(), d)
	return
}

func NewReturnError(ds ...string) (error, interface{}) {
	d := ""
	for _, s := range ds {
		d += s + " "
	}
	e := errors.New(strings.Trim(d, " "))
	log.Println(flagErr(), d)
	return e, nil
}

func NewErrors(d ...string) (es []error) {
	for _, s := range d {
		e := errors.New(s)
		es = append(es, e)
		log.Println(flagErr(), s)
	}
	return
}

func flagErr() string {
	return "[ERROR-" + time.Now().Format("20060102-150405.0000") + "]"
}

func flagInfo() string {
	return "[INFO-" + time.Now().Format("20060102-150405.0000") + "]"
}
