package pack

import (
	"errors"
	"searchengine/kitex_gen/collectionModel"
	"searchengine/pkg/errno"
	"time"
)

func BuildBaseResp(err error) *collectionModel.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *collectionModel.BaseResp {
	return &collectionModel.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
		ServiceTime:   time.Now().Unix(),
	}
}
