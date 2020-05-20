package main

import (
	"os"
	"time"

	_ "github.com/Group5-HCMUS/hasagi/config"
	"github.com/Group5-HCMUS/hasagi/pkg/allocationrepo"
	"github.com/Group5-HCMUS/hasagi/pkg/authservice"
	"github.com/Group5-HCMUS/hasagi/pkg/lchistoryrepo"
	"github.com/Group5-HCMUS/hasagi/pkg/middleware"
	"github.com/Group5-HCMUS/hasagi/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/phamvinhdat/httpclient"
	"github.com/spf13/viper"
)

func main() {
	// get env
	env := os.Getenv("env")
	if env == "" {
		env = "dev"
	}
	viper := viper.Sub(env)

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
	httpClient := httpclient.NewClient()
	authURL := viper.GetString("auth.url")
	authService := authservice.New(authURL, httpClient)
	maxDistance := viper.GetFloat64("alert.max_distance")
	maxTime := viper.GetInt64("alert.max_time")
	serviceRepo := service.NewRepository(maxDistance, time.Minute*time.Duration(maxTime),
		aLocationRepo, lcHistoryRepo)
	sv := service.NewService(serviceRepo)
	authMiddleWare := middleware.VerifyToken(authService)
	sv.Register(router.Group("/api/v1", authMiddleWare))

	// run
	port := os.Getenv("PORT")
	if port == "" {
		port = viper.GetString("port")
	}

	_ = router.Run(":" + port)
}
