package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lbrulet/Go-api-test/api"
	"github.com/lbrulet/Go-api-test/api/handler"
	"github.com/lbrulet/Go-api-test/pkg/user"
	"log"
)

var (
	port   string
	dbHost string
)

func init() {
	flag.StringVar(&port, "port", "8080", "use to define the http port (ex: 8080)")
	flag.StringVar(&dbHost, "db_host", "localhost", "use to define the database host (ex: 8080)")
}

func main() {

	flag.Parse()

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=root dbname=database_test password=root sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := user.NewPostgresRepository(db)

	userService := user.NewService(userRepository)

	userEndpointService := handler.NewUserEndpointService()

	router := gin.New()
	router.Use(gin.Recovery())

	httpService := api.NewHttpService(router, userService, userEndpointService)
	httpService.SetupRouter()

	_ = httpService.Router().Run(":" + port)
}
