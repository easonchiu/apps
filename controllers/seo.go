package controllers

import "github.com/gin-gonic/gin"

func (Controller) PageRobots(ctx *gin.Context) {
	ctx.String(200, "Sitemap: https://www.ysgamestudio.com/sitemap.xml")
}

func (Controller) PageSitemap(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/xml")
	// 构建Sitemap数据
	sitemap := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>https://www.ysgamestudio.com</loc>
    </url>
    <url>
        <loc>https://www.ysgamestudio.com/games/sudoku-crown</loc>
    </url>
    <url>
        <loc>https://www.ysgamestudio.com/games/digit-merge</loc>
    </url>
    <url>
        <loc>https://www.ysgamestudio.com/games/block-cuties</loc>
    </url>
    <url>
        <loc>https://www.ysgamestudio.com/privacy</loc>
    </url>
    <url>
        <loc>https://www.ysgamestudio.com/contact</loc>
    </url>
</urlset>`

	// 返回XML响应
	ctx.String(200, sitemap)
}
