package repositories

import(
	"gorm.io/gorm"
)

type dbConnection struct{
	connection *gorm.DB
}