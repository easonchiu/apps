package routers

import (
	"io"
	"strconv"
	"strings"
	"ysgame/controllers"
	"ysgame/ui"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(g *gin.Engine) {
	// API routes
	api := g.Group("/wwwapi")
	{
		api.POST("/feedback", controllers.SubmitFeedback)
	}

	g.GET("/robots.txt", controllers.ReturnRobotsTxt)
	g.GET("/sitemap.xml", controllers.ReturnSitemap)
	g.GET("/app-ads.txt", controllers.ReturnAppAds)

	// Get the embedded file system
	distFS := ui.GetDistFS()

	// Helper function to serve a file from the embedded FS
	serveFile := func(c *gin.Context, filename string) bool {
		file, err := distFS.Open(filename)
		if err != nil {
			return false
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			return false
		}

		// Determine content type
		contentType := "application/octet-stream"
		if strings.HasSuffix(filename, ".html") {
			contentType = "text/html; charset=utf-8"
		} else if strings.HasSuffix(filename, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(filename, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(filename, ".svg") {
			contentType = "image/svg+xml"
		} else if strings.HasSuffix(filename, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(filename, ".ico") {
			contentType = "image/x-icon"
		}

		c.Header("Content-Type", contentType)
		c.Header("Content-Length", strconv.FormatInt(stat.Size(), 10))
		c.Status(200)
		io.Copy(c.Writer, file)
		return true
	}

	// Serve all requests
	g.NoRoute(func(c *gin.Context) {
		path := strings.TrimPrefix(c.Request.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}

		// Try to serve the requested file
		if serveFile(c, path) {
			return
		}

		// File doesn't exist, serve index.html for SPA routing
		if !serveFile(c, "index.html") {
			c.String(500, "Error loading application")
		}
	})
}
