package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 代理 Vite 开发服务器
	viteDevServer, _ := url.Parse("http://localhost:5173")
	proxy := httputil.NewSingleHostReverseProxy(viteDevServer)

	// 定义 API 路由组（优先级高）
	api := r.Group("/api")
	{
		api.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Hello from Gin!"})
		})
		api.GET("/home", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"title": "Welcome to the Home Page"})
		})
		api.GET("/about", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"title": "Welcome to the About Page"})
		})
	}

	// 非 `/api` 请求 → 代理到 Vite
	r.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// 如果请求的是 `/api` 但未定义 → 返回 404
		if strings.HasPrefix(path, "/api") {
			c.JSON(http.StatusNotFound, gin.H{"error": "API endpoint not found"})
			return
		}

		// 其他请求（前端路由/静态文件）→ 代理到 Vite
		proxy.ServeHTTP(c.Writer, c.Request)
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
		return
	}
}
