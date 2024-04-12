package response

import "DFES-Web/model/do"

type TreeResponse struct {
	do.FileNode
	Children *[]TreeResponse
}
