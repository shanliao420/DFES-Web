package service

import (
	"DFES-Web/db"
	"DFES-Web/model/do"
	"gorm.io/gorm"
	"log"
)

type FileSystemService struct {
}

func (fss *FileSystemService) InitUserRootNode(user *do.UserModel) {
	err := db.GlobalMySQLClient.Transaction(func(tx *gorm.DB) error {
		rootNode := &do.FileNode{
			Name:   user.Username + "-space",
			Parent: uint64(0),
		}
		tx.Create(rootNode)
		userRootMap := &do.UserTreeRoot{
			UserId:     user.ID,
			TreeRootId: rootNode.ID,
		}
		tx.Create(userRootMap)
		tx.Commit()
		return nil
	})
	if err != nil {
		log.Println("init user space err:", err)
		return
	}
}

var FileSystemServiceInstance = new(FileSystemService)
