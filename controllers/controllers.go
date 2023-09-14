package controllers

import "github.com/gin-gonic/gin"

type Controller struct{}

func (Controller) PageHomepage(ctx *gin.Context) {
	ctx.HTML(200, "home.html", gin.H{
		"active":      "games",
		"description": "Enjoy hours of entertainment with our exciting lineup of casual games designed to captivate players of all ages.",
		"keywords":    "Sudoku Crown, Block Cuties, Digit Merge!, Casual games, Addictive games, Relaxing games, Fun games, Engaging games, Puzzle games, Arcade games, Mobile games, Online games, Family-friendly games, Free games, Challenging games, Adventure games, Strategy games, Multiplayer games",
	})
}

func (Controller) PageSudokuCrown(ctx *gin.Context) {
	ctx.HTML(200, "games.html", gin.H{
		"active":      "games",
		"title":       "Sudoku Crown",
		"description": "Challenge Your Brain with Sudoku Game!",
		"keywords":    "Sudoku Crown, Sudoku game, Free sudoku, Puzzle game, Number game, Logic game, Brain teaser, Mind game, Strategy game, Board game, Sudoku solver, Sudoku puzzles, Challenging puzzles, Sudoku strategies, Sudoku techniques, Sudoku tips, Sudoku variations, Daily Sudoku, Online Sudoku, Printable Sudoku, Beginner Sudoku, Advanced Sudoku",
		"game":        "sudoku_crown",
	})
}

func (Controller) PageBlockCuties(ctx *gin.Context) {
	ctx.HTML(200, "games.html", gin.H{
		"active":      "games",
		"title":       "Block Cuties",
		"description": "An Addictive Block Blast Game! Experience Endless Fun with Block Cuties!",
		"keywords":    "Block Cuties, Block Puzzle game, Puzzle game, Tetris-style game, Brain teaser, Mind game, Logic game, Tetromino game, Strategy game, Tile game, Challenging puzzles, Addictive gameplay, Spatial awareness, Pattern recognition, Shape matching, Block manipulation, Casual game, Mobile game, Offline game, Free-to-play game, Relaxing game",
		"game":        "block_cuties",
	})
}

func (Controller) PageDigitMerge(ctx *gin.Context) {
	ctx.HTML(200, "games.html", gin.H{
		"active":      "games",
		"title":       "Digit Merge!",
		"description": "Merge Gears and Balls, Challenge Your Number Intelligence!",
		"keywords":    "Digit Merge!, Physics-based game, Digit puzzle game, Digit elimination game, Digit puzzle-solving game, Numeric logic game, Math game, Digit strategy game, Digit challenge game, Digit slide game, Digit stacking game, Physics simulation game, Digit synthesis game, Digit collision game, Digit reasoning game, Digit maze game, Digit sorting game, Digit calculation game, Digit linking game",
		"game":        "digit_merge",
	})
}

func (Controller) PagePrivacy(ctx *gin.Context) {
	ctx.HTML(200, "privacy.html", gin.H{
		"active":      "privacy_policy",
		"title":       "Privacy Policy",
		"description": "Enjoy hours of entertainment with our exciting lineup of casual games designed to captivate players of all ages.",
		"keywords":    "Sudoku Crown, Block Cuties, Digit Merge!, Casual games, Addictive games, Relaxing games, Fun games, Engaging games, Puzzle games, Arcade games, Mobile games, Online games, Family-friendly games, Free games, Challenging games, Adventure games, Strategy games, Multiplayer games",
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
		"description": "Enjoy hours of entertainment with our exciting lineup of casual games designed to captivate players of all ages.",
		"keywords":    "Sudoku Crown, Block Cuties, Digit Merge!, Casual games, Addictive games, Relaxing games, Fun games, Engaging games, Puzzle games, Arcade games, Mobile games, Online games, Family-friendly games, Free games, Challenging games, Adventure games, Strategy games, Multiplayer games",
	})
}

func (Controller) PageNoFound(ctx *gin.Context) {
	ctx.HTML(200, "404.html", nil)
}

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
