package Db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"Basic/Config"
)
func Link()(bool){
	db, err := gorm.Open("mysql", Config.MSQ)
	defer db.Close()
	if err != nil {
		log.Panic("mysql db connect faild --- " + err.Error())
		return false
	}
	return  true

}
