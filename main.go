package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"notepad/config"
	"notepad/router"

	"github.com/gin-gonic/gin"
)

//go:embed static/*
var Static embed.FS

func init() {
	config.SetConfig()
}

func main() {
	r := gin.Default()
	static, _ := fs.Sub(Static, "static")
	r.StaticFS("/static", http.FS(static))
	templ := template.Must(template.New("").ParseFS(Static, "static/*.html"))
	r.SetHTMLTemplate(templ)
	router.Jump(r)
	router.NotePad(r)
	r.Run("127.0.0.1:" + config.Port)
}
