package controllers

import "github.com/gin-gonic/gin"

func ReturnSitemap(c *gin.Context) {
	c.Header("Content-Type", "application/xml")
	sitemap := `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>https://www.yaoshang.tech</loc>
    </url>
    <url>
        <loc>https://www.yaoshang.tech/games/sudoku-crown</loc>
    </url>
    <url>
        <loc>https://www.yaoshang.tech/games/digit-merge</loc>
    </url>
    <url>
        <loc>https://www.yaoshang.tech/games/block-cuties</loc>
    </url>
    <url>
        <loc>https://www.yaoshang.tech/privacy</loc>
    </url>
    <url>
        <loc>https://www.yaoshang.tech/contact</loc>
    </url>
</urlset>`
	c.String(200, sitemap)
}

func ReturnRobotsTxt(c *gin.Context) {
	c.String(200, "Sitemap: https://www.yaoshang.tech/sitemap.xml")
}
