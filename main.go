package main

import (
	"log"
	"os"

	_ "github.com/Group5-HCMUS/hasagi/config"
	"github.com/Group5-HCMUS/hasagi/pkg/allocationrepo"
	"github.com/Group5-HCMUS/hasagi/pkg/lchistoryrepo"
	"github.com/Group5-HCMUS/hasagi/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func main() {
	// database
	db, err := gorm.Open(viper.GetString("database.dialect"),
		viper.GetString("database.url"))
	if err != nil {
		panic(err)
	}

	// repository
	aLocationRepo := allocationrepo.New(db)
	lcHistoryRepo := lchistoryrepo.New(db)

	// gin
	router := gin.Default()

	// service
	serviceRepo := service.NewRepository(aLocationRepo, lcHistoryRepo)
	sv := service.NewService(serviceRepo)
	sv.Register(router.Group("/api/v1"))

	// run
	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString("port")
	}

	log.Println("service running on port: ", port)
	_ = router.Run(":" + port)
}
