package plugin

import (
	"github.com/gin-gonic/gin"
)

type Agent interface {
	Base
	RegisterUnAuthRouter(r *gin.RouterGroup)
	RegisterAuthUserRouter(r *gin.RouterGroup)
	RegisterAuthAdminRouter(r *gin.RouterGroup)
}

var (
	CallAgent,
	registerAgent = MakePlugin[Agent](true)
)
