package service

import (
	"DFES-Web/db"
	"DFES-Web/model/do"
	"gorm.io/gorm"
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

var FileSystemServiceInstance = new(FileSystemService)
