/*
 * Copyright (c) 2019
 * Created at 5/20/19 10:15 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config interface {
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetStrings(key string) []string
	GetStringSlice(key string) []string
	Init(string)
}

type viperConfig struct{}

func (v *viperConfig) Init(prefix string) {
	viper.SetEnvPrefix(`go-clean`)
	viper.AutomaticEnv()

	osEnv := os.Getenv("OS_ENV")
	env := "env"
	if osEnv != "" {
		env = osEnv
	}

	if prefix != "" {
		env = prefix + "." + env
	}

	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigFile(env + `.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (v *viperConfig) GetStringSlice(key string) (c []string) {
	c = viper.GetStringSlice(key)
	return
}

func (v *viperConfig) GetStrings(key string) (c []string) {
	val := viper.GetString(key)
	c = strings.Split(val, ",")
	return
}

func NewConfig() Config {
	v := &viperConfig{}
	v.Init("")
	return v
}

func NewWithPrefix() Config {
	v := &viperConfig{}
	prefix := v.GetString("prefix")
	fmt.Println("Init config with prefix", prefix)
	v.Init(prefix)
	return v
}
