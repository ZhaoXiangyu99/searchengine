package handlers

import (
	"context"
	"searchengine/cmd/api/idgen"
	"searchengine/cmd/api/rpc"
	searchapi "searchengine/kitex_gen/SearchApi"
	"searchengine/pkg/errno"

	"github.com/gin-gonic/gin"
)

// 添加索引直接接口
func Add(c *gin.Context) {
	var addVar AddParam
	if err := c.ShouldBind(addVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}

	//请求字段特殊情况处理
	//TODO

	//生成ID
	addVar.Id = idgen.GetID()
	addVar.Text = c.Query("text")
	addVar.Url = c.Query("url")

	//RPC实际请求
	req := &searchapi.AddRequest{
		Id:   addVar.Id,
		Text: addVar.Text,
		Url:  addVar.Url,
	}
	err := rpc.Add(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	} else {
		SendResponse(c, errno.Success, nil)
	}
}
