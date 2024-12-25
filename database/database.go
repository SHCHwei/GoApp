package database

import (
    "gorm.io/gorm"
)

var MariaDB *gorm.DB

var Mongo MongoDB


func init(){
    mariaConnect()
    Mongo.connect()
}