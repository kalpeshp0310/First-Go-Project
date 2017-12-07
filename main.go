package main

import (
	"github.com/kalpeshp0310/hellogo/models"
	"github.com/kalpeshp0310/hellogo/web"
)

func main() {
	db := models.ConnectDB()
	defer db.Close()
	db.AutoMigrate(&models.List{}, &models.User{}, &models.Item{})
	web.StartServer()
}
