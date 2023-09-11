package controllers

import "github.com/gin-gonic/gin"

type Controller struct{}

func (Controller) PageSudokuCrown(ctx *gin.Context) {
	ctx.HTML(200, "games.html", gin.H{
		"active": "games",
		"title":  "Games",
		"game":   "sudoku_crown",
	})
}

func (Controller) PageBlockCuties(ctx *gin.Context) {
	ctx.HTML(200, "games.html", gin.H{
		"active": "games",
		"title":  "Games",
		"game":   "block_cuties",
	})
}

func (Controller) PageDigitMerge(ctx *gin.Context) {
	ctx.HTML(200, "games.html", gin.H{
		"active": "games",
		"title":  "Games",
		"game":   "digit_merge",
	})
}

func (Controller) PagePrivacyPolicy(ctx *gin.Context) {
	ctx.HTML(200, "privacy_policy.html", gin.H{
		"active": "privacy_policy",
		"title":  "Privacy Policy",
	})
}

func (Controller) PageNews(ctx *gin.Context) {
	ctx.HTML(200, "news.html", gin.H{
		"active": "news",
		"title":  "News",
	})
}

func (Controller) PageContact(ctx *gin.Context) {
	ctx.HTML(200, "contact.html", gin.H{
		"active": "contact",
		"title":  "Content",
	})
}

func (Controller) PageNoFound(ctx *gin.Context) {
	ctx.HTML(200, "404.html", nil)
}