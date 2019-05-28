/*
 * Copyright (c) 2019
 * Created at 5/20/19 10:21 PM
 * Created by oktoprima
 * Email octoprima93@gmail.com
 * Github https://github.com/oktopriima
 */

package config

import (
	"log"

	"gopkg.in/mgo.v2"
)

func NewMongoDBConfig(cfg Config) (error, *mgo.Database) {
	address := cfg.GetString(`mongodb.address`)
	database := cfg.GetString(`mongodb.database`)
	log.Println("mongo config", address, database)
	session, err := mgo.Dial(address)
	session.SetMode(mgo.Monotonic, true)
	return err, session.DB(database)
}
