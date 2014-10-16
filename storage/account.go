package storage

import (
	"labix.org/v2/mgo"
	_ "github.com/harboe/massive-avenger/model"
)

type AccountRepository struct {

}

func (repo AccountRepository) Post(usr model.User) (model.User, error) {

	connectTable("users", func(c *mgo.Collection) {
		err := c.Insert(usr);
		return usr, err;
	});

}

func (repo AccountRepository) Put(usr model.User) (model.UserModel, error) {

}