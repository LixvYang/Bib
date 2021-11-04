package framework

import (
	"log"
	"net/http"
)

type Core struct {
	router map[string]ControllerHandler
}

//初始化框架核心架构
func NewCore() *Core {
	return &Core{}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	log.Println("core.serveHTTP")
	ctx := NewContext(request, response)

	// 一个简单的路由选择器，这里直接写死为测试路由foo
	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")

	router(ctx)

}