package router

import (
	"DFES-Web/api"
	"github.com/gin-gonic/gin"
)

type FileSystemRouter struct {
}

func (fsr *FileSystemRouter) InitFileSystemRouter(router *gin.RouterGroup) {
	router.POST("/operator", api.FileSystemApiInstance.OperateFileSystem)
}

var FileSystemRouterInstance = new(FileSystemRouter)
