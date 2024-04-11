package api

import (
	"DFES-Web/model/do"
	"DFES-Web/model/request"
	"DFES-Web/model/response"
	"DFES-Web/service"
	"github.com/gin-gonic/gin"
	"log"
)

type FileSystemApi struct {
}

func (fsa *FileSystemApi) OperateFileSystem(c *gin.Context) {
	var fileSystemOperateInfo request.FileSystemOperateInfo
	err := c.ShouldBindJSON(&fileSystemOperateInfo)
	if err != nil {
		log.Println("operate file system err:", err)
		response.FailWithMessage("请求参数有误", c)
		return
	}
	switch fileSystemOperateInfo.OperatorCode {
	case request.OpCreate:
		if !service.FileSystemServiceInstance.Exists(fileSystemOperateInfo.Parent) {
			response.FailWithMessage("新增节点所选父节点不存在", c)
			return
		}
		if !CheckKind(fileSystemOperateInfo.Kind) {
			response.FailWithMessage("新增节点类型错误", c)
			return
		}
		err = service.FileSystemServiceInstance.CreateNode(&do.FileNode{
			Name:   fileSystemOperateInfo.Name,
			Parent: fileSystemOperateInfo.Parent,
			Kind:   fileSystemOperateInfo.Kind,
		})
	case request.OpDelete:
		if !service.FileSystemServiceInstance.Exists(fileSystemOperateInfo.ID) {
			response.FailWithMessage("删除操作所选节点不存在", c)
			return
		}
		err = service.FileSystemServiceInstance.DeleteNode(fileSystemOperateInfo.ID)
	case request.OpUpdate:
		if !service.FileSystemServiceInstance.Exists(fileSystemOperateInfo.Parent) {
			response.FailWithMessage("更新节点所选父节点不存在", c)
			return
		}
		if !CheckKind(fileSystemOperateInfo.Kind) {
			response.FailWithMessage("更新节点类型错误", c)
			return
		}
		err = service.FileSystemServiceInstance.ModifyNode(&do.FileNode{
			Name:   fileSystemOperateInfo.Name,
			Parent: fileSystemOperateInfo.Parent,
			Kind:   fileSystemOperateInfo.Kind,
		})
	default:
		response.FailWithMessage("请求参数有误", c)
		return
	}
	if err != nil {
		log.Println("process operator:", fileSystemOperateInfo.OperatorCode, " err:", err)
		response.FailWithMessage("操作失败，请重试", c)
		return
	}
	response.OkWithMessage("操作成功", c)
}

func (fsa *FileSystemApi) UploadFile(c *gin.Context) {
	//fileHeader, err := c.FormFile("file")
	//if err != nil {
	//	log.Println("get file header from request err:", err)
	//	response.FailWithMessage("请重试", c)
	//	return
	//}
	//file, err := fileHeader.Open()
	//if err != nil {
	//	log.Println("open file header err:", err)
	//	response.FailWithMessage("请重试", c)
	//	return
	//}
	//var dataId string
	//if fileHeader.Size > 1024*1024*1024 {
	//	dataId, err = db.GlobalDFESClient.PushStream(context.Background(), file)
	//	if err != nil {
	//		log.Println("上传失败，请重试", c)
	//		return
	//	}
	//} else {
	//	b, err := io.ReadAll(file)
	//	if err != nil {
	//		log.Println("read bytes err:", err)
	//		response.FailWithMessage("请重试", c)
	//		return
	//	}
	//	dataId, err = db.GlobalDFESClient.Push(context.Background(), b)
	//	if err != nil {
	//		log.Println("上传失败，请重试", c)
	//		return
	//	}
	//}
	//parentId := c.PostForm("parent")

}

func CheckKind(kind byte) bool {
	if kind == do.FileKind || kind == do.DirectoryKind {
		return true
	}
	return false
}

var FileSystemApiInstance = new(FileSystemApi)
