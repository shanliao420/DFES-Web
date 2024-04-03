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

var FileSystemServiceInstance = new(FileSystemService)
