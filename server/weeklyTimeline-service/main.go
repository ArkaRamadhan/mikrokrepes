package main

import (
	"github.com/arkaramadhan/its-vo/common/middleware"
	"github.com/arkaramadhan/its-vo/weeklyTimeline-service/controllers"
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

	// ********** Route Meeting Schedule ********** //
	r.GET("/meetingSchedule", controllers.MeetingListIndex)
	r.POST("/meetingSchedule", controllers.MeetingListCreate)
	r.GET("/meetingSchedule/:id", controllers.MeetingListShow)
	r.PUT("/meetingSchedule/:id", controllers.MeetingListUpdate)
	r.DELETE("/meetingSchedule/:id", controllers.MeetingListDelete)
	r.GET("/exportMeetingList", controllers.CreateExcelMeetingList)
	r.POST("/uploadMeetingList", controllers.ImportExcelMeetingList)

	r.POST("/uploadFileMeetingList", controllers.UploadHandlerMeetingList)
	r.GET("/downloadMeetingList/:id/:filename", controllers.DownloadFileHandlerMeetingList)
	r.DELETE("/deleteMeetingList/:id/:filename", controllers.DeleteFileHandlerMeetingList)
	r.GET("/filesMeetingList/:id", controllers.GetFilesByIDMeetingList)

	// ********** Route Timeline Project ********** //
	r.GET("/timelineProject", controllers.GetEventsProject)
	r.POST("/timelineProject", controllers.CreateEventProject)
	r.DELETE("/timelineProject/:id", controllers.DeleteEventProject)
	r.GET("/resourceProject", controllers.GetResourcesProject)
	r.POST("/resourceProject", controllers.CreateResourceProject)
	r.DELETE("/resourceProject/:id", controllers.DeleteResourceProject)
	r.GET("/exportTimelineProject", controllers.ExportTimelineProjectToExcel)

	r.Run(":8085")
}
