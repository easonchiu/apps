package controllers

import "github.com/gin-gonic/gin"

type Controller struct{}

func (Controller) PageHomepage(ctx *gin.Context) {
	ctx.HTML(200, "home.html", gin.H{
		"active":      "games",
		"description": "",
		"keywords":    "",
	})
}

func (Controller) PageSudokuCrown(ctx *gin.Context) {
	ctx.HTML(200, "games.html", gin.H{
		"active":      "games",
		"title":       "Sudoku Crown",
		"description": "Challenge Your Brain with Sudoku Game!",
		"keywords":    "",
		"game":        "sudoku_crown",
	})
}

func (Controller) PageBlockCuties(ctx *gin.Context) {
	ctx.HTML(200, "games.html", gin.H{
		"active":      "games",
		"title":       "Block Cuties",
		"description": "An Addictive Block Blast Game! Experience Endless Fun with Block Cuties!",
		"keywords":    "",
		"game":        "block_cuties",
	})
}

func (Controller) PageDigitMerge(ctx *gin.Context) {
	ctx.HTML(200, "games.html", gin.H{
		"active":      "games",
		"title":       "Digit Merge!",
		"description": "Merge Gears and Balls, Challenge Your Number Intelligence!",
		"keywords":    "",
		"game":        "digit_merge",
	})
}

func (Controller) PagePrivacy(ctx *gin.Context) {
	ctx.HTML(200, "privacy.html", gin.H{
		"active":      "privacy_policy",
		"title":       "Privacy Policy",
		"description": "",
		"keywords":    "",
	})
}

func (Controller) PageNews(ctx *gin.Context) {
	ctx.HTML(200, "news.html", gin.H{
		"active":      "news",
		"title":       "News",
		"description": "",
		"keywords":    "",
	})
}

func (Controller) PageContact(ctx *gin.Context) {
	ctx.HTML(200, "contact.html", gin.H{
		"active":      "contact",
		"title":       "Content",
		"description": "",
		"keywords":    "",
	})
}

func (Controller) PageNoFound(ctx *gin.Context) {
	ctx.HTML(200, "404.html", nil)
}
