package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type Person struct {
	Id    int
	Email string
	RegId string
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Person))
	beego.Info("Added model Person")
}

func Test() {
	beego.Info("start insert")
	p := Person{Email:"test@gmail.com"}
	beego.Info("insert p", p, Orm)
	id, err := Orm.Insert(&p)
	beego.Info("err", err)
	if err != nil {
		panic(err)
	}
	beego.Info("new Id: ", id)
}
