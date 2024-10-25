package main

import (
	"github.com/arkaramadhan/its-vo/common/middleware"
	"github.com/arkaramadhan/its-vo/project-service/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(middleware.CORS())

	// ********** Middleware ********** //
	r.Use(middleware.TokenAuthMiddleware())
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// ********** Project ********** //
	r.POST("/Project", controllers.ProjectCreate)
	r.PUT("/Project/:id", controllers.ProjectUpdate)
	r.GET("/Project", controllers.ProjectIndex)
	r.GET("/Project/:id", controllers.ProjectShow)
	r.DELETE("/Project/:id", controllers.ProjectDelete)
	r.GET("/exportProject", controllers.ExportProjectHandler)
	r.POST("/uploadProject", controllers.ImportExcelProject)

	r.POST("/uploadFileProject", controllers.UploadHandlerProject)
	r.GET("/downloadProject/:id/:filename", controllers.DownloadFileHandlerProject)
	r.DELETE("/deleteProject/:id/:filename", controllers.DeleteFileHandlerProject)
	r.GET("/filesProject/:id", controllers.GetFilesByIDProject)

	r.Run(":8086")
}
