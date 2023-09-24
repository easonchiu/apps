package routers

import (
	"ysgame/controllers"
	"ysgame/ui"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(g *gin.Engine) {
	g.SetHTMLTemplate(ui.GetTemplate())

	g.StaticFS("/static", ui.StaticFS())

	ctl := new(controllers.Controller)

	group := g.Group("")
	group.GET("/", ctl.PageHomepage)
	group.GET("/games/sudoku-crown", ctl.PageSudokuCrown)
	group.GET("/games/block-cuties", ctl.PageBlockCuties)
	group.GET("/games/digit-merge", ctl.PageDigitMerge)
	group.GET("/news", ctl.PageNews)
	group.GET("/privacy", ctl.PagePrivacy)
	group.GET("/contact", ctl.PageContact)

	group.GET("/robots.txt", ctl.PageRobots)
	group.GET("/sitemap.xml", ctl.PageSitemap)

	g.NoRoute(ctl.PageNoFound)
}
