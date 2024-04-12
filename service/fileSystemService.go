package service

import (
	"DFES-Web/db"
	"DFES-Web/model/do"
	"context"
	"gorm.io/gorm"
	"io"
	"log"
)

type FileSystemService struct {
}

func (fss *FileSystemService) InitUserRootNode(user *do.UserModel) error {
	err := db.GlobalMySQLClient.Transaction(func(tx *gorm.DB) error {
		rootNode := &do.FileNode{
			Name:   user.Username + "-space",
			Parent: uint64(0),
		}
		if err := tx.Create(rootNode).Error; err != nil {
			tx.Rollback()
			return err
		}
		userRootMap := &do.UserTreeRoot{
			UserId:     user.ID,
			TreeRootId: rootNode.ID,
		}
		if err := tx.Create(userRootMap).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
	if err != nil {
		log.Println("init user space err:", err)
		return err
	}
	return nil
}

func (fss *FileSystemService) CreateNode(newNode *do.FileNode) error {
	return db.GlobalMySQLClient.Create(newNode).Error
}

func (fss *FileSystemService) DeleteNode(id uint64) error {
	return db.GlobalMySQLClient.Delete(&do.FileNode{}, id).Error
}

func (fss *FileSystemService) ModifyNode(node *do.FileNode) error {
	return db.GlobalMySQLClient.Updates(node).Error
}

func (fss *FileSystemService) Exists(id uint64) bool {
	var cnt int64
	db.GlobalMySQLClient.First(&do.FileNode{}, id).Count(&cnt)
	return cnt == 1
}

func (fss *FileSystemService) ExistsInDirectory(parent uint64, filename string) bool {
	var cnt int64
	db.GlobalMySQLClient.Model(&do.FileNode{}).Where("parent_id = ? and name = ?", parent, filename).Count(&cnt)
	return cnt == 1
}

func (fss *FileSystemService) Upload(node *do.FileNode, reader io.Reader, stream bool) error {
	if stream {
		dataId, err := db.GlobalDFESClient.PushStream(context.Background(), reader)
		if err != nil {
			log.Println("upload err:", err)
			return err
		}
		node.DataId = dataId
		db.GlobalMySQLClient.Create(node)
		return nil
	}
	b, err := io.ReadAll(reader)
	if err != nil {
		log.Println("read reader in upload err:", err)
		return err
	}
	dataId, err := db.GlobalDFESClient.Push(context.Background(), b)
	node.DataId = dataId
	db.GlobalMySQLClient.Create(node)
	return nil
}

var FileSystemServiceInstance = new(FileSystemService)
