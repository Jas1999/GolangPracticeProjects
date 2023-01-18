package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//variable for db connection
var (
	DBConn *gorm.DB
)
