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
	group.GET("/", ctl.PageSudokuCrown)
	group.GET("/games/sudoku-crown", ctl.PageSudokuCrown)
	group.GET("/games/block-cuties", ctl.PageBlockCuties)
	group.GET("/games/digit-merge", ctl.PageDigitMerge)
	group.GET("/news", ctl.PageNews)
	group.GET("/privacy-policy", ctl.PagePrivacyPolicy)
	group.GET("/contact", ctl.PageContact)
	g.NoRoute(ctl.PageNoFound)
}
