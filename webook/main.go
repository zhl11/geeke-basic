package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)
import "github.com/zhl11/geeke-basic/webook/internal"

func main() {
	router := gin.Default()
	user := internal.NewUserHandler()
	router.Use(cors.Default())
	//router.Use(cors.Default())
	//cors.Default()
	// 先注册 CORS 中间件
	router.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}, // 确保包含 POST
		AllowHeaders:     []string{"Content-Type", "authorization"},
		ExposeHeaders:    []string{"Content-Type", "authorization"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			if strings.HasSuffix(origin, "http://localhost") {
				return true
			}

			return strings.Contains(origin, "yuming")
		},
		MaxAge: 12 * time.Hour,
	}))

	user.RegistryServer(router)
	router.Run(":8080")
}
