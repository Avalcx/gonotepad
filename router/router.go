package router

import (
	"net/http"
	"notepad/config"
	"notepad/controller"
	"notepad/model"

	"github.com/gin-gonic/gin"
)

var tmp string = "tmp/"

func NotePad(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		// 获取path method,根据get or post 实现不同的方法
		textName := c.Request.URL.Path[1:]
		method := c.Request.Method

		// get时 仅读取
		if method == "GET" {
			err := controller.IsExsit(textName, config.TextType)
			if err != nil {
				c.JSON(http.StatusNotFound, err.Error())
				return
			}
			text, err := controller.Read(textName, config.TextType)
			if err != nil {
				c.JSON(http.StatusForbidden, err.Error())
				return
			}

			// query参数为json时，仅返回数据
			_, jsonIsExsit := c.GetQuery("json")
			if jsonIsExsit {
				c.JSON(http.StatusOK, text)
				return
			}
			// query参数为空时，返回html页面
			c.HTML(http.StatusOK, "index.html", gin.H{
				"Title":     "Notepad",
				"Content":   text,
				"CleanTime": config.CleanTime,
			})
			// post 时 仅写入
		} else if method == "POST" {

			_, jsonIsExsit := c.GetQuery("json")
			textForm, _ := c.GetPostForm("text")
			// 2种写入方式
			// query参数为json时，仅返回ok
			// 存在query参数时，不再往下执行
			if jsonIsExsit {
				err := controller.Write(textName, textForm, config.TextType)

				if err != nil {
					c.JSON(http.StatusForbidden, err.Error())
					return
				}
				c.JSON(http.StatusOK, "ok")
				return
			} else {
				// query参数为空时，返回html页面
				err := controller.Write(textName, textForm, config.TextType)
				if err != nil {
					c.JSON(http.StatusForbidden, err.Error())
					return
				}
				c.HTML(http.StatusOK, "index.html", gin.H{
					"Title":   "Notepad",
					"Content": textForm,
				})
			}
		}
	})
}

func CleanData(r *gin.Engine) {
	r.GET("/cleanData", func(c *gin.Context) {
		textName, ok := c.GetQuery("textName")
		if ok {
			ok = model.CleanList[textName]
			if ok {
				controller.Clean(textName)
				c.JSON(http.StatusOK, "clean success")
			} else {
				c.JSON(http.StatusOK, "textName is not exist")
			}
			return
		}
		c.JSON(http.StatusForbidden, "arguments is invaild")
	})
}
