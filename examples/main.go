package main

import (
	"github.com/gin-gonic/gin"
	"github.com/penglin1995/webflow"
	"github.com/penglin1995/webflow/examples/controllers"
	"github.com/penglin1995/webflow/logx"
)

func main() {
	webflow.Init(logx.New(logx.LevelError, "./"))

	e := gin.New()

	e.POST("/user", webflow.UseController(new(controllers.AddUserCTL)))

	e.Run(":8090")
}
