package router

import (
	"DFES-Web/api"
	"github.com/gin-gonic/gin"
)

type FileSystemRouter struct {
}

func (fsr *FileSystemRouter) InitFileSystemRouter(router *gin.RouterGroup) {
	router.POST("/operator", api.FileSystemApiInstance.OperateFileSystem)
	router.POST("/upload", api.FileSystemApiInstance.UploadFile)
	router.GET("/download", api.FileSystemApiInstance.DownloadFile)
}

var FileSystemRouterInstance = new(FileSystemRouter)
