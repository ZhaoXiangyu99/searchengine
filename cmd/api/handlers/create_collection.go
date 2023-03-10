package handlers

import (
	"searchengine/cmd/api/rpc"
	"searchengine/kitex_gen/collectionModel"
	"searchengine/pkg/constants"
	"searchengine/pkg/errno"

	jwt "github.com/appleboy/gin-jwt/v2"

	"github.com/gin-gonic/gin"
)

func CreateCollection(c *gin.Context) {
	var createColltVar CreateColltParam
	name := c.PostForm("name")
	if len(name) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}
	createColltVar.Name = name

	claim := jwt.ExtractClaims(c)
	UserId := int64(claim[constants.IdentityKey].(float64))
	createColltVar.UserID = UserId

	req := &collectionModel.CreateColltRequest{
		UserId: createColltVar.UserID,
		Name:   createColltVar.Name,
	}
	err := rpc.CreateCollection(c, req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	} else {
		SendResponse(c, errno.Success, nil)
	}
}
