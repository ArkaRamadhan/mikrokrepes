package main

import (
	"github.com/arkaramadhan/its-vo/common/initializers"
	"github.com/arkaramadhan/its-vo/controllers"
	"github.com/gin-gonic/gin"
)

func init() {

	db.LoadEnvVariables()
	db.ConnectToDB()

}

func main() {

	r := gin.Default()
	// Routes for Memo
	r.GET("/memo", controllers.MemoIndex)
	r.POST("/memo", controllers.MemoCreate)
	r.GET("/memo/:id", controllers.MemoShow)
	r.PUT("/memo/:id", controllers.MemoUpdate)
	r.DELETE("/memo/:id", controllers.MemoDelete)
	r.GET("/exportMemo", controllers.ExportMemoHandler)
	r.POST("/uploadMemo", controllers.ImportExcelMemo)
}
