package main

import (
	"github.com/arkaramadhan/its-vo/common/middleware"
	"github.com/arkaramadhan/its-vo/user-service/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware.CORS())

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.Run(":8084")
}
