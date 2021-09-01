package irouter

import (
	"github.com/gin-gonic/gin"
	"jp/jp"
)

type IndexRouter struct {
}

func NewIndexRouter() *IndexRouter {
	return &IndexRouter{}
}

func (this *IndexRouter) Index() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{"result": "hello golang "})
	}
}

func (this *IndexRouter) BuildRouter(jp *jp.Jp) {
	jp.Handle("GET", "/", this.Index())
}
