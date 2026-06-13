package http

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	AddUser(ctx *gin.Context)
	AddProvider(ctx *gin.Context)
	FirstBuy(ctx *gin.Context)
	RepeatBuy(ctx *gin.Context)
}

func NewRouter(handler HttpHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/user", handler.AddUser)
	r.POST("/provider", handler.AddProvider)
	r.POST("/first-buy", handler.FirstBuy)
	r.POST("/repeat-buy", handler.RepeatBuy)

	return r
}
