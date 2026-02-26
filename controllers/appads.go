package controllers

import "github.com/gin-gonic/gin"

func ReturnAppAds(c *gin.Context) {
	c.Header("Content-Type", "text/plain")
	c.Header("Cache-Control", "max-age=3600")
	txt := `google.com, pub-9798070448976857, DIRECT, f08c47fec0942fa0`

	// 返回
	c.String(200, txt)
}
