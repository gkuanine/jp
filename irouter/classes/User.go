package classes

import (
	"github.com/gin-gonic/gin"
	"nidemall/jp"
)

type UserClass struct{

}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass)UserList()gin.HandlerFunc{
	return func(context *gin.Context) {
			context.JSON(200,gin.H{"result":"sucess"})
	}
}
func (this *UserClass)Build(jp *jp.Jp){
	jp.Handle("GET","/user",this.UserList())
}
