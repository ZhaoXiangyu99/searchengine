// Code generated by Kitex v0.3.1. DO NOT EDIT.
package search

import (
	"github.com/cloudwego/kitex/server"
	searchapi "searchengine3090ti/kitex_gen/SearchApi"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler searchapi.Search, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
