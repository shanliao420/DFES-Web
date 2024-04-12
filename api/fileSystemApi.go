package api

import (
	"DFES-Web/db"
	"DFES-Web/model/do"
	"DFES-Web/model/request"
	"DFES-Web/model/response"
	"DFES-Web/service"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"strconv"
)

const (
	DefaultMaxSingleFile = 100 * 1024 * 1024 // 上传的文件超过100M使用流式传输
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
	fileHeader, err := c.FormFile("file")
	if err != nil {
		log.Println("get file header from request err:", err)
		response.FailWithMessage("请重试", c)
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		log.Println("open file header err:", err)
		response.FailWithMessage("请重试", c)
		return
	}
	name := c.PostForm("name")
	parent, err := strconv.ParseUint(c.PostForm("parent"), 10, 64)
	if err != nil {
		log.Println("parse parent err:", err)
		response.FailWithMessage("参数错误", c)
		return
	}
	if !service.FileSystemServiceInstance.Exists(parent) {
		response.FailWithMessage("新增节点所选父节点不存在", c)
		return
	}
	if service.FileSystemServiceInstance.ExistsInDirectory(parent, name) {
		response.FailWithMessage("所选目录存在同名文件", c)
		return
	}
	node := &do.FileNode{
		Name:     name,
		Parent:   parent,
		Kind:     do.FileKind,
		FileSize: uint64(fileHeader.Size),
	}
	err = service.FileSystemServiceInstance.Upload(node, file, fileHeader.Size > DefaultMaxSingleFile)
	if err != nil {
		response.FailWithMessage("请稍后重试", c)
		return
	}
	response.OkWithMessage("上传成功", c)
}

func (fsa *FileSystemApi) DownloadFile(c *gin.Context) {
	id, err := string2uint64(c.Query("id"))
	if err != nil {
		log.Println("get id err:", err)
		response.FailWithMessage("参数错误", c)
		return
	}
	node := service.FileSystemServiceInstance.GetNode(id)
	reader, err := db.GlobalDFESClient.GetStream(context.Background(), node.DataId)
	if err != nil {
		log.Println("get file stream err:", err)
		response.FailWithMessage("请稍后重试", c)
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+node.Name)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Length", strconv.FormatUint(node.FileSize, 10))
	_, err = io.Copy(c.Writer, reader)
	if err != nil {
		log.Println("copy err:", err)
		response.FailWithMessage("请稍后重试", c)
		return
	}
	c.Writer.Flush()
}

func CheckKind(kind byte) bool {
	if kind == do.FileKind || kind == do.DirectoryKind {
		return true
	}
	return false
}

func string2uint64(num string) (uint64, error) {
	return strconv.ParseUint(num, 10, 64)
}

var FileSystemApiInstance = new(FileSystemApi)
