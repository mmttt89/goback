// route/route.go

package route

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	UserRouter(db, r)
	TestRouter(r)
}
