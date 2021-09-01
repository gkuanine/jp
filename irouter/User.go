package irouter

import (
	"github.com/gin-gonic/gin"
	"jp/jp"
)

type UserRouter struct {
}

func NewUserRouter() *UserRouter {
	return &UserRouter{}
}

func (this *UserRouter) UserList() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{"result": "sucess"})
	}
}
func (this *UserRouter) BuildRouter(jp *jp.Jp) {
	jp.Handle("GET", "/user", this.UserList())
}
