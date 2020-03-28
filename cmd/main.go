package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lbrulet/Go-api-test/api"
	"github.com/lbrulet/Go-api-test/pkg/user"
	"log"
)

func main() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=root dbname=database_test password=root sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := user.NewPostgresRepository(db)

	userService := user.NewService(userRepository)

	router := gin.Default()
	httpService := api.NewHttpService(router, userService)
	httpService.SetupRouter()

	_ = httpService.Router().Run(":8080")
}
