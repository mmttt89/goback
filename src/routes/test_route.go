package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestRouter(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "hello, i'm working")
	})
}
