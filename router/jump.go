package router

import (
	"net/http"
	"notepad/utils"

	"github.com/gin-gonic/gin"
)

func Jump(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		textName := utils.RandStr()
		c.Redirect(http.StatusFound, "/"+textName)
	})
}
