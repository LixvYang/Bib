package framework

import "net/http"

type Core struct {

}

//初始化框架核心架构
func NewCore() *Core {
	return &Core{}
}

// 框架核心结构实现Handler接口
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) { 
	// TODO
}