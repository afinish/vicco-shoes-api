package main

import (
	"github.com/heavenmise/vicco-shoes-api/pkg/config"
	"github.com/heavenmise/vicco-shoes-api/pkg/models"
	"github.com/heavenmise/vicco-shoes-api/pkg/routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	if err != nil {
		panic(err)
	}
	defer config.DB.Close()

	config.DB.AutoMigrate(&models.Shoe{})

	r := routes.RegisterNewRouter()
	r.Run()
}