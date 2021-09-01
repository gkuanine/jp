package jp

import "github.com/gin-gonic/gin"

type Jp struct {
	*gin.Engine
}

func Ignite() *Jp {
	return &Jp{Engine: gin.New()}
}
func (this *Jp) Launch() {
	this.Run(":8080")
}
func (this *Jp) Mount(routers ...IRouter) *Jp {
	for _, router := range routers {
		router.BuildRouter(this)
	}
	return this
}
