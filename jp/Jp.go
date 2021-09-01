package jp

import "github.com/gin-gonic/gin"

type Jp struct{
	*gin.Engine
}
func Ignite() *Jp {
	return &Jp{Engine: gin.New()}
}
func (this *Jp)Launch(){

}
func (this *Jp)Mount(classes ...IClass)*Jp {
	for _,class:=range classes{
		class.Build(this)
	}
	return this;
}